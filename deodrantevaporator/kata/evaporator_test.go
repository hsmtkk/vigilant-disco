package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/deodrantevaporator/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 22, kata.Evaporator(10, 10, 10))
}

func Test1(t *testing.T) {
	assert.Equal(t, 29, kata.Evaporator(10, 10, 5))
}

func Test2(t *testing.T) {
	assert.Equal(t, 59, kata.Evaporator(100, 5, 5))
}
