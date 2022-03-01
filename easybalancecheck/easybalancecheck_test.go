package easybalancecheck_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hsmtkk/vigilant-disco/easybalancecheck"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	content, err := fileAsString("book0.csv")
	assert.Nil(t, err)
	book, err := easybalancecheck.NewBookParser().ParseBook(content)
	assert.Nil(t, err)
	got := easybalancecheck.NewReportMaker().MakeReport(book)
	want, err := fileAsString("report0.txt")
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func fileAsString(file string) (string, error) {
	bs, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file; %s; %w", file, err)
	}
	return string(bs), nil
}
