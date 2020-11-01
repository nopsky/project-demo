// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/nopsky/project-demo/internal/user/entity"
	mock "github.com/stretchr/testify/mock"
)

// IUserRepo is an autogenerated mock type for the IUserRepo type
type IUserRepo struct {
	mock.Mock
}

// GetUserByUsername provides a mock function with given fields: ctx, username
func (_m *IUserRepo) GetUserByUsername(ctx context.Context, username string) (*entity.UserEntity, error) {
	ret := _m.Called(ctx, username)

	var r0 *entity.UserEntity
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.UserEntity); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *IUserRepo) Save(ctx context.Context, _a1 *entity.UserEntity) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.UserEntity) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}