// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	net "perun.network/go-perun/wire/net"
)

// Listener is an autogenerated mock type for the Listener type
type Listener struct {
	mock.Mock
}

// Accept provides a mock function with given fields:
func (_m *Listener) Accept() (net.Conn, error) {
	ret := _m.Called()

	var r0 net.Conn
	if rf, ok := ret.Get(0).(func() net.Conn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Conn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Listener) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
