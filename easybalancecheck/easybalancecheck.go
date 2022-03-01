package easybalancecheck

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

type ReportMaker interface {
	MakeReport(book Book) string
}

type BookParser interface {
	ParseBook(content string) (Book, error)
}

type Record struct {
	CheckNumber uint
	Category    string
	Amount      float64
}

func NewRecord(elems []string) (Record, error) {
	if len(elems) != 3 {
		return Record{}, fmt.Errorf("incorrect line; %v", elems)
	}
	checkNumber, err := strconv.ParseUint(elems[0], 10, 64)
	if err != nil {
		return Record{}, fmt.Errorf("failed to parse check number; %s; %w", elems[0], err)
	}
	amount, err := strconv.ParseFloat(elems[2], 64)
	if err != nil {
		return Record{}, fmt.Errorf("failed to parse amount; %s; %w", elems[2], err)
	}
	return Record{CheckNumber: uint(checkNumber), Category: elems[1], Amount: amount}, nil
}

type Book struct {
	OriginalBalance float64
	Records         []Record
}

func NewReportMaker() ReportMaker {
	return &reportMakerImpl{}
}

type reportMakerImpl struct{}

func (r *reportMakerImpl) MakeReport(book Book) string {
	balance := book.OriginalBalance
	totalExpense := 0.0
	lines := []string{}
	lines = append(lines, fmt.Sprintf("Original Balance: %.2f", book.OriginalBalance))
	for _, record := range book.Records {
		balance -= record.Amount
		totalExpense += record.Amount
		lines = append(lines, fmt.Sprintf("%d %s %.2f Balance %.2f", record.CheckNumber, record.Category, record.Amount, balance))
	}
	averageExpense := totalExpense / float64(len(book.Records))
	lines = append(lines, fmt.Sprintf("Total expense  %.2f", totalExpense))
	lines = append(lines, fmt.Sprintf("Average expense  %.2f", averageExpense))
	return strings.Join(lines, "\n") + "\n"
}

func NewBookParser() BookParser {
	return &bookParserImpl{}
}

type bookParserImpl struct{}

func (b *bookParserImpl) ParseBook(content string) (Book, error) {
	reader := csv.NewReader(strings.NewReader(content))
	lines, err := reader.ReadAll()
	if err != nil {
		return Book{}, fmt.Errorf("failed to read CSV; %w", err)
	}
	originalBalance, err := strconv.ParseFloat(lines[0][0], 64)
	if err != nil {
		return Book{}, fmt.Errorf("failed to read original balance; %s; %w", lines[0], err)
	}
	records := []Record{}
	for _, line := range lines[1:] {
		record, err := NewRecord(line)
		if err != nil {
			return Book{}, fmt.Errorf("failed to make record; %s; %w", line, err)
		}
		records = append(records, record)
	}
	return Book{OriginalBalance: originalBalance, Records: records}, nil
}
