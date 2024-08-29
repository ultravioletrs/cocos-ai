// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	sevsnp "github.com/google/go-sev-guest/proto/sevsnp"
	mock "github.com/stretchr/testify/mock"
)

// QuoteProvider is an autogenerated mock type for the QuoteProvider type
type QuoteProvider struct {
	mock.Mock
}

// GetRawQuote provides a mock function with given fields: reportData
func (_m *QuoteProvider) GetRawQuote(reportData [64]byte) ([]uint8, error) {
	ret := _m.Called(reportData)

	if len(ret) == 0 {
		panic("no return value specified for GetRawQuote")
	}

	var r0 []uint8
	var r1 error
	if rf, ok := ret.Get(0).(func([64]byte) ([]uint8, error)); ok {
		return rf(reportData)
	}
	if rf, ok := ret.Get(0).(func([64]byte) []uint8); ok {
		r0 = rf(reportData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint8)
		}
	}

	if rf, ok := ret.Get(1).(func([64]byte) error); ok {
		r1 = rf(reportData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsSupported provides a mock function with given fields:
func (_m *QuoteProvider) IsSupported() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSupported")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Product provides a mock function with given fields:
func (_m *QuoteProvider) Product() *sevsnp.SevProduct {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Product")
	}

	var r0 *sevsnp.SevProduct
	if rf, ok := ret.Get(0).(func() *sevsnp.SevProduct); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sevsnp.SevProduct)
		}
	}

	return r0
}

// NewQuoteProvider creates a new instance of QuoteProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuoteProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *QuoteProvider {
	mock := &QuoteProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
