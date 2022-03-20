package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/happybirthday/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 162, kata.WrapPresent(17, 32, 11))
}

func Test1(t *testing.T) {
	assert.Equal(t, 124, kata.WrapPresent(13, 13, 13))
}

func Test2(t *testing.T) {
	assert.Equal(t, 32, kata.WrapPresent(1, 3, 1))
}
