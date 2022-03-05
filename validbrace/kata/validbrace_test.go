package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/validbrace/kata"
	"github.com/stretchr/testify/assert"
)

func TestValidBraces(t *testing.T) {
	assert.True(t, kata.ValidBraces("(){}[]"))
	assert.True(t, kata.ValidBraces("([{}])"))
	assert.False(t, kata.ValidBraces("(}"))
	assert.False(t, kata.ValidBraces("[(])"))
	assert.False(t, kata.ValidBraces("[({)](]"))
	assert.False(t, kata.ValidBraces("(("))
	assert.False(t, kata.ValidBraces("]}"))
}
