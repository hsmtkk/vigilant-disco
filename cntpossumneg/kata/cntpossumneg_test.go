package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/cntpossumneg/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	res := []int{10, -65}
	assert.Equal(t, res, kata.CountPositivesSumNegatives(arr))
}
