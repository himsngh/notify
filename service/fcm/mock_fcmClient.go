package fcm

import (
	"github.com/appleboy/go-fcm"
	"github.com/stretchr/testify/mock"
)

type mockFcmClient struct {
	mock.Mock
}

func (_m *mockFcmClient) Send(msg *fcm.Message) (*fcm.Response, error) {
	ret := _m.Called(msg)

	var r0 *fcm.Response
	if rf, ok := ret.Get(0).(func(msg *fcm.Message) *fcm.Response); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Get(0).(*fcm.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(msg *fcm.Message) error); ok {
		r1 = rf(msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
