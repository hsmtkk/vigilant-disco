package validpin_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/validpin"
	"github.com/stretchr/testify/assert"
)

func TestInvalidLength(t *testing.T) {
	pins := []string{"1", "12", "123", "12345", "1234567", "-1234", "1.234", "-1.234", "00000000"}
	for _, pin := range pins {
		got := validpin.Validate(pin)
		assert.False(t, got)
	}
}

func TestNonDigit(t *testing.T) {
	pins := []string{"a234", ".234"}
	for _, pin := range pins {
		got := validpin.Validate(pin)
		assert.False(t, got)
	}
}

func TestValid(t *testing.T) {
	pins := []string{"0000", "1111", "123456", "098765", "000000", "123456", "090909"}
	for _, pin := range pins {
		got := validpin.Validate(pin)
		assert.True(t, got)
	}
}
