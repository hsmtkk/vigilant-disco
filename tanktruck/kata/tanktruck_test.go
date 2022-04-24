package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/tanktruck/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 2940, kata.TankVol(5, 7, 3848))
	assert.Equal(t, 907, kata.TankVol(2, 7, 3848))
	assert.Equal(t, 1750, kata.TankVol(3, 6, 3500))
	assert.Equal(t, 1750, kata.TankVol(3, 6, 3501))
}

func Test1(t *testing.T) {
	assert.Equal(t, 907, kata.TankVol(2, 7, 3848))
}

func Test2(t *testing.T) {
	assert.Equal(t, 1750, kata.TankVol(3, 6, 3500))
}

func Test3(t *testing.T) {
	assert.Equal(t, 1750, kata.TankVol(3, 6, 3501))
}
