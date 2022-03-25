package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/allinclusive/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.True(t, kata.ContainAllRots("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}))
}

func Test1(t *testing.T) {
	assert.False(t, kata.ContainAllRots("XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}))
}

func Test2(t *testing.T) {
	assert.True(t, kata.ContainAllRots("", []string{}))
}
