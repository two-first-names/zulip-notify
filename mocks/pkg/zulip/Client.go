// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// SendStreamMessage provides a mock function with given fields: content, stream, topic
func (_m *Client) SendStreamMessage(content string, stream string, topic string) error {
	ret := _m.Called(content, stream, topic)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(content, stream, topic)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
