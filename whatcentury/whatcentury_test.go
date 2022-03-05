package whatcentury_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/whatcentury"
	"github.com/stretchr/testify/assert"
)

type yearCentury struct {
	Year    uint32
	Century string
}

func TestWhatCentury(t *testing.T) {
	yearCenturies := []yearCentury{
		{1999, "20th"},
		{2011, "21st"},
		{2154, "22nd"},
		{2259, "23rd"},
		{1234, "13th"},
		{1023, "11th"},
		{2000, "20th"},
		{3210, "33rd"},
	}
	for _, yC := range yearCenturies {
		got := whatcentury.WhatCentury(yC.Year)
		assert.Equal(t, yC.Century, got)
	}
}
