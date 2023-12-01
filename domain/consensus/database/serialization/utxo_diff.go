package serialization

import (
	"github.com/Pyrinpyi/pyipad/domain/consensus/model/externalapi"
	"github.com/Pyrinpyi/pyipad/domain/consensus/utils/utxo"
)

// UTXODiffToDBUTXODiff converts UTXODiff to DbUtxoDiff
func UTXODiffToDBUTXODiff(diff externalapi.UTXODiff) (*DbUtxoDiff, error) {
	toAdd, err := utxoCollectionToDBUTXOCollection(diff.ToAdd())
	if err != nil {
		return nil, err
	}

	toRemove, err := utxoCollectionToDBUTXOCollection(diff.ToRemove())
	if err != nil {
		return nil, err
	}

	return &DbUtxoDiff{
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}, nil
}

// DBUTXODiffToUTXODiff converts DbUtxoDiff to UTXODiff
func DBUTXODiffToUTXODiff(diff *DbUtxoDiff) (externalapi.UTXODiff, error) {
	toAdd, err := dbUTXOCollectionToUTXOCollection(diff.ToAdd)
	if err != nil {
		return nil, err
	}

	toRemove, err := dbUTXOCollectionToUTXOCollection(diff.ToRemove)
	if err != nil {
		return nil, err
	}

	return utxo.NewUTXODiffFromCollections(toAdd, toRemove)
}
