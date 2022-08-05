// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-1-project/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserAccrualer is an autogenerated mock type for the UserAccrualer type
type UserAccrualer struct {
	mock.Mock
}

type UserAccrualer_Expecter struct {
	mock *mock.Mock
}

func (_m *UserAccrualer) EXPECT() *UserAccrualer_Expecter {
	return &UserAccrualer_Expecter{mock: &_m.Mock}
}

// Accrual provides a mock function with given fields: ctx, user, accrual
func (_m *UserAccrualer) Accrual(ctx context.Context, user *entity.User, accrual float32) error {
	ret := _m.Called(ctx, user, accrual)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User, float32) error); ok {
		r0 = rf(ctx, user, accrual)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserAccrualer_Accrual_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Accrual'
type UserAccrualer_Accrual_Call struct {
	*mock.Call
}

// Accrual is a helper method to define mock.On call
//  - ctx context.Context
//  - user *entity.User
//  - accrual float32
func (_e *UserAccrualer_Expecter) Accrual(ctx interface{}, user interface{}, accrual interface{}) *UserAccrualer_Accrual_Call {
	return &UserAccrualer_Accrual_Call{Call: _e.mock.On("Accrual", ctx, user, accrual)}
}

func (_c *UserAccrualer_Accrual_Call) Run(run func(ctx context.Context, user *entity.User, accrual float32)) *UserAccrualer_Accrual_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.User), args[2].(float32))
	})
	return _c
}

func (_c *UserAccrualer_Accrual_Call) Return(_a0 error) *UserAccrualer_Accrual_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewUserAccrualer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserAccrualer creates a new instance of UserAccrualer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserAccrualer(t mockConstructorTestingTNewUserAccrualer) *UserAccrualer {
	mock := &UserAccrualer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
