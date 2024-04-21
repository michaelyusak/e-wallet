// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "e-wallet/entity"

	decimal "github.com/shopspring/decimal"

	mock "github.com/stretchr/testify/mock"

	repository "e-wallet/repository"
)

// WalletRepository is an autogenerated mock type for the WalletRepository type
type WalletRepository struct {
	mock.Mock
}

// GetCashFlow provides a mock function with given fields: ctx, wallet
func (_m *WalletRepository) GetCashFlow(ctx context.Context, wallet *entity.Wallet) error {
	ret := _m.Called(ctx, wallet)

	if len(ret) == 0 {
		panic("no return value specified for GetCashFlow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Wallet) error); ok {
		r0 = rf(ctx, wallet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetWalletByNum provides a mock function with given fields: ctx, walletNum
func (_m *WalletRepository) GetWalletByNum(ctx context.Context, walletNum string) (*entity.Wallet, error) {
	ret := _m.Called(ctx, walletNum)

	if len(ret) == 0 {
		panic("no return value specified for GetWalletByNum")
	}

	var r0 *entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Wallet, error)); ok {
		return rf(ctx, walletNum)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Wallet); ok {
		r0 = rf(ctx, walletNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, walletNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWalletByUserId provides a mock function with given fields: ctx, userId
func (_m *WalletRepository) GetWalletByUserId(ctx context.Context, userId int) (*entity.Wallet, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetWalletByUserId")
	}

	var r0 *entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*entity.Wallet, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Wallet); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostOneTransaction provides a mock function with given fields: ctx, transferReq
func (_m *WalletRepository) PostOneTransaction(ctx context.Context, transferReq entity.Transaction) error {
	ret := _m.Called(ctx, transferReq)

	if len(ret) == 0 {
		panic("no return value specified for PostOneTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Transaction) error); ok {
		r0 = rf(ctx, transferReq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostOneWallet provides a mock function with given fields: ctx, userId
func (_m *WalletRepository) PostOneWallet(ctx context.Context, userId int) (*entity.Wallet, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for PostOneWallet")
	}

	var r0 *entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*entity.Wallet, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Wallet); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBalance provides a mock function with given fields: ctx, walletId, amount
func (_m *WalletRepository) UpdateBalance(ctx context.Context, walletId int, amount decimal.Decimal) error {
	ret := _m.Called(ctx, walletId, amount)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBalance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, decimal.Decimal) error); ok {
		r0 = rf(ctx, walletId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateGachaTrial provides a mock function with given fields: ctx, userId, getTrial
func (_m *WalletRepository) UpdateGachaTrial(ctx context.Context, userId int, getTrial decimal.Decimal) error {
	ret := _m.Called(ctx, userId, getTrial)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGachaTrial")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, decimal.Decimal) error); ok {
		r0 = rf(ctx, userId, getTrial)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithTx provides a mock function with given fields: ctx, fn
func (_m *WalletRepository) WithTx(ctx context.Context, fn func(repository.WalletRepository) error) error {
	ret := _m.Called(ctx, fn)

	if len(ret) == 0 {
		panic("no return value specified for WithTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(repository.WalletRepository) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWalletRepository creates a new instance of WalletRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWalletRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *WalletRepository {
	mock := &WalletRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
