package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/equabletriangle/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.True(t, kata.EquableTriangle(5, 12, 13))
}

func Test1(t *testing.T) {
	assert.False(t, kata.EquableTriangle(2, 3, 4))
}
