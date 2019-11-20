package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"strings"

	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var SYSTEM = sdk.Export(_init)
var PUBLIC = sdk.Export(registerMedia, areRegistered, getMedia)

var OWNER_KEY = []byte("__CONTRACT_OWNER__")
var INDEX_KEY = []byte("__INDEX__")

func _init() {
	state.WriteBytes(OWNER_KEY, address.GetSignerAddress())
}

func compress(data string) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(data)); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	return b.Bytes()
}

func decompress(data []byte) string {
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	return string(s)
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

func registerMedia(mediaID, metadata string) {
	if !bytes.Equal(state.ReadBytes(OWNER_KEY), address.GetSignerAddress()) {
		panic("Only contract owner can register media")
	}
	key := []byte(mediaID)
	if isRegistered(mediaID) {
		panic("The record already exists")
	}
	compressedData := compress(metadata)
	state.WriteBytes(key, compressedData)
	state.WriteString(INDEX_KEY, "|"+state.ReadString(INDEX_KEY))
}

func getMedia(id string) string {
	return decompress(state.ReadBytes([]byte(id)))
}

func main() {}
