//go:build integration

package tests

import (
	"fmt"
	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/infrastructure/handler"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/stretchr/testify/assert"
)

func SetupTestBalance(a *app.App) {
	db := a.Storage.Connect()
	db.Exec("INSERT INTO users(id, login, password) VALUES ($1, $2, $3)",
		"dbac0532-eaa4-44f8-8845-72be1b25a6ac", "login", "password")
	db.Exec("INSERT INTO users(id, login, password, current) VALUES ($1, $2, $3, $4)",
		"9a110553-f962-4ab4-ba92-7c3fb7a107f3", "login", "password", 555.5)
	db.Exec("INSERT INTO withdraws(id, order_number, user_id, sum, processed_at) VALUES ($1, $2, $3, $4, $5)",
		uuid.NewString(), "444444", "9a110553-f962-4ab4-ba92-7c3fb7a107f3", 666.6, time.Now())

}

func TearDownTestBalance(a *app.App) {
	db := a.Storage.Connect()
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM withdraws")
}

func (s *IntegrationTestSuite) TestBalance() {
	SetupTestBalance(s.app)
	defer func() {
		TearDownTestBalance(s.app)
	}()

	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "success",
			request: request{
				method:             http.MethodGet,
				target:             "/api/user/balance",
				sessionCookieName:  s.app.Config.SessionCookieName,
				sessionCookieValue: "dbac0532-eaa4-44f8-8845-72be1b25a6ac",
				body:               "",
				contentType:        "application/json",
			},
			want: want{
				code:        http.StatusOK,
				response:    `{"current":0, "withdrawn":0}`,
				contentType: "application/json",
			},
		},
		{
			name: "success with balance",
			request: request{
				method:             http.MethodGet,
				target:             "/api/user/balance",
				sessionCookieName:  s.app.Config.SessionCookieName,
				sessionCookieValue: "9a110553-f962-4ab4-ba92-7c3fb7a107f3",
				body:               "",
				contentType:        "application/json",
			},
			want: want{
				code:        http.StatusOK,
				response:    `{"current":555.5, "withdrawn":666.6}`,
				contentType: "application/json",
			},
		},
		{
			name: "fail with not authenticated",
			request: request{
				method:             http.MethodGet,
				target:             "/api/user/balance",
				sessionCookieName:  s.app.Config.SessionCookieName,
				sessionCookieValue: "",
				body:               "",
				contentType:        "application/json",
			},
			want: want{
				code:        http.StatusUnauthorized,
				response:    fmt.Sprintf(`{"warning":"%s"}`, usecase.ErrNotAuthenticated),
				contentType: "application/json",
			},
		},
	}
	t := s.T()
	for _, tt := range tests {
		s.Run(tt.name, func() {
			h := handler.BalanceHandler(s.app)

			w := s.MakeTestRequest(tt.request, h)
			res := w.Result()
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				s.logger.Fatal(err)
			}

			assert.Equal(t, tt.want.code, res.StatusCode)
			if tt.want.contentType == "application/json" {
				assert.JSONEqf(t, tt.want.response, string(resBody), w.Body.String())
			} else {
				assert.Equalf(t, tt.want.response, string(resBody), w.Body.String())
			}
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
