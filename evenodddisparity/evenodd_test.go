package evenodddisparity_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/evenodddisparity"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	got := evenodddisparity.EvenOddDisparity([]string{"0", "1", "2", "3"})
	assert.Equal(t, 0, got)
}

func Test1(t *testing.T) {
	got := evenodddisparity.EvenOddDisparity([]string{"0", "1", "2", "3", "a", "b"})
	assert.Equal(t, 0, got)
}

func Test2(t *testing.T) {
	got := evenodddisparity.EvenOddDisparity([]string{"0", "15", "z", "16", "m", "13", "14", "c", "9", "10", "13", "u", "4", "3"})
	assert.Equal(t, 0, got)
}

func Test3(t *testing.T) {
	got := evenodddisparity.EvenOddDisparity([]string{"5", "15", "16", "10", "6", "4", "16", "t", "13", "n", "14", "k", "n", "0", "q", "d", "7", "9"})
	assert.Equal(t, 2, got)
}
