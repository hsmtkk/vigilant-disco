package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/removefirstlastchar/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, "loquen", kata.RemoveChar("eloquent"))
}

func Test1(t *testing.T) {
	assert.Equal(t, "ountr", kata.RemoveChar("country"))
}

func Test2(t *testing.T) {
	assert.Equal(t, "erso", kata.RemoveChar("person"))
}

func Test3(t *testing.T) {
	assert.Equal(t, "lac", kata.RemoveChar("place"))
}
