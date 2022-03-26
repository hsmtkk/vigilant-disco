package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/reversedstring/kata"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	assert.Equal(t, "dlrow", kata.Solution("world"))
}
