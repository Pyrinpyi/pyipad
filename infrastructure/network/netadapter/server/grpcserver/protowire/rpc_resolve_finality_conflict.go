package protowire

import (
	"github.com/Pyrinpyi/pyipad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PyipadMessage_ResolveFinalityConflictRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_ResolveFinalityConflictRequest is nil")
	}
	return x.ResolveFinalityConflictRequest.toAppMessage()
}

func (x *PyipadMessage_ResolveFinalityConflictRequest) fromAppMessage(message *appmessage.ResolveFinalityConflictRequestMessage) error {
	x.ResolveFinalityConflictRequest = &ResolveFinalityConflictRequestMessage{
		FinalityBlockHash: message.FinalityBlockHash,
	}
	return nil
}

func (x *ResolveFinalityConflictRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "ResolveFinalityConflictRequestMessage is nil")
	}
	return &appmessage.ResolveFinalityConflictRequestMessage{
		FinalityBlockHash: x.FinalityBlockHash,
	}, nil
}

func (x *PyipadMessage_ResolveFinalityConflictResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_ResolveFinalityConflictResponse is nil")
	}
	return x.ResolveFinalityConflictResponse.toAppMessage()
}

func (x *PyipadMessage_ResolveFinalityConflictResponse) fromAppMessage(message *appmessage.ResolveFinalityConflictResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.ResolveFinalityConflictResponse = &ResolveFinalityConflictResponseMessage{
		Error: err,
	}
	return nil
}

func (x *ResolveFinalityConflictResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "ResolveFinalityConflictResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.ResolveFinalityConflictResponseMessage{
		Error: rpcErr,
	}, nil
}
