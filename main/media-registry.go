package main

import (
	"bytes"
	"strings"

	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var SYSTEM = sdk.Export(_init)
var PUBLIC = sdk.Export(registerMedia, getMedia)

var OWNER_KEY = []byte("__CONTRACT_OWNER__")
var PHASHES_KEY = []byte("__PHASHES__")

func _init() {
	state.WriteBytes(OWNER_KEY, address.GetSignerAddress())
}

func isRegistered(pHash, binaryHash string) bool {
	key := []byte(pHash)
	records := state.ReadString(key)
	return strings.Contains(records, binaryHash)
}

func isValidURL(url string) bool {
	if len(url) == 0 {
		return false
	}
	if !strings.HasPrefix(url, "http") {
		return false
	}
	return true
}

func validateInput(pHash, imageURL, postedAt, copyrights, binaryHash string) {
	if len(pHash) != 64 {
		panic("Invalid phash: length should be 64")
	}
	if !isValidURL(imageURL) {
		panic("Image URL is invalid")
	}
	if len(postedAt) == 0 {
		panic("Image timestamp is not provided")
	}
	if len(binaryHash) == 0 {
		panic("File image hash is not provided")
	}
}

func registerMedia(pHash, imageURL, postedAt, copyrights, binaryHash string) {
	if !bytes.Equal(state.ReadBytes(OWNER_KEY), address.GetSignerAddress()) {
		panic("Only contract owner can register media")
	}

	validateInput(pHash, imageURL, postedAt, copyrights, binaryHash)

	if isRegistered(pHash, binaryHash) {
		panic("Record with the following url already exists " + imageURL)
	}

	record := strings.Join([]string{imageURL, postedAt, copyrights, binaryHash}, ",")

	key := []byte(pHash)
	state.WriteString(key, state.ReadString(key)+"|"+record)

	if keys := state.ReadString(PHASHES_KEY); !strings.Contains(keys, pHash) {
		state.WriteString(PHASHES_KEY, keys+"|"+pHash)
	}
}

func getMedia(pHash string) []string {
	records := strings.TrimLeft(state.ReadString([]byte(pHash)), "|")
	return strings.Split(records, "|")
}

func main() {}
