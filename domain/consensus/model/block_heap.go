package model

import "github.com/Pyrinpyi/pyipad/domain/consensus/model/externalapi"

// BlockHeap represents a heap of block hashes, providing a priority-queue functionality
type BlockHeap interface {
	Push(blockHash *externalapi.DomainHash) error
	PushSlice(blockHash []*externalapi.DomainHash) error
	Pop() *externalapi.DomainHash
	Len() int
	ToSlice() []*externalapi.DomainHash
}
