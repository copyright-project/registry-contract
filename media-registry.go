package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

// Structure of the state
// [instagram_id: string]: {
//	  imageUrl: string
//	  postUrl: string
//	  postedAt: string
//	  ownerId: string
// }

var SYSTEM = sdk.Export(_init)
var PUBLIC = sdk.Export(registerMedia)

func _init() {}

func registerMedia(instagramID, metadata string) {
	key := []byte(instagramID)
	if state.ReadString(key) != "" {
		panic("The record already exists")
	}
	state.WriteString(key, metadata)
}
