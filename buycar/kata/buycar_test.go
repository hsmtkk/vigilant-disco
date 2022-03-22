package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/buycar/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := [2]int{6, 766}
	got := kata.NbMonths(2000, 8000, 1000, 1.5)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := [2]int{0, 4000}
	got := kata.NbMonths(12000, 8000, 1000, 1.5)
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	want := [2]int{8, 597}
	got := kata.NbMonths(8000, 12000, 500, 1.0)
	assert.Equal(t, want, got)
}

func Test3(t *testing.T) {
	want := [2]int{8, 332}
	got := kata.NbMonths(18000, 32000, 1500, 1.25)
	assert.Equal(t, want, got)
}

func Test4(t *testing.T) {
	want := [2]int{25, 122}
	got := kata.NbMonths(7500, 32000, 300, 1.55)
	assert.Equal(t, want, got)
}

func Test5(t *testing.T) {
	want := [2]int{1, 110}
	got := kata.NbMonths(2600, 3500, 1000, 1.2)
	assert.Equal(t, want, got)
}
