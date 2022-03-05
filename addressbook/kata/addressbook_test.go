package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/addressbook/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	const input = `John Daggett, 341 King Road, Plymouth MA
	Alice Ford, 22 East Broadway, Richmond VA
	Orville Thomas, 11345 Oak Bridge Road, Tulsa OK
	Terry Kalkas, 402 Lans Road, Beaver Falls PA
	Eric Adams, 20 Post Road, Sudbury MA
	Hubert Sims, 328A Brook Road, Roanoke VA
	Amy Wilde, 334 Bayshore Pkwy, Mountain View CA
	Sal Carpenter, 73 6th Street, Boston MA`
	const want = "California\n..... Amy Wilde 334 Bayshore Pkwy Mountain View California\n Massachusetts\n..... Eric Adams 20 Post Road Sudbury Massachusetts\n..... John Daggett 341 King Road Plymouth Massachusetts\n..... Sal Carpenter 73 6th Street Boston Massachusetts\n Oklahoma\n..... Orville Thomas 11345 Oak Bridge Road Tulsa Oklahoma\n Pennsylvania\n..... Terry Kalkas 402 Lans Road Beaver Falls Pennsylvania\n Virginia\n..... Alice Ford 22 East Broadway Richmond Virginia\n..... Hubert Sims 328A Brook Road Roanoke Virginia"
	got := kata.ByState(input)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	const input = `John Daggett, 341 King Road, Plymouth MA
	Alice Ford, 22 East Broadway, Richmond VA
	Sal Carpenter, 73 6th Street, Boston MA`
	const want = "Massachusetts\n..... John Daggett 341 King Road Plymouth Massachusetts\n..... Sal Carpenter 73 6th Street Boston Massachusetts\n Virginia\n..... Alice Ford 22 East Broadway Richmond Virginia"
	got := kata.ByState(input)
	assert.Equal(t, want, got)
}

func TestRecordsFromCSV(t *testing.T) {
	const input = `John Daggett, 341 King Road, Plymouth MA
	Alice Ford, 22 East Broadway, Richmond VA
	Sal Carpenter, 73 6th Street, Boston MA`
	want := kata.Records{
		{"John Daggett", "341 King Road", "Plymouth", "Massachusetts"},
		{"Alice Ford", "22 East Broadway", "Richmond", "Virginia"},
		{"Sal Carpenter", "73 6th Street", "Boston", "Massachusetts"},
	}
	got, err := kata.RecordsFromCSV(input)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestStates(t *testing.T) {
	const input = `John Daggett, 341 King Road, Plymouth MA
	Alice Ford, 22 East Broadway, Richmond VA
	Sal Carpenter, 73 6th Street, Boston MA`
	records, err := kata.RecordsFromCSV(input)
	assert.Nil(t, err)
	got := records.States()
	want := []string{"Massachusetts", "Virginia"}
	assert.Equal(t, want, got)
}
