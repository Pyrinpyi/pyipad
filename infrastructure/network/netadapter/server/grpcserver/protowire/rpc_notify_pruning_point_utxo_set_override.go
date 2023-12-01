package protowire

import (
	"github.com/Pyrinpyi/pyipad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PyipadMessage_NotifyPruningPointUTXOSetOverrideRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_NotifyPruningPointUTXOSetOverrideRequest is nil")
	}
	return &appmessage.NotifyPruningPointUTXOSetOverrideRequestMessage{}, nil
}

func (x *PyipadMessage_NotifyPruningPointUTXOSetOverrideRequest) fromAppMessage(_ *appmessage.NotifyPruningPointUTXOSetOverrideRequestMessage) error {
	x.NotifyPruningPointUTXOSetOverrideRequest = &NotifyPruningPointUTXOSetOverrideRequestMessage{}
	return nil
}

func (x *PyipadMessage_NotifyPruningPointUTXOSetOverrideResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_NotifyPruningPointUTXOSetOverrideResponse is nil")
	}
	return x.NotifyPruningPointUTXOSetOverrideResponse.toAppMessage()
}

func (x *PyipadMessage_NotifyPruningPointUTXOSetOverrideResponse) fromAppMessage(message *appmessage.NotifyPruningPointUTXOSetOverrideResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.NotifyPruningPointUTXOSetOverrideResponse = &NotifyPruningPointUTXOSetOverrideResponseMessage{
		Error: err,
	}
	return nil
}

func (x *NotifyPruningPointUTXOSetOverrideResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NotifyPruningPointUTXOSetOverrideResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.NotifyPruningPointUTXOSetOverrideResponseMessage{
		Error: rpcErr,
	}, nil
}

func (x *PyipadMessage_PruningPointUTXOSetOverrideNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_PruningPointUTXOSetOverrideNotification is nil")
	}
	return &appmessage.PruningPointUTXOSetOverrideNotificationMessage{}, nil
}

func (x *PyipadMessage_PruningPointUTXOSetOverrideNotification) fromAppMessage(_ *appmessage.PruningPointUTXOSetOverrideNotificationMessage) error {
	x.PruningPointUTXOSetOverrideNotification = &PruningPointUTXOSetOverrideNotificationMessage{}
	return nil
}

func (x *PyipadMessage_StopNotifyingPruningPointUTXOSetOverrideRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_StopNotifyingPruningPointUTXOSetOverrideRequest is nil")
	}
	return &appmessage.StopNotifyingPruningPointUTXOSetOverrideRequestMessage{}, nil
}

func (x *PyipadMessage_StopNotifyingPruningPointUTXOSetOverrideRequest) fromAppMessage(_ *appmessage.StopNotifyingPruningPointUTXOSetOverrideRequestMessage) error {
	x.StopNotifyingPruningPointUTXOSetOverrideRequest = &StopNotifyingPruningPointUTXOSetOverrideRequestMessage{}
	return nil
}

func (x *PyipadMessage_StopNotifyingPruningPointUTXOSetOverrideResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_StopNotifyingPruningPointUTXOSetOverrideResponse is nil")
	}
	return x.StopNotifyingPruningPointUTXOSetOverrideResponse.toAppMessage()
}

func (x *PyipadMessage_StopNotifyingPruningPointUTXOSetOverrideResponse) fromAppMessage(
	message *appmessage.StopNotifyingPruningPointUTXOSetOverrideResponseMessage) error {

	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.StopNotifyingPruningPointUTXOSetOverrideResponse = &StopNotifyingPruningPointUTXOSetOverrideResponseMessage{
		Error: err,
	}
	return nil
}

func (x *StopNotifyingPruningPointUTXOSetOverrideResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "StopNotifyingPruningPointUTXOSetOverrideResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.StopNotifyingPruningPointUTXOSetOverrideResponseMessage{
		Error: rpcErr,
	}, nil
}
