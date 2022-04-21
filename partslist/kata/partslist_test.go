package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/partslist/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := "(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)"
	got := kata.PartList([]string{"I", "wish", "I", "hadn't", "come"})
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := "(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)"
	got := kata.PartList([]string{"cdIw", "tzIy", "xDu", "rThG"})
	assert.Equal(t, want, got)
}
