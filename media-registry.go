package main

import (
	"bytes"
	"strings"

	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var SYSTEM = sdk.Export(_init)
var PUBLIC = sdk.Export(registerMedia, areRegistered, getMedia)

var OWNER_KEY = []byte("__CONTRACT_OWNER__")
var PHASHES_KEY = []byte("__PHASHES__")

func _init() {
	state.WriteBytes(OWNER_KEY, address.GetSignerAddress())
}

func isRegistered(id string) bool {
	key := []byte(id)
	return state.ReadString(key) != ""
}

func areRegistered(ids string) string {
	res := ""
	for _, id := range strings.Split(ids, ",") {
		if isRegistered(id) {
			res = res + "1"
		} else {
			res = res + "0"
		}
	}
	return res
}

func registerMedia(mediaID, pHash, metadata string) {
	if len(mediaID) == 0 {
		panic("Media ID must be provided")
	}
	if len(pHash) == 0 {
		panic("PHash must be provided")
	}
	if !bytes.Equal(state.ReadBytes(OWNER_KEY), address.GetSignerAddress()) {
		panic("Only contract owner can register media")
	}
	key := []byte(mediaID)
	if isRegistered(mediaID) {
		panic("The record already exists")
	}
	state.WriteString(key, metadata)
	state.WriteString(PHASHES_KEY, ","+pHash)
}

func getMedia(id string) string {
	return state.ReadString([]byte(id))
}

func main() {}
