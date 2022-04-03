package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/twiceold/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	fatherAge int
	childAge  int
	want      int
}

func TestTwiceAsOld(t *testing.T) {
	testCases := []testCase{
		{36, 7, 22},
		{55, 30, 5},
		{42, 21, 0},
		{22, 1, 20},
		{29, 0, 29},
	}
	for _, testCase := range testCases {
		got := kata.TwiceAsOld(testCase.fatherAge, testCase.childAge)
		assert.Equal(t, testCase.want, got)
	}
}
