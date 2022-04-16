package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/squarenot/kata"
	"github.com/stretchr/testify/assert"
)

func TestIsSquareNumber(t *testing.T) {
	assert.Equal(t, 2, kata.Process(4))
	assert.Equal(t, 25, kata.Process(5))
}

func Test0(t *testing.T) {
	assert.Equal(t, []int{2, 9, 3, 49, 4, 1}, kata.SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1}))
}

/*
{{4, 3, 9, 7, 2, 1},{2, 9, 3, 49, 4, 1}},
{{100, 101, 5, 5, 1, 1},{10, 10201, 25, 25, 1, 1}},
{{1, 2, 3, 4, 5, 6},{1, 4, 9, 2, 25, 36}},
*/
