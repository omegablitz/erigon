package lightclient

import (
	"bytes"
	"fmt"

	"github.com/tendermint/tendermint/crypto/merkle"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	chainIDLength              uint64 = 32
	heightLength               uint64 = 8
	validatorSetHashLength     uint64 = 32
	validatorPubkeyLength      uint64 = 32
	validatorVotingPowerLength uint64 = 8
	appHashLength              uint64 = 32
	storeNameLengthBytesLength uint64 = 32
	keyLengthBytesLength       uint64 = 32
	valueLengthBytesLength     uint64 = 32
	maxConsensusStateLength    uint64 = 32 * (128 - 1) // maximum validator quantity 99
)

type ConsensusState struct {
	ChainID             string
	Height              uint64
	AppHash             []byte
	CurValidatorSetHash []byte
	NextValidatorSet    *tmtypes.ValidatorSet
}

// input:
// | chainID   | height   | appHash  | curValidatorSetHash | [{validator pubkey, voting power}] |
// | 32 bytes  | 8 bytes  | 32 bytes | 32 bytes            | [{32 bytes, 8 bytes}]              |
func DecodeConsensusState(input []byte) (ConsensusState, error) {
	panic("REMOVED")
}

// output:
// | chainID   | height   | appHash  | curValidatorSetHash | [{validator pubkey, voting power}] |
// | 32 bytes  | 8 bytes  | 32 bytes | 32 bytes            | [{32 bytes, 8 bytes}]              |
func (cs ConsensusState) EncodeConsensusState() ([]byte, error) {
	panic("REMOVED")
}

func (cs *ConsensusState) ApplyHeader(header *Header) (bool, error) {
	panic("REMOVED")
}

type Header struct {
	tmtypes.SignedHeader
	ValidatorSet     *tmtypes.ValidatorSet `json:"validator_set"`
	NextValidatorSet *tmtypes.ValidatorSet `json:"next_validator_set"`
}

func (h *Header) Validate(chainID string) error {
	if err := h.SignedHeader.ValidateBasic(chainID); err != nil {
		return err
	}
	if h.ValidatorSet == nil {
		return fmt.Errorf("invalid header: validator set is nil")
	}
	if h.NextValidatorSet == nil {
		return fmt.Errorf("invalid header: next validator set is nil")
	}
	if !bytes.Equal(h.ValidatorsHash, h.ValidatorSet.Hash()) {
		return fmt.Errorf("invalid header: validator set does not match hash")
	}
	if !bytes.Equal(h.NextValidatorsHash, h.NextValidatorSet.Hash()) {
		return fmt.Errorf("invalid header: next validator set does not match hash")
	}
	return nil
}

func (h *Header) EncodeHeader() ([]byte, error) {
	bz, err := Cdc.MarshalBinaryLengthPrefixed(h)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func DecodeHeader(input []byte) (*Header, error) {
	var header Header
	err := Cdc.UnmarshalBinaryLengthPrefixed(input, &header)
	if err != nil {
		return nil, err
	}
	return &header, nil
}

type KeyValueMerkleProof struct {
	Key       []byte
	Value     []byte
	StoreName string
	AppHash   []byte
	Proof     *merkle.Proof
}

func (kvmp *KeyValueMerkleProof) Validate() bool {
	panic("REMOVED")
}

// input:
// | storeName | key length | key | value length | value | appHash  | proof |
// | 32 bytes  | 32 bytes   |     | 32 bytes     |       | 32 bytes |       |
func DecodeKeyValueMerkleProof(input []byte) (*KeyValueMerkleProof, error) {
	panic("REMOVED")
}
