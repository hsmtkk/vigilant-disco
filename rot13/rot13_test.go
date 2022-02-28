package rot13_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/rot13"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := "grfg"
	got := rot13.Rotate("test")
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := "Grfg"
	got := rot13.Rotate("Test")
	assert.Equal(t, want, got)
}
