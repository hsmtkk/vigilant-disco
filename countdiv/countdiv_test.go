package countdiv_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/countdiv"
	"github.com/stretchr/testify/assert"
)

func TestNumOfDivisors(t *testing.T) {
	testCases := map[int]int{
		1:    1,
		4:    3,
		5:    2,
		12:   6,
		25:   3,
		4096: 13,
	}
	for input, want := range testCases {
		got := countdiv.NumOfDivisors(input)
		assert.Equal(t, want, got)
	}
}
