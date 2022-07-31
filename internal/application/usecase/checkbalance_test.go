package usecase

import (
	"context"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCheckBalance(t *testing.T) {
	type mock struct {
		userGetter  *mocks.UserGetter
		withdrawner *mocks.Withdrawner
	}

	type args struct {
		ctx    context.Context
		userID uuid.UUID
		user   *entity.User
	}

	userUUID := uuid.New()
	userMock := &entity.User{
		ID:           userUUID,
		Login:        "User login",
		PasswordHash: "Password hash",
		Current:      555.55,
	}
	mockPrepare := func(m *mock, ctx context.Context, userID uuid.UUID, user *entity.User) {
		userGetterMock := m.userGetter.On("Get", ctx, userID)
		if userID == userMock.ID {
			userGetterMock.Return(userMock, nil)
		} else {
			userGetterMock.Return(nil, ErrUserNotFound)
		}

		withdrawnerMock := m.withdrawner.On("GetWithdrawn", ctx, user)
		if user.ID == userMock.ID {
			withdrawnerMock.Return(float32(666.66), nil)
		} else {
			withdrawnerMock.Return(float32(0), nil)
		}
	}

	tests := []struct {
		name        string
		args        *args
		want        *Balance
		wantErr     error
		mockPrepare func(m *mock, ctx context.Context, userID uuid.UUID, user *entity.User)
	}{
		{
			"success",
			&args{
				context.Background(),
				userUUID,
				userMock,
			},
			&Balance{Current: 555.55, Withdrawn: 666.66},
			nil,
			mockPrepare,
		},
		{
			"fail",
			&args{
				context.Background(),
				uuid.New(),
				userMock,
			},
			nil,
			ErrUserNotFound,
			mockPrepare,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mock{
				userGetter:  new(mocks.UserGetter),
				withdrawner: new(mocks.Withdrawner),
			}

			if tt.mockPrepare != nil {
				tt.mockPrepare(m, tt.args.ctx, tt.args.userID, tt.args.user)
			}

			balance, err := CheckBalance(
				tt.args.ctx,
				tt.args.userID.String(),
				m.userGetter,
				m.withdrawner,
			)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Current, balance.Current)
			assert.Equal(t, tt.want.Withdrawn, balance.Withdrawn)
		})
	}
}
