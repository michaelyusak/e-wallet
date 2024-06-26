// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "e-wallet/entity"

	mock "github.com/stretchr/testify/mock"
)

// RPTRepository is an autogenerated mock type for the RPTRepository type
type RPTRepository struct {
	mock.Mock
}

// DeleteExisting provides a mock function with given fields: ctx, userId
func (_m *RPTRepository) DeleteExisting(ctx context.Context, userId int) error {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteExisting")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTokenExpiredAt provides a mock function with given fields: ctx, token
func (_m *RPTRepository) GetTokenExpiredAt(ctx context.Context, token string) (*entity.ResetPasswordToken, error) {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenExpiredAt")
	}

	var r0 *entity.ResetPasswordToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.ResetPasswordToken, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.ResetPasswordToken); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ResetPasswordToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostOneRPT provides a mock function with given fields: ctx, rptReq, userId
func (_m *RPTRepository) PostOneRPT(ctx context.Context, rptReq entity.ResetPasswordToken, userId int) error {
	ret := _m.Called(ctx, rptReq, userId)

	if len(ret) == 0 {
		panic("no return value specified for PostOneRPT")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.ResetPasswordToken, int) error); ok {
		r0 = rf(ctx, rptReq, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRPTRepository creates a new instance of RPTRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRPTRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RPTRepository {
	mock := &RPTRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
