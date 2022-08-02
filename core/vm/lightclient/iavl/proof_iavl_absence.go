package iavl

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto/merkle"
	proto "github.com/tendermint/tendermint/proto/tendermint/crypto"
)

const ProofOpIAVLAbsence = "iavl:a"

// IAVLAbsenceOp takes a key as its only argument
//
// If the produced root hash matches the expected hash, the proof
// is good.
type IAVLAbsenceOp struct {
	// Encoded in ProofOp.Key.
	key []byte

	// To encode in ProofOp.Data.
	// Proof is nil for an empty tree.
	// The hash of an empty tree is nil.
	Proof *RangeProof `json:"proof"`
}

var _ merkle.ProofOperator = IAVLAbsenceOp{}

func NewIAVLAbsenceOp(key []byte, proof *RangeProof) IAVLAbsenceOp {
	return IAVLAbsenceOp{
		key:   key,
		Proof: proof,
	}
}

func IAVLAbsenceOpDecoder(pop proto.ProofOp) (merkle.ProofOperator, error) {
	panic("REMOVED")
}

func (op IAVLAbsenceOp) ProofOp() proto.ProofOp {
	bz := cdc.MustMarshalBinaryLengthPrefixed(op)
	return proto.ProofOp{
		Type: ProofOpIAVLAbsence,
		Key:  op.key,
		Data: bz,
	}
}

func (op IAVLAbsenceOp) String() string {
	return fmt.Sprintf("IAVLAbsenceOp{%v}", op.GetKey())
}

func (op IAVLAbsenceOp) Run(args [][]byte) ([][]byte, error) {
	panic("REMOVED")
}

func (op IAVLAbsenceOp) GetKey() []byte {
	return op.key
}
