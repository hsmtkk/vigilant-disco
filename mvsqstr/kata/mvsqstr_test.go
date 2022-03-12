package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/mvsqstr/kata"
	"github.com/stretchr/testify/assert"
)

func TestVert(t *testing.T) {
	input := "hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu"
	got := kata.Oper(kata.VertMirror, input)
	want := "QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw"
	assert.Equal(t, want, got)
}

func TestHor(t *testing.T) {
	input := "lVHt\nJVhv\nCSbg\nyeCt"
	got := kata.Oper(kata.HorMirror, input)
	want := "yeCt\nCSbg\nJVhv\nlVHt"
	assert.Equal(t, want, got)
}
