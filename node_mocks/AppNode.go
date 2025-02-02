// Code generated by mockery v1.0.0. DO NOT EDIT.

package node_mocks

import common "github.com/anycable/anycable-go/common"
import mock "github.com/stretchr/testify/mock"
import node "github.com/anycable/anycable-go/node"

// AppNode is an autogenerated mock type for the AppNode type
type AppNode struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: s
func (_m *AppNode) Authenticate(s *node.Session) (*common.ConnectResult, error) {
	ret := _m.Called(s)

	var r0 *common.ConnectResult
	if rf, ok := ret.Get(0).(func(*node.Session) *common.ConnectResult); ok {
		r0 = rf(s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.ConnectResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*node.Session) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Disconnect provides a mock function with given fields: s
func (_m *AppNode) Disconnect(s *node.Session) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*node.Session) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandlePubSub provides a mock function with given fields: msg
func (_m *AppNode) HandlePubSub(msg []byte) {
	_m.Called(msg)
}

// LookupSession provides a mock function with given fields: id
func (_m *AppNode) LookupSession(id string) *node.Session {
	ret := _m.Called(id)

	var r0 *node.Session
	if rf, ok := ret.Get(0).(func(string) *node.Session); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*node.Session)
		}
	}

	return r0
}

// Perform provides a mock function with given fields: s, msg
func (_m *AppNode) Perform(s *node.Session, msg *common.Message) (*common.CommandResult, error) {
	ret := _m.Called(s, msg)

	var r0 *common.CommandResult
	if rf, ok := ret.Get(0).(func(*node.Session, *common.Message) *common.CommandResult); ok {
		r0 = rf(s, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.CommandResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*node.Session, *common.Message) error); ok {
		r1 = rf(s, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: s, msg
func (_m *AppNode) Subscribe(s *node.Session, msg *common.Message) (*common.CommandResult, error) {
	ret := _m.Called(s, msg)

	var r0 *common.CommandResult
	if rf, ok := ret.Get(0).(func(*node.Session, *common.Message) *common.CommandResult); ok {
		r0 = rf(s, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.CommandResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*node.Session, *common.Message) error); ok {
		r1 = rf(s, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unsubscribe provides a mock function with given fields: s, msg
func (_m *AppNode) Unsubscribe(s *node.Session, msg *common.Message) (*common.CommandResult, error) {
	ret := _m.Called(s, msg)

	var r0 *common.CommandResult
	if rf, ok := ret.Get(0).(func(*node.Session, *common.Message) *common.CommandResult); ok {
		r0 = rf(s, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.CommandResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*node.Session, *common.Message) error); ok {
		r1 = rf(s, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
