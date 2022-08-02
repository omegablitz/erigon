package iavl

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto/merkle"
)

const ProofOpIAVLValue = "iavl:v"

// IAVLValueOp takes a key and a single value as argument and
// produces the root hash.
//
// If the produced root hash matches the expected hash, the proof
// is good.
type IAVLValueOp struct {
	// Encoded in ProofOp.Key.
	key []byte

	// To encode in ProofOp.Data.
	// Proof is nil for an empty tree.
	// The hash of an empty tree is nil.
	Proof *RangeProof `json:"proof"`
}

var _ merkle.ProofOperator = IAVLValueOp{}

func NewIAVLValueOp(key []byte, proof *RangeProof) IAVLValueOp {
	return IAVLValueOp{
		key:   key,
		Proof: proof,
	}
}

func IAVLValueOpDecoder(pop merkle.ProofOp) (merkle.ProofOperator, error) {
	panic("REMOVED")
}

func (op IAVLValueOp) ProofOp() merkle.ProofOp {
	bz := cdc.MustMarshalBinaryLengthPrefixed(op)
	return merkle.ProofOp{
		Type: ProofOpIAVLValue,
		Key:  op.key,
		Data: bz,
	}
}

func (op IAVLValueOp) String() string {
	return fmt.Sprintf("IAVLValueOp{%v}", op.GetKey())
}

func (op IAVLValueOp) Run(args [][]byte) ([][]byte, error) {
	panic("REMOVED")
}

func (op IAVLValueOp) GetKey() []byte {
	return op.key
}
