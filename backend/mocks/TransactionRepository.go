// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "e-wallet/entity"

	mock "github.com/stretchr/testify/mock"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// GetTransactions provides a mock function with given fields: ctx, walletNumber, param
func (_m *TransactionRepository) GetTransactions(ctx context.Context, walletNumber string, param entity.PaginationParameter) (*entity.TransactionList, error) {
	ret := _m.Called(ctx, walletNumber, param)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactions")
	}

	var r0 *entity.TransactionList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.PaginationParameter) (*entity.TransactionList, error)); ok {
		return rf(ctx, walletNumber, param)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.PaginationParameter) *entity.TransactionList); ok {
		r0 = rf(ctx, walletNumber, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.TransactionList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.PaginationParameter) error); ok {
		r1 = rf(ctx, walletNumber, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}