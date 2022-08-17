// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"
	request "faba_traning_project/internal/httpbody/request"

	mock "github.com/stretchr/testify/mock"

	response "faba_traning_project/internal/httpbody/response"
)

// IUser is an autogenerated mock type for the IUser type
type IUser struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, requestUser
func (_m *IUser) Create(ctx context.Context, requestUser request.CreateUser) (response.User, error) {
	ret := _m.Called(ctx, requestUser)

	var r0 response.User
	if rf, ok := ret.Get(0).(func(context.Context, request.CreateUser) response.User); ok {
		r0 = rf(ctx, requestUser)
	} else {
		r0 = ret.Get(0).(response.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, request.CreateUser) error); ok {
		r1 = rf(ctx, requestUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewIUserT interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUser creates a new instance of IUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUser(t NewIUserT) *IUser {
	mock := &IUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}