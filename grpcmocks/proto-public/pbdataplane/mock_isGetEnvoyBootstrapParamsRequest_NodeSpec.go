// Code generated by mockery v2.41.0. DO NOT EDIT.

package mockpbdataplane

import mock "github.com/stretchr/testify/mock"

// isGetEnvoyBootstrapParamsRequest_NodeSpec is an autogenerated mock type for the isGetEnvoyBootstrapParamsRequest_NodeSpec type
type isGetEnvoyBootstrapParamsRequest_NodeSpec struct {
	mock.Mock
}

type isGetEnvoyBootstrapParamsRequest_NodeSpec_Expecter struct {
	mock *mock.Mock
}

func (_m *isGetEnvoyBootstrapParamsRequest_NodeSpec) EXPECT() *isGetEnvoyBootstrapParamsRequest_NodeSpec_Expecter {
	return &isGetEnvoyBootstrapParamsRequest_NodeSpec_Expecter{mock: &_m.Mock}
}

// isGetEnvoyBootstrapParamsRequest_NodeSpec provides a mock function with given fields:
func (_m *isGetEnvoyBootstrapParamsRequest_NodeSpec) isGetEnvoyBootstrapParamsRequest_NodeSpec() {
	_m.Called()
}

// isGetEnvoyBootstrapParamsRequest_NodeSpec_isGetEnvoyBootstrapParamsRequest_NodeSpec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isGetEnvoyBootstrapParamsRequest_NodeSpec'
type isGetEnvoyBootstrapParamsRequest_NodeSpec_isGetEnvoyBootstrapParamsRequest_NodeSpec_Call struct {
	*mock.Call
}

// isGetEnvoyBootstrapParamsRequest_NodeSpec is a helper method to define mock.On call
func (_e *isGetEnvoyBootstrapParamsRequest_NodeSpec_Expecter) isGetEnvoyBootstrapParamsRequest_NodeSpec() *isGetEnvoyBootstrapParamsRequest_NodeSpec_isGetEnvoyBootstrapParamsRequest_NodeSpec_Call {
	return &isGetEnvoyBootstrapParamsRequest_NodeSpec_isGetEnvoyBootstrapParamsRequest_NodeSpec_Call{Call: _e.mock.On("isGetEnvoyBootstrapParamsRequest_NodeSpec")}
}

func (_c *isGetEnvoyBootstrapParamsRequest_NodeSpec_isGetEnvoyBootstrapParamsRequest_NodeSpec_Call) Run(run func()) *isGetEnvoyBootstrapParamsRequest_NodeSpec_isGetEnvoyBootstrapParamsRequest_NodeSpec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *isGetEnvoyBootstrapParamsRequest_NodeSpec_isGetEnvoyBootstrapParamsRequest_NodeSpec_Call) Return() *isGetEnvoyBootstrapParamsRequest_NodeSpec_isGetEnvoyBootstrapParamsRequest_NodeSpec_Call {
	_c.Call.Return()
	return _c
}

func (_c *isGetEnvoyBootstrapParamsRequest_NodeSpec_isGetEnvoyBootstrapParamsRequest_NodeSpec_Call) RunAndReturn(run func()) *isGetEnvoyBootstrapParamsRequest_NodeSpec_isGetEnvoyBootstrapParamsRequest_NodeSpec_Call {
	_c.Call.Return(run)
	return _c
}

// newIsGetEnvoyBootstrapParamsRequest_NodeSpec creates a new instance of isGetEnvoyBootstrapParamsRequest_NodeSpec. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newIsGetEnvoyBootstrapParamsRequest_NodeSpec(t interface {
	mock.TestingT
	Cleanup(func())
}) *isGetEnvoyBootstrapParamsRequest_NodeSpec {
	mock := &isGetEnvoyBootstrapParamsRequest_NodeSpec{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
