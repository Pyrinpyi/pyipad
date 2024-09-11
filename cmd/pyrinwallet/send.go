package main

import (
	"context"
	"fmt"
	"github.com/Pyrinpyi/pyipad/cmd/pyrinwallet/utils"
	"os"
	"strings"

	"github.com/Pyrinpyi/pyipad/cmd/pyrinwallet/daemon/client"
	"github.com/Pyrinpyi/pyipad/cmd/pyrinwallet/daemon/pb"
	"github.com/Pyrinpyi/pyipad/cmd/pyrinwallet/keys"
	"github.com/Pyrinpyi/pyipad/cmd/pyrinwallet/libpyrinwallet"
	"github.com/pkg/errors"
)

func send(conf *sendConfig) error {
	keysFile, err := keys.ReadKeysFile(conf.NetParams(), conf.KeysFile)
	if err != nil {
		return err
	}

	if len(keysFile.ExtendedPublicKeys) > len(keysFile.EncryptedMnemonics) {
		return errors.Errorf("Cannot use 'send' command for multisig wallet without all of the keys")
	}

	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()

	var sendAmountLeor uint64
	if !conf.IsSendAll {
		sendAmountLeor, err = utils.PyiToLeor(conf.SendAmount)

		if err != nil {
			return err
		}
	}

	createUnsignedTransactionsResponse, err :=
		daemonClient.CreateUnsignedTransactions(ctx, &pb.CreateUnsignedTransactionsRequest{
			From:                     conf.FromAddresses,
			Address:                  conf.ToAddress,
			Amount:                   sendAmountLeor,
			IsSendAll:                conf.IsSendAll,
			UseExistingChangeAddress: conf.UseExistingChangeAddress,
		})
	if err != nil {
		return err
	}

	if len(conf.Password) == 0 {
		conf.Password = keys.GetPassword("Password:")
	}
	mnemonics, err := keysFile.DecryptMnemonics(conf.Password)
	if err != nil {
		if strings.Contains(err.Error(), "message authentication failed") {
			fmt.Fprintf(os.Stderr, "Password decryption failed. Sometimes this is a result of not "+
				"specifying the same keys file used by the wallet daemon process.\n")
		}
		return err
	}

	signedTransactions := make([][]byte, len(createUnsignedTransactionsResponse.UnsignedTransactions))
	for i, unsignedTransaction := range createUnsignedTransactionsResponse.UnsignedTransactions {
		signedTransaction, err := libpyrinwallet.Sign(conf.NetParams(), mnemonics, unsignedTransaction, keysFile.ECDSA)
		if err != nil {
			return err
		}
		signedTransactions[i] = signedTransaction
	}

	fmt.Printf("Broadcasting %d transaction(s)\n", len(signedTransactions))
	// Since we waited for user input when getting the password, which could take unbound amount of time -
	// create a new context for broadcast, to reset the timeout.
	broadcastCtx, broadcastCancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer broadcastCancel()

	const chunkSize = 100 // To avoid sending a message bigger than the gRPC max message size, we split it to chunks
	for offset := 0; offset < len(signedTransactions); offset += chunkSize {
		end := len(signedTransactions)
		if offset+chunkSize <= len(signedTransactions) {
			end = offset + chunkSize
		}

		chunk := signedTransactions[offset:end]
		response, err := daemonClient.Broadcast(broadcastCtx, &pb.BroadcastRequest{Transactions: chunk})
		if err != nil {
			return err
		}
		fmt.Printf("Broadcasted %d transaction(s) (broadcasted %.2f%% of the transactions so far)\n", len(chunk), 100*float64(end)/float64(len(signedTransactions)))
		fmt.Println("Broadcasted Transaction ID(s): ")
		for _, txID := range response.TxIDs {
			fmt.Printf("\t%s\n", txID)
		}
	}

	if conf.Verbose {
		fmt.Println("Serialized Transaction(s) (can be parsed via the `parse` command or resent via `broadcast`): ")
		for _, signedTx := range signedTransactions {
			fmt.Printf("\t%x\n\n", signedTx)
		}
	}

	return nil
}
