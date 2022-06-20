package pairofint_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/pairofint"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := []pairofint.Pair{{2, 2}, {2, 3}, {2, 4}, {3, 3}, {3, 4}, {4, 4}}
	got := pairofint.GeneratePairs(2, 4)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := []pairofint.Pair{{0, 0}, {0, 1}, {1, 1}}
	got := pairofint.GeneratePairs(0, 1)
	assert.Equal(t, want, got)
}
func Test2(t *testing.T) {
	want := []pairofint.Pair{{0, 0}}
	got := pairofint.GeneratePairs(0, 0)
	assert.Equal(t, want, got)
}
