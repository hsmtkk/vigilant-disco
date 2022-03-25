package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/playpassphrase/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	got := kata.PlayPass("I LOVE YOU!!!", 1)
	want := "!!!vPz fWpM J"
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	got := kata.PlayPass("I LOVE YOU!!!", 0)
	want := "!!!uOy eVoL I"
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	got := kata.PlayPass("AAABBCCY", 1)
	want := "zDdCcBbB"
	assert.Equal(t, want, got)
}
