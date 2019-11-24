package main

import (
	"bytes"

	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/events"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var SYSTEM = sdk.Export(_init)
var PUBLIC = sdk.Export(registerMedia)
var EVENTS = sdk.Export(Log)

var OWNER_KEY = []byte("__CONTRACT_OWNER__")

func _init() {
	state.WriteBytes(OWNER_KEY, address.GetSignerAddress())
}

func Log(mediaID, phash, copyrights, timestamp, imageURL string) {}

func registerMedia(mediaID, phash, copyrights, timestamp, imageURL string) {
	if !bytes.Equal(state.ReadBytes(OWNER_KEY), address.GetSignerAddress()) {
		panic("Only contract owner can register media")
	}
	events.EmitEvent(Log, mediaID, phash, copyrights, timestamp, imageURL)
}
