// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-1-project/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserAdder is an autogenerated mock type for the UserAdder type
type UserAdder struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, user
func (_m *UserAdder) Add(ctx context.Context, user *entity.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserAdder interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserAdder creates a new instance of UserAdder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserAdder(t mockConstructorTestingTNewUserAdder) *UserAdder {
	mock := &UserAdder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
