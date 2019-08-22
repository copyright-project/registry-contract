package main

import (
	"bytes"

	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

// Structure of the state
// [media_id: string]: {
//	  imageUrl: string
//	  postUrl: string
//	  postedAt: string
//	  ownerId: string
//	  hash: string
// }

var SYSTEM = sdk.Export(_init)
var PUBLIC = sdk.Export(registerMedia)

var OWNER_KEY = []byte("__CONTRACT_OWNER__")

func _init() {
	state.WriteBytes(OWNER_KEY, address.GetCallerAddress())
}

func registerMedia(mediaID, metadata string) {
	if !bytes.Equal(state.ReadBytes(OWNER_KEY), address.GetCallerAddress()) {
		panic("Only contract owner can register media")
	}
	key := []byte(mediaID)
	if state.ReadString(key) != "" {
		panic("The record already exists")
	}
	state.WriteString(key, metadata)
}

func main() {}
