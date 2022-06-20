package lengthtwovals_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/lengthtwovals"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := []string{"true", "false", "true", "false", "true"}
	got := lengthtwovals.Alternate(5, "true", "false")
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := []string{"blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue",
		"red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"}
	got := lengthtwovals.Alternate(20, "blue", "red")
	assert.Equal(t, want, got)
}
