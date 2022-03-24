package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/tortoiserace/kata"
	"github.com/stretchr/testify/assert"
)

func TestRace0(t *testing.T) {
	want := [3]int{0, 32, 18}
	got := kata.Race(720, 850, 70)
	assert.Equal(t, want, got)
}

func TestRace1(t *testing.T) {
	want := [3]int{-1, -1, -1}
	got := kata.Race(820, 81, 550)
	assert.Equal(t, want, got)
}

func TestRace2(t *testing.T) {
	want := [3]int{3, 21, 49}
	got := kata.Race(80, 91, 37)
	assert.Equal(t, want, got)
}
