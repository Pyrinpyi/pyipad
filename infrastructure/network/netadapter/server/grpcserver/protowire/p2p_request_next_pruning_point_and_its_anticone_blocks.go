package protowire

import (
	"github.com/Pyrinpyi/pyipad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PyipadMessage_RequestNextPruningPointAndItsAnticoneBlocks) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_DonePruningPointAndItsAnticoneBlocks is nil")
	}
	return &appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks{}, nil
}

func (x *PyipadMessage_RequestNextPruningPointAndItsAnticoneBlocks) fromAppMessage(_ *appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks) error {
	return nil
}
