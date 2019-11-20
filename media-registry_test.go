package main

import (
	"testing"

	. "github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
)

func TestPanicContractCalledNotByOwner(t *testing.T) {
	InServiceScope(nil, []byte("some-signer"), func(m Mockery) {
		require.Panics(t, func() {
			registerMedia("test-media-id", "test-metadata")
		})
	})
}

func TestPanicWhenRegisteringSameMedia(t *testing.T) {
	InServiceScope(nil, nil, func(m Mockery) {
		_init()
		registerMedia("test-media-id-1", "test-metadata")
		require.Panics(t, func() {
			registerMedia("test-media-id-1", "test-metadata")
		})
	})
}

func TestAreRegistered(t *testing.T) {
	InServiceScope(nil, nil, func(m Mockery) {
		_init()
		registerMedia("id1", "test-metadata")
		registerMedia("id2", "test-metadata")
		require.Equal(t, areRegistered("id1,id2,id3"), "110")
	})
}

func TestRecordRetrieval(t *testing.T) {
	InServiceScope(nil, nil, func(m Mockery) {
		_init()
		registerMedia("id1", "test-metadata")
		require.Equal(t, getMedia("id1"), "test-metadata")
	})
}
