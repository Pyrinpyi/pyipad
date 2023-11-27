package rpchandlers

import (
	"github.com/Pyrinpyi/pyipad/app/appmessage"
	"github.com/Pyrinpyi/pyipad/app/rpc/rpccontext"
	"github.com/Pyrinpyi/pyipad/infrastructure/network/netadapter/router"
)

// HandleGetVirtualSelectedParentBlueScore handles the respectively named RPC command
func HandleGetVirtualSelectedParentBlueScore(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	c := context.Domain.Consensus()
	selectedParent, err := c.GetVirtualSelectedParent()
	if err != nil {
		return nil, err
	}
	blockInfo, err := c.GetBlockInfo(selectedParent)
	if err != nil {
		return nil, err
	}
	return appmessage.NewGetVirtualSelectedParentBlueScoreResponseMessage(blockInfo.BlueScore), nil
}
