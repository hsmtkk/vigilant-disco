package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/abbrevname/kata"
	"github.com/stretchr/testify/assert"
)

func TestAbbrevName(t *testing.T) {
	nameWant := map[string]string{
		"Sam Harris":     "S.H",
		"Patrick Feenan": "P.F",
		"Evan Cole":      "E.C",
		"P Favuzzi":      "P.F",
		"David Mendieta": "D.M",
	}
	for name, want := range nameWant {
		got := kata.AbbrevName(name)
		assert.Equal(t, want, got)
	}
}
