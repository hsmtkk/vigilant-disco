package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/gocinema/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 43, kata.Movie(500, 15, 0.9))
}

func Test1(t *testing.T) {
	assert.Equal(t, 2, kata.Movie(0, 10, 0.95))
}
