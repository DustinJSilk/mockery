// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// VariadicReturnFunc is an autogenerated mock type for the VariadicReturnFunc type
type VariadicReturnFunc struct {
	mock.Mock
}

type VariadicReturnFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *VariadicReturnFunc) EXPECT() *VariadicReturnFunc_Expecter {
	return &VariadicReturnFunc_Expecter{mock: &_m.Mock}
}

// SampleMethod provides a mock function with given fields: str
func (_m *VariadicReturnFunc) SampleMethod(str string) func(string, []int, ...interface{}) {
	ret := _m.Called(str)

	var r0 func(string, []int, ...interface{})
	if rf, ok := ret.Get(0).(func(string) func(string, []int, ...interface{})); ok {
		r0 = rf(str)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(string, []int, ...interface{}))
		}
	}

	return r0
}

// VariadicReturnFunc_SampleMethod_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SampleMethod'
type VariadicReturnFunc_SampleMethod_Call struct {
	*mock.Call
}

// SampleMethod is a helper method to define mock.On call
//   - str string
func (_e *VariadicReturnFunc_Expecter) SampleMethod(str interface{}) *VariadicReturnFunc_SampleMethod_Call {
	return &VariadicReturnFunc_SampleMethod_Call{Call: _e.mock.On("SampleMethod", str)}
}

func (_c *VariadicReturnFunc_SampleMethod_Call) Run(run func(str string)) *VariadicReturnFunc_SampleMethod_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *VariadicReturnFunc_SampleMethod_Call) Return(_a0 func(string, []int, ...interface{})) *VariadicReturnFunc_SampleMethod_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VariadicReturnFunc_SampleMethod_Call) RunAndReturn(run func(string) func(string, []int, ...interface{})) *VariadicReturnFunc_SampleMethod_Call {
	_c.Call.Return(run)
	return _c
}

// NewVariadicReturnFunc creates a new instance of VariadicReturnFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVariadicReturnFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *VariadicReturnFunc {
	mock := &VariadicReturnFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
