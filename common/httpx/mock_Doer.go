// Code generated by mockery v2.14.0. DO NOT EDIT.

package httpx

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockDoer is an autogenerated mock type for the Doer type
type MockDoer struct {
	mock.Mock
}

// Do provides a mock function with given fields: ctx, req
func (_m *MockDoer) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	ret := _m.Called(ctx, req)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *http.Response); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockDoer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDoer creates a new instance of MockDoer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDoer(t mockConstructorTestingTNewMockDoer) *MockDoer {
	mock := &MockDoer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
