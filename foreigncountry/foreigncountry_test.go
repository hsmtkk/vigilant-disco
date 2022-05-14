package foreigncountry_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/foreigncountry"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	trips []foreigncountry.Span
	want  int
}

func TestRepresented(t *testing.T) {
	testCases := []testCase{
		{[]foreigncountry.Span{{10, 15}, {25, 35}}, 17},
		{[]foreigncountry.Span{{2, 8}, {220, 229}, {10, 16}}, 24},
		{[]foreigncountry.Span{{2, 8}, {6, 16}, {17, 26}}, 25},
		{[]foreigncountry.Span{{1, 365}}, 365},
	}
	for _, tc := range testCases {
		got := foreigncountry.DaysRepresented(tc.trips)
		assert.Equal(t, tc.want, got)
	}
}
