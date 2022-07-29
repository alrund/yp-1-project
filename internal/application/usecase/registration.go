package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

type RegistrationData struct {
	Login    string
	Password string
}

func Registration(
	ctx context.Context,
	regData RegistrationData,
	sessionCookieName, sessionCookieDuration string,
	repository port.UserRepository,
	encryptor port.Encryptor,
	cooker port.Cooker,
	hasher port.PasswordHasher,
	w http.ResponseWriter,
) error {
	user, err := repository.GetByLogin(ctx, regData.Login)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return err
	}

	if user != nil {
		return ErrLoginAlreadyUse
	}

	user = &entity.User{
		ID:           uuid.New(),
		Login:        regData.Login,
		PasswordHash: hasher.Hash(regData.Password),
	}
	err = repository.Add(ctx, user)
	if err != nil {
		return err
	}

	encryptedUserID, err := encryptor.Encrypt(user.ID.String())
	if err != nil {
		return err
	}

	err = cooker.AddCookieWithDuration(sessionCookieName, encryptedUserID, sessionCookieDuration, w)
	if err != nil {
		return err
	}

	return nil
}
