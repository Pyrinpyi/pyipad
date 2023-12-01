package testutils

import (
	"github.com/Pyrinpyi/pyipad/domain/consensus/model/externalapi"
	"github.com/Pyrinpyi/pyipad/domain/consensus/utils/constants"
	"github.com/Pyrinpyi/pyipad/domain/consensus/utils/txscript"
	"github.com/pkg/errors"
)

// OpTrueScript returns a P2SH script paying to an anyone-can-spend address,
// The second return value is a redeemScript to be used with txscript.PayToScriptHashSignatureScript
func OpTrueScript() (*externalapi.ScriptPublicKey, []byte) {
	var err error
	redeemScript := []byte{txscript.OpTrue}
	scriptPublicKeyScript, err := txscript.PayToScriptHashScript(redeemScript)
	if err != nil {
		panic(errors.Wrapf(err, "Couldn't parse opTrueScript. This should never happen"))
	}
	scriptPublicKey := &externalapi.ScriptPublicKey{Script: scriptPublicKeyScript, Version: constants.MaxScriptPublicKeyVersion}
	return scriptPublicKey, redeemScript
}
