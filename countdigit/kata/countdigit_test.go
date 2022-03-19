package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/countdigit/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := 4
	got := kata.NbDig(10, 1)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := 11
	got := kata.NbDig(25, 1)
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	want := 213
	got := kata.NbDig(550, 5)
	assert.Equal(t, want, got)
}

func Test3(t *testing.T) {
	want := 4700
	got := kata.NbDig(5750, 0)
	assert.Equal(t, want, got)
}
