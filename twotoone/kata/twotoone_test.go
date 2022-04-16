package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/twotoone/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, "aehrsty", kata.TwoToOne("aretheyhere", "yestheyarehere"))
}

func Test1(t *testing.T) {
	assert.Equal(t, "abcdefghilnoprstu", kata.TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding"))
}
