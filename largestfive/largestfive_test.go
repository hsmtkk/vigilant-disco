package largestfive_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/largestfive"
	"github.com/stretchr/testify/assert"
)

func TestSolve0(t *testing.T) {
	var want uint = 67890
	got := largestfive.Solve("1234567890")
	assert.Equal(t, want, got)
}

func TestSequences(t *testing.T) {
	want := []uint{12345, 23456, 34567, 45678, 56789, 67890}
	got := largestfive.Sequences("1234567890")
	assert.Equal(t, want, got)
}

func TestSolve1(t *testing.T) {
	var want uint = 74765
	got := largestfive.Solve("731674765")
	assert.Equal(t, want, got)
}
