package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/mul35/kata"
	"github.com/stretchr/testify/assert"
)

func TestMul35(t *testing.T) {
	assert.Equal(t, 23, kata.Multiple3And5(10))
}
