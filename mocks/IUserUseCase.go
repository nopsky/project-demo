// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	user "github.com/nopsky/project-demo/proto/user"
)

// IUserUseCase is an autogenerated mock type for the IUserUseCase type
type IUserUseCase struct {
	mock.Mock
}

// RegisterUser provides a mock function with given fields: ctx, req
func (_m *IUserUseCase) RegisterUser(ctx context.Context, req *user.RegisterUserRequest) (*user.RegisterUserResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *user.RegisterUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.RegisterUserRequest) *user.RegisterUserResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.RegisterUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.RegisterUserRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
