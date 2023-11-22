package protowire

import (
	"github.com/Pyrinpyi/pyipad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PyipadMessage_NotifyVirtualSelectedParentChainChangedRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_NotifyVirtualSelectedParentChainChangedRequest is nil")
	}
	return &appmessage.NotifyVirtualSelectedParentChainChangedRequestMessage{
		IncludeAcceptedTransactionIDs: x.NotifyVirtualSelectedParentChainChangedRequest.IncludeAcceptedTransactionIds,
	}, nil
}

func (x *PyipadMessage_NotifyVirtualSelectedParentChainChangedRequest) fromAppMessage(appmessage *appmessage.NotifyVirtualSelectedParentChainChangedRequestMessage) error {
	x.NotifyVirtualSelectedParentChainChangedRequest = &NotifyVirtualSelectedParentChainChangedRequestMessage{
		IncludeAcceptedTransactionIds: appmessage.IncludeAcceptedTransactionIDs,
	}
	return nil
}

func (x *PyipadMessage_NotifyVirtualSelectedParentChainChangedResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_NotifyVirtualSelectedParentChainChangedResponse is nil")
	}
	return x.NotifyVirtualSelectedParentChainChangedResponse.toAppMessage()
}

func (x *PyipadMessage_NotifyVirtualSelectedParentChainChangedResponse) fromAppMessage(message *appmessage.NotifyVirtualSelectedParentChainChangedResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.NotifyVirtualSelectedParentChainChangedResponse = &NotifyVirtualSelectedParentChainChangedResponseMessage{
		Error: err,
	}
	return nil
}

func (x *NotifyVirtualSelectedParentChainChangedResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NotifyVirtualSelectedParentChainChangedResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.NotifyVirtualSelectedParentChainChangedResponseMessage{
		Error: rpcErr,
	}, nil
}

func (x *PyipadMessage_VirtualSelectedParentChainChangedNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_VirtualSelectedParentChainChangedNotification is nil")
	}
	return x.VirtualSelectedParentChainChangedNotification.toAppMessage()
}

func (x *PyipadMessage_VirtualSelectedParentChainChangedNotification) fromAppMessage(message *appmessage.VirtualSelectedParentChainChangedNotificationMessage) error {
	x.VirtualSelectedParentChainChangedNotification = &VirtualSelectedParentChainChangedNotificationMessage{
		RemovedChainBlockHashes: message.RemovedChainBlockHashes,
		AddedChainBlockHashes:   message.AddedChainBlockHashes,
		AcceptedTransactionIds:  make([]*AcceptedTransactionIds, len(message.AcceptedTransactionIDs)),
	}

	for i, acceptedTransactionIDs := range message.AcceptedTransactionIDs {
		x.VirtualSelectedParentChainChangedNotification.AcceptedTransactionIds[i] = &AcceptedTransactionIds{}
		x.VirtualSelectedParentChainChangedNotification.AcceptedTransactionIds[i].fromAppMessage(acceptedTransactionIDs)
	}
	return nil
}

func (x *VirtualSelectedParentChainChangedNotificationMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "VirtualSelectedParentChainChangedNotificationMessage is nil")
	}
	message := &appmessage.VirtualSelectedParentChainChangedNotificationMessage{
		RemovedChainBlockHashes: x.RemovedChainBlockHashes,
		AddedChainBlockHashes:   x.AddedChainBlockHashes,
		AcceptedTransactionIDs:  make([]*appmessage.AcceptedTransactionIDs, len(x.AcceptedTransactionIds)),
	}

	for i, acceptedTransactionIds := range x.AcceptedTransactionIds {
		message.AcceptedTransactionIDs[i] = acceptedTransactionIds.toAppMessage()
	}
	return message, nil
}
