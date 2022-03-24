package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/ballupward/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 10, kata.MaxBall(37))
}

func Test1(t *testing.T) {
	assert.Equal(t, 13, kata.MaxBall(45))
}

func Test2(t *testing.T) {
	assert.Equal(t, 28, kata.MaxBall(99))
}

func Test3(t *testing.T) {
	assert.Equal(t, 24, kata.MaxBall(85))
}

func Test4(t *testing.T) {
	assert.Equal(t, 39, kata.MaxBall(136))
}

func Test5(t *testing.T) {
	assert.Equal(t, 4, kata.MaxBall(15))
}
