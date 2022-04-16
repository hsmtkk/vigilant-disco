package findmiddle_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/findmiddle"
	"github.com/stretchr/testify/assert"
)

func TestFindMiddle(t *testing.T) {
	assert.Equal(t, 0, findmiddle.FindMiddle([]int{2, 3, 1}))
	assert.Equal(t, 0, findmiddle.FindMiddle([]int{-2, -3, -1}))
	assert.Equal(t, 1, findmiddle.FindMiddle([]int{5, 10, 14}))
}
