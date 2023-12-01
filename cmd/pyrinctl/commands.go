package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Pyrinpyi/pyipad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.PyipadMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.PyipadMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.PyipadMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.PyipadMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.PyipadMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.PyipadMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.PyipadMessage_BanRequest{}),
	reflect.TypeOf(protowire.PyipadMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
