package lightclient

import (
	"github.com/tendermint/go-amino"
)

type Codec = amino.Codec

var Cdc *Codec

func init() {
	cdc := amino.NewCodec()
	Cdc = cdc.Seal()
}
