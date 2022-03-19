package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/growthpopulation/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 15, kata.NbYear(1500, 5, 100, 5000))
}

func Test1(t *testing.T) {
	assert.Equal(t, 10, kata.NbYear(1500000, 2.5, 10000, 2000000))
}

func Test2(t *testing.T) {
	assert.Equal(t, 94, kata.NbYear(1500000, 0.25, 1000, 2000000))
}

func Test3(t *testing.T) {
	assert.Equal(t, 151, kata.NbYear(1500000, 0.25, -1000, 2000000))
}
