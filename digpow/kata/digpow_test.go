package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/digpow/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, kata.DigPow(89, 1))
}

func Test1(t *testing.T) {
	assert.Equal(t, -1, kata.DigPow(92, 1))
}
