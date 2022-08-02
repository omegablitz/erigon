package iavl

import (
	"fmt"
	"strings"
)

type RangeProof struct {
	// You don't need the right path because
	// it can be derived from what we have.
	LeftPath   PathToLeaf      `json:"left_path"`
	InnerNodes []PathToLeaf    `json:"inner_nodes"`
	Leaves     []proofLeafNode `json:"leaves"`

	// memoize
	rootVerified bool
	rootHash     []byte // valid iff rootVerified is true
	treeEnd      bool   // valid iff rootVerified is true

}

// Keys returns all the keys in the RangeProof.  NOTE: The keys here may
// include more keys than provided by tree.GetRangeWithProof or
// MutableTree.GetVersionedRangeWithProof.  The keys returned there are only
// in the provided [startKey,endKey){limit} range.  The keys returned here may
// include extra keys, such as:
// - the key before startKey if startKey is provided and doesn't exist;
// - the key after a queried key with tree.GetWithProof, when the key is absent.
func (proof *RangeProof) Keys() (keys [][]byte) {
	if proof == nil {
		return nil
	}
	for _, leaf := range proof.Leaves {
		keys = append(keys, leaf.Key)
	}
	return keys
}

// String returns a string representation of the proof.
func (proof *RangeProof) String() string {
	if proof == nil {
		return "<nil-RangeProof>"
	}
	return proof.StringIndented("")
}

func (proof *RangeProof) StringIndented(indent string) string {
	istrs := make([]string, 0, len(proof.InnerNodes))
	for _, ptl := range proof.InnerNodes {
		istrs = append(istrs, ptl.stringIndented(indent+"    "))
	}
	lstrs := make([]string, 0, len(proof.Leaves))
	for _, leaf := range proof.Leaves {
		lstrs = append(lstrs, leaf.stringIndented(indent+"    "))
	}
	return fmt.Sprintf(`RangeProof{
%s  LeftPath: %v
%s  InnerNodes:
%s    %v
%s  Leaves:
%s    %v
%s  (rootVerified): %v
%s  (rootHash): %X
%s  (treeEnd): %v
%s}`,
		indent, proof.LeftPath.stringIndented(indent+"  "),
		indent,
		indent, strings.Join(istrs, "\n"+indent+"    "),
		indent,
		indent, strings.Join(lstrs, "\n"+indent+"    "),
		indent, proof.rootVerified,
		indent, proof.rootHash,
		indent, proof.treeEnd,
		indent)
}

// The index of the first leaf (of the whole tree).
// Returns -1 if the proof is nil.
func (proof *RangeProof) LeftIndex() int64 {
	if proof == nil {
		return -1
	}
	return proof.LeftPath.Index()
}

// Also see LeftIndex().
// Verify that a key has some value.
// Does not assume that the proof itself is valid, call Verify() first.
func (proof *RangeProof) VerifyItem(key, value []byte) error {
	panic("REMOVED")
}

// Verify that proof is valid absence proof for key.
// Does not assume that the proof itself is valid.
// For that, use Verify(root).
func (proof *RangeProof) VerifyAbsence(key []byte) error {
	panic("REMOVED")
}

// Verify that proof is valid.
func (proof *RangeProof) Verify(root []byte) error {
	panic("REMOVED")
}

func (proof *RangeProof) verify(root []byte) (err error) {
	panic("REMOVED")
}

// ComputeRootHash computes the root hash with leaves.
// Returns nil if error or proof is nil.
// Does not verify the root hash.
func (proof *RangeProof) ComputeRootHash() []byte {
	if proof == nil {
		return nil
	}
	rootHash, _ := proof.computeRootHash()
	return rootHash
}

func (proof *RangeProof) computeRootHash() (rootHash []byte, err error) {
	rootHash, treeEnd, err := proof._computeRootHash()
	if err == nil {
		proof.rootHash = rootHash // memoize
		proof.treeEnd = treeEnd   // memoize
	}
	return rootHash, err
}

func (proof *RangeProof) _computeRootHash() (rootHash []byte, treeEnd bool, err error) {
	panic("REMOVED")
}

///////////////////////////////////////////////////////////////////////////////
