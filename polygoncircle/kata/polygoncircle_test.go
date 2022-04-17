package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/polygoncircle/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	radius  float64
	polygon int
	want    float64
}

func Test0(t *testing.T) {
	testCases := []testCase{
		{3, 3, 11.691},
		{5.8, 7, 92.053},
		{4, 5, 38.042},
	}
	for _, testCase := range testCases {
		got := kata.AreaOfPolygonInsideCircle(testCase.radius, testCase.polygon)
		assert.Equal(t, testCase.want, got)
	}
}
