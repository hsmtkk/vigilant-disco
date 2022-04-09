package findscreensize_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/findscreensize"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, "1024x768", findscreensize.FindScreenSize(1024, "4:3"))
}

func Test1(t *testing.T) {
	assert.Equal(t, "1280x720", findscreensize.FindScreenSize(1280, "16:9"))
}

func Test2(t *testing.T) {
	assert.Equal(t, "3840x1080", findscreensize.FindScreenSize(3840, "32:9"))
}
