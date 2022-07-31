// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TransactSupporter is an autogenerated mock type for the TransactSupporter type
type TransactSupporter struct {
	mock.Mock
}

// WithinTransaction provides a mock function with given fields: ctx, tFunc
func (_m *TransactSupporter) WithinTransaction(ctx context.Context, tFunc func(context.Context) error) error {
	ret := _m.Called(ctx, tFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, tFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTransactSupporter interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactSupporter creates a new instance of TransactSupporter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactSupporter(t mockConstructorTestingTNewTransactSupporter) *TransactSupporter {
	mock := &TransactSupporter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
