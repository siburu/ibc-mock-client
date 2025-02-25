package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
)

// CheckForMisbehaviour never detects misbehaviour and always returns false.
func (cs ClientState) CheckForMisbehaviour(_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.ClientMessage) bool {
	return false
}
