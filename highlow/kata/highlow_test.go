package kata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/hsmtkk/vigilant-disco/highlow/kata"
)

func Test0(t *testing.T) {
	input := "8 3 -5 42 -1 0 0 -9 4 7 4 -4"
	got := kata.HighAndLow(input)
	want := "42 -9"
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	input := "1 2 3"
	got := kata.HighAndLow(input)
	want := "3 1"
	assert.Equal(t, want, got)
}

func TestSingleInput(t *testing.T) {
	input := "42"
	got := kata.HighAndLow(input)
	want := "42 42"
	assert.Equal(t, want, got)
}
