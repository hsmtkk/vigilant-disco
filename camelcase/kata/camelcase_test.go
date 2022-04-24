package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/camelcase/kata"
	"github.com/stretchr/testify/assert"
)

func TestCamelCase(t *testing.T) {
	testCases := map[string]string{
		"test case":         "TestCase",
		"camel case method": "CamelCaseMethod",
		"say hello ":        "SayHello",
		" camel case word":  "CamelCaseWord",
		"":                  "",
	}
	for input, want := range testCases {
		got := kata.CamelCase(input)
		assert.Equal(t, want, got)
	}
}
