package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/glowingplant/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	upSpeed       int
	downSpeed     int
	desiredHeight int
	want          int
}

func Test0(t *testing.T) {
	testCases := []testCase{
		{100, 10, 910, 10},
		{10, 9, 4, 1},
		{5, 2, 5, 1},
		{5, 2, 6, 2},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, kata.GrowingPlant(testCase.upSpeed, testCase.downSpeed, testCase.desiredHeight))
	}
}
