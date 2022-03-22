package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/sumdig/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 7, kata.DigitalRoot(16))
}

func Test1(t *testing.T) {
	assert.Equal(t, 6, kata.DigitalRoot(195))
}

func Test2(t *testing.T) {
	assert.Equal(t, 2, kata.DigitalRoot(992))
}

func Test3(t *testing.T) {
	assert.Equal(t, 9, kata.DigitalRoot(167346))
}

func Test4(t *testing.T) {
	assert.Equal(t, 0, kata.DigitalRoot(0))
}
