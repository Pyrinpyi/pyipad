package p2p

import (
	"github.com/daglabs/btcd/config"
	"github.com/daglabs/btcd/peer"
	"github.com/daglabs/btcd/wire"
	"time"
)

// OnAddr is invoked when a peer receives an addr bitcoin message and is
// used to notify the server about advertised addresses.
func (sp *Peer) OnAddr(_ *peer.Peer, msg *wire.MsgAddr) {
	// Ignore addresses when running on the simulation test network.  This
	// helps prevent the network from becoming another public test network
	// since it will not be able to learn about other peers that have not
	// specifically been provided.
	if config.MainConfig().SimNet {
		return
	}

	// A message that has no addresses is invalid.
	if len(msg.AddrList) == 0 {
		peerLog.Errorf("Command [%s] from %s does not contain any addresses",
			msg.Command(), sp.Peer)
		sp.Disconnect()
		return
	}

	if msg.IncludeAllSubnetworks {
		peerLog.Errorf("Got unexpected IncludeAllSubnetworks=true in [%s] command from %s",
			msg.Command(), sp.Peer)
		sp.Disconnect()
		return
	} else if !msg.SubnetworkID.IsEqual(config.MainConfig().SubnetworkID) && msg.SubnetworkID != nil {
		peerLog.Errorf("Only full nodes and %s subnetwork IDs are allowed in [%s] command, but got subnetwork ID %s from %s",
			config.MainConfig().SubnetworkID, msg.Command(), msg.SubnetworkID, sp.Peer)
		sp.Disconnect()
		return
	}

	for _, na := range msg.AddrList {
		// Don't add more address if we're disconnecting.
		if !sp.Connected() {
			return
		}

		// Set the timestamp to 5 days ago if it's more than 24 hours
		// in the future so this address is one of the first to be
		// removed when space is needed.
		now := time.Now()
		if na.Timestamp.After(now.Add(time.Minute * 10)) {
			na.Timestamp = now.Add(-1 * time.Hour * 24 * 5)
		}

		// Add address to known addresses for this peer.
		sp.addKnownAddresses([]*wire.NetAddress{na})
	}

	// Add addresses to server address manager.  The address manager handles
	// the details of things such as preventing duplicate addresses, max
	// addresses, and last seen updates.
	// XXX bitcoind gives a 2 hour time penalty here, do we want to do the
	// same?
	sp.server.addrManager.AddAddresses(msg.AddrList, sp.NA(), msg.SubnetworkID)
}