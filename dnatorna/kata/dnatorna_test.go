package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/dnatorna/kata"
	"github.com/stretchr/testify/assert"
)

func TestDNAtoRNA(t *testing.T) {
	assert.Equal(t, "GCAU", kata.DNAtoRNA("GCAT"))
}
