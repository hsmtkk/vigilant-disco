package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/sumevenfib/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 10, kata.SumEvenFibonacci(8))
}

func Test1(t *testing.T) {
	assert.Equal(t, 60696, kata.SumEvenFibonacci(111111))
}

func Test2(t *testing.T) {
	assert.Equal(t, 4613732, kata.SumEvenFibonacci(8675309))
}
