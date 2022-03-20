package dscorder_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/dscorder"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input uint
	want  uint
}

func Test0(t *testing.T) {
	testCases := []testCase{
		{0, 0},
		{1, 1},
		{15, 51},
		{1021, 2110},
		{123456789, 987654321},
		{145263, 654321},
		{1254859723, 9875543221},
	}
	for _, testCase := range testCases {
		got := dscorder.DscOrder(testCase.input)
		assert.Equal(t, testCase.want, got)
	}
}
