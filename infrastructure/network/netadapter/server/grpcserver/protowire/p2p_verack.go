package protowire

import (
	"github.com/Pyrinpyi/pyipad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PyipadMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *PyipadMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
