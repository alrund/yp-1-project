package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/infrastructure/helper"
)

func LoginHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !helper.HasContentType(r, "application/json") {
			a.Warn(w, r, http.StatusBadRequest, usecase.ErrInvalidRequestFormat)

			return
		}

		var cred usecase.Credential
		if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
			a.Warn(w, r, http.StatusBadRequest, usecase.ErrInvalidRequestFormat)

			return
		}

		user, err := usecase.Login(
			r.Context(),
			cred,
			a.UserRepository,
			a.Hasher,
		)
		if err != nil {
			switch {
			case errors.Is(err, usecase.ErrUserNotFound):
				a.Warn(w, r, http.StatusUnauthorized, usecase.ErrNotAuthenticated)
			default:
				a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)
			}

			return
		}

		err = usecase.SetCookie(
			user.ID,
			a.Config.SessionCookieName, a.Config.SessionCookieDuration,
			a.Encryptor,
			a.Cooker,
			w,
		)
		if err != nil {
			a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)

			return
		}

		a.PlainRespond(w, r, http.StatusOK, []byte(http.StatusText(http.StatusOK)))
	}

	return http.HandlerFunc(fn)
}
