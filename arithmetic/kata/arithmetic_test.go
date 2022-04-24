package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/arithmetic/kata"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 10, kata.Arithmetic(8, 2, "add"))
}

func TestSub(t *testing.T) {
	assert.Equal(t, 6, kata.Arithmetic(8, 2, "subtract"))
}

func TestMul(t *testing.T) {
	assert.Equal(t, 16, kata.Arithmetic(8, 2, "multiply"))
}

func TestDiv(t *testing.T) {
	assert.Equal(t, 4, kata.Arithmetic(8, 2, "divide"))
}
