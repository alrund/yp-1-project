// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-1-project/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserRegistrator is an autogenerated mock type for the UserRegistrator type
type UserRegistrator struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, user
func (_m *UserRegistrator) Add(ctx context.Context, user *entity.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByLogin provides a mock function with given fields: ctx, login
func (_m *UserRegistrator) GetByLogin(ctx context.Context, login string) (*entity.User, error) {
	ret := _m.Called(ctx, login)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRegistrator interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRegistrator creates a new instance of UserRegistrator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRegistrator(t mockConstructorTestingTNewUserRegistrator) *UserRegistrator {
	mock := &UserRegistrator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
