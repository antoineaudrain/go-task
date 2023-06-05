// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// Login provides a mock function with given fields: email, password
func (_m *Service) Login(email string, password string) (string, error) {
	ret := _m.Called(email, password)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type Service_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - email string
//   - password string
func (_e *Service_Expecter) Login(email interface{}, password interface{}) *Service_Login_Call {
	return &Service_Login_Call{Call: _e.mock.On("Login", email, password)}
}

func (_c *Service_Login_Call) Run(run func(email string, password string)) *Service_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Service_Login_Call) Return(_a0 string, _a1 error) *Service_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_Login_Call) RunAndReturn(run func(string, string) (string, error)) *Service_Login_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: email, password
func (_m *Service) Register(email string, password string) (string, error) {
	ret := _m.Called(email, password)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type Service_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - email string
//   - password string
func (_e *Service_Expecter) Register(email interface{}, password interface{}) *Service_Register_Call {
	return &Service_Register_Call{Call: _e.mock.On("Register", email, password)}
}

func (_c *Service_Register_Call) Run(run func(email string, password string)) *Service_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Service_Register_Call) Return(_a0 string, _a1 error) *Service_Register_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_Register_Call) RunAndReturn(run func(string, string) (string, error)) *Service_Register_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
