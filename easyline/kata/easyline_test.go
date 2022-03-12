package kata_test

import (
	"math/big"
	"testing"

	"github.com/hsmtkk/vigilant-disco/easyline/kata"
	"github.com/stretchr/testify/assert"
)

func Test7(t *testing.T) {
	assert.Equal(t, "3432", kata.EasyLine(7))
}

func Test13(t *testing.T) {
	assert.Equal(t, "10400600", kata.EasyLine(13))
}

func Test50(t *testing.T) {
	assert.Equal(t, "100891344545564193334812497256", kata.EasyLine(50))
}

func Test100(t *testing.T) {
	assert.Equal(t, "90548514656103281165404177077484163874504589675413336841320", kata.EasyLine(100))
}

func TestEasyLineNum(t *testing.T) {
	want := big.NewInt(10400600)
	got := kata.EasyLineNum(13)
	assert.Equal(t, want, got)
}

func TestEasyLineNumbers(t *testing.T) {
	want := []*big.Int{big.NewInt(1), big.NewInt(4), big.NewInt(6), big.NewInt(4), big.NewInt(1)}
	got := kata.EasyLineNumbers(4)
	assert.Equal(t, want, got)
}

func TestSumOfSquares(t *testing.T) {
	input := []*big.Int{big.NewInt(1), big.NewInt(4), big.NewInt(6), big.NewInt(4), big.NewInt(1)}
	want := big.NewInt(70)
	got := kata.SumOfSquares(input)
	assert.Equal(t, want, got)
}
