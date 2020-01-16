package main

import (
	"math/rand"
	"testing"

	. "github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateStringOfLength(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestPanicContractCalledNotByOwner(t *testing.T) {
	InServiceScope(nil, []byte("some-signer"), func(m Mockery) {
		require.Panics(t, func() {
			pHash := generateStringOfLength(64)
			binaryHash := generateStringOfLength(64)
			registerMedia(pHash, "https://some-url", "123456789", "by me", binaryHash)
		})
	})
}

func TestPanicWithInvalidArguments(t *testing.T) {
	InServiceScope(nil, []byte("some-signer"), func(m Mockery) {
		require.Panics(t, func() {
			_init()
			binaryHash := generateStringOfLength(64)
			registerMedia("", "https://some-url", "123456789", "by me", binaryHash)
		})
		require.Panics(t, func() {
			_init()
			pHash := generateStringOfLength(64)
			binaryHash := generateStringOfLength(64)
			registerMedia(pHash, "some-url", "123456789", "by me", binaryHash)
		})
		require.Panics(t, func() {
			_init()
			pHash := generateStringOfLength(64)
			binaryHash := generateStringOfLength(64)
			registerMedia(pHash, "", "123456789", "by me", binaryHash)
		})
	})
}

func TestPanicRegisterExistingRecord(t *testing.T) {
	InServiceScope(nil, []byte("some-signer"), func(m Mockery) {
		require.Panics(t, func() {
			_init()
			pHash := generateStringOfLength(64)
			binaryHash := generateStringOfLength(64)
			registerMedia(pHash, "https://some-url", "123456789", "by me", binaryHash)
			registerMedia(pHash, "https://some-url", "123456789", "by me", binaryHash)
		})
	})
}

func TestRecordRetrieval(t *testing.T) {
	InServiceScope(nil, nil, func(m Mockery) {
		_init()
		pHash := generateStringOfLength(64)
		binaryHash := generateStringOfLength(64)
		m.MockEmitEvent(mediaRegistered, pHash)
		registerMedia(pHash, "https://some-url", "123456789", "by me", binaryHash)
		require.Equal(t, []string{"https://some-url,123456789,by me," + binaryHash}, getMedia(pHash))
	})
}
func TestRegisterRecordWithSamePHash(t *testing.T) {
	InServiceScope(nil, nil, func(m Mockery) {
		_init()
		pHash := generateStringOfLength(64)
		binaryHash1 := generateStringOfLength(64)
		binaryHash2 := generateStringOfLength(64)
		m.MockEmitEvent(mediaRegistered, pHash)
		registerMedia(pHash, "https://some-url-1", "123456789", "by me", binaryHash1)
		registerMedia(pHash, "https://some-url-2", "123456789", "by me", binaryHash2)
		require.Equal(t, []string{
			"https://some-url-1,123456789,by me," + binaryHash1,
			"https://some-url-2,123456789,by me," + binaryHash2,
		}, getMedia(pHash))
	})
}
