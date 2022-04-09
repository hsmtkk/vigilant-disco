package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/squarepi/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 8, kata.SquarePi(5))
}

func Test1(t *testing.T) {
	assert.Equal(t, 15, kata.SquarePi(10))
}
