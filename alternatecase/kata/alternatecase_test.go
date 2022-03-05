package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/alternatecase/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	orig := "String.prototype.toAlternatingCase"
	want := "sTRING.PROTOTYPE.TOaLTERNATINGcASE"
	got := kata.ToAlternatingCase(orig)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	orig := "altERnaTIng cAsE"
	want := "ALTerNAtiNG CaSe"
	got := kata.ToAlternatingCase(orig)
	assert.Equal(t, want, got)
}
