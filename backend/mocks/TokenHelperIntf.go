// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	constants "e-wallet/constants"

	mock "github.com/stretchr/testify/mock"
)

// TokenHelperIntf is an autogenerated mock type for the TokenHelperIntf type
type TokenHelperIntf struct {
	mock.Mock
}

// CreateAndSign provides a mock function with given fields: userId, purpose
func (_m *TokenHelperIntf) CreateAndSign(userId int, purpose constants.Purpose) (string, error) {
	ret := _m.Called(userId, purpose)

	if len(ret) == 0 {
		panic("no return value specified for CreateAndSign")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int, constants.Purpose) (string, error)); ok {
		return rf(userId, purpose)
	}
	if rf, ok := ret.Get(0).(func(int, constants.Purpose) string); ok {
		r0 = rf(userId, purpose)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int, constants.Purpose) error); ok {
		r1 = rf(userId, purpose)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenHelperIntf creates a new instance of TokenHelperIntf. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenHelperIntf(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenHelperIntf {
	mock := &TokenHelperIntf{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}