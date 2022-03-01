package pascaltriangle_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/pascaltriangle"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	want := []uint{1}
	got := pascaltriangle.PascalTriangle(1)
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	want := []uint{1, 1, 1}
	got := pascaltriangle.PascalTriangle(2)
	assert.Equal(t, want, got)
}

func Test4(t *testing.T) {
	want := []uint{1, 1, 1, 1, 2, 1, 1, 3, 3, 1}
	got := pascaltriangle.PascalTriangle(4)
	assert.Equal(t, want, got)
}

func Test6(t *testing.T) {
	want := []uint{1, 1, 1, 1, 2, 1, 1, 3, 3, 1, 1, 4, 6, 4, 1, 1, 5, 10, 10, 5, 1}
	got := pascaltriangle.PascalTriangle(6)
	assert.Equal(t, want, got)
}
