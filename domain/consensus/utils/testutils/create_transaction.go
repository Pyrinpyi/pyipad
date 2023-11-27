package testutils

import (
	"github.com/Pyrinpyi/pyipad/domain/consensus/model/externalapi"
	"github.com/Pyrinpyi/pyipad/domain/consensus/utils/consensushashing"
	"github.com/Pyrinpyi/pyipad/domain/consensus/utils/constants"
	"github.com/Pyrinpyi/pyipad/domain/consensus/utils/txscript"
)

// CreateTransaction create a transaction that spends the first output of provided transaction.
// Assumes that the output being spent has opTrueScript as it's scriptPublicKey
// Creates the value of the spent output minus 1 leor
func CreateTransaction(txToSpend *externalapi.DomainTransaction, fee uint64) (*externalapi.DomainTransaction, error) {
	scriptPublicKey, redeemScript := OpTrueScript()

	signatureScript, err := txscript.PayToScriptHashSignatureScript(redeemScript, nil)
	if err != nil {
		return nil, err
	}
	input := &externalapi.DomainTransactionInput{
		PreviousOutpoint: externalapi.DomainOutpoint{
			TransactionID: *consensushashing.TransactionID(txToSpend),
			Index:         0,
		},
		SignatureScript: signatureScript,
		Sequence:        constants.MaxTxInSequenceNum,
	}
	output := &externalapi.DomainTransactionOutput{
		ScriptPublicKey: scriptPublicKey,
		Value:           txToSpend.Outputs[0].Value - fee,
	}
	return &externalapi.DomainTransaction{
		Version: constants.MaxTransactionVersion,
		Inputs:  []*externalapi.DomainTransactionInput{input},
		Outputs: []*externalapi.DomainTransactionOutput{output},
		Payload: []byte{},
	}, nil
}
