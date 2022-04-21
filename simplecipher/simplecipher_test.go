package simplecipher_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/simplecipher"
	"github.com/stretchr/testify/assert"
)

func TestCipher(t *testing.T) {
	cipher := simplecipher.New("abcdefghijklmnopqrstuvwxyz", "etaoinshrdlucmfwypvbgkjqxz")
	assert.Equal(t, "eta", cipher.Encode("abc"))
	assert.Equal(t, "qxz", cipher.Encode("xyz"))
	assert.Equal(t, "aeiou", cipher.Decode("eirfg"))
	assert.Equal(t, "aikcfu", cipher.Decode("erlang"))
}
