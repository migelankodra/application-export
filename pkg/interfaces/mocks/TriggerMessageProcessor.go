// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/migelankodra/application-export/pkg/interfaces"
	mock "github.com/stretchr/testify/mock"

	types "github.com/edgexfoundry/go-mod-messaging/v2/pkg/types"
)

// TriggerMessageProcessor is an autogenerated mock type for the TriggerMessageProcessor type
type TriggerMessageProcessor struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, envelope
func (_m *TriggerMessageProcessor) Execute(ctx interfaces.AppFunctionContext, envelope types.MessageEnvelope) error {
	ret := _m.Called(ctx, envelope)

	var r0 error
	if rf, ok := ret.Get(0).(func(interfaces.AppFunctionContext, types.MessageEnvelope) error); ok {
		r0 = rf(ctx, envelope)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
