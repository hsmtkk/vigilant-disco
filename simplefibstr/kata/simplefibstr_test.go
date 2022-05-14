package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/simplefibstr/kata"
	"github.com/stretchr/testify/assert"
)

func TestSimpleFibStr(t *testing.T) {
	testCases := map[int]string{
		0: "0",
		1: "01",
		2: "010",
		3: "01001",
		5: "0100101001001",
	}
	for input, want := range testCases {
		got := kata.Solve(input)
		assert.Equal(t, want, got)
	}
}
