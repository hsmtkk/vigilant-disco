package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/whicharein/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	got := kata.InArray([]string{"live", "arp", "strong"}, []string{"lively", "alive", "harp", "sharp", "armstrong"})
	want := []string{"arp", "live", "strong"}
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	got := kata.InArray([]string{"cod", "code", "wars", "ewar", "ar"}, []string{})
	want := []string{}
	assert.Equal(t, want, got)
}
