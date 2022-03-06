package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/bookseller/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	stocks := []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	categories := []string{"A", "B", "C", "D"}
	want := "(A : 0) - (B : 1290) - (C : 515) - (D : 600)"
	got := kata.StockList(stocks, categories)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	stocks := []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
	categories := []string{"A", "B"}
	want := "(A : 200) - (B : 1140)"
	got := kata.StockList(stocks, categories)
	assert.Equal(t, want, got)
}

/*
 [ROXANNE 102 RHODODE 123 BKWRKAA 125 BTSQZFG 239 DRTYMKH 060] [B R D X]
*/

func TestCategoryStockFromString(t *testing.T) {
	stocks := []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	want := []kata.CategoryStock{
		{"B", 150},
		{"C", 515},
		{"B", 250},
		{"B", 890},
		{"D", 600},
	}
	got := []kata.CategoryStock{}
	for _, stock := range stocks {
		cs := kata.CategoryStockFromString(stock)
		got = append(got, cs)
	}
	assert.Equal(t, want, got)
}
