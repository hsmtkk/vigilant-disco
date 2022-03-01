package cannonready_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/cannonready"
	"github.com/stretchr/testify/assert"
)

const AYE = "aye"
const NAY = "nay"

func Test0(t *testing.T) {
	gunners := []cannonready.Gunner{
		{"Mike", AYE},
		{"Joe", AYE},
		{"Johnson", AYE},
		{"Peter", AYE},
	}
	want := "Fire!"
	got := cannonready.CannonReady(gunners)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	gunners := []cannonready.Gunner{
		{"Mike", AYE},
		{"Joe", NAY},
		{"Johnson", AYE},
		{"Peter", AYE},
	}
	want := "Shiver me timbers!"
	got := cannonready.CannonReady(gunners)
	assert.Equal(t, want, got)
}
