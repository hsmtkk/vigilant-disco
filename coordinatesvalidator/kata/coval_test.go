package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/coordinatesvalidator/kata"
	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	validCoordinates := []string{
		"-23, 25",
		"4, -3",
		"24.53525235, 23.45235",
		"04, -23.234235",
		"43.91343345, 143"}
	for _, input := range validCoordinates {
		assert.True(t, kata.IsValidCoordinates(input), input)
	}
}

func TestInvalid(t *testing.T) {
	invalidCoordinates := []string{
		"23.234, - 23.4234",
		"2342.43536, 34.324236",
		"N23.43345, E32.6457",
		"99.234, 12.324",
		"6.325624, 43.34345.345",
		"0, 1,2",
		"0.342q0832, 1.2324",
		"23.245, 1e1"}
	for _, input := range invalidCoordinates {
		assert.False(t, kata.IsValidCoordinates(input), input)
	}
}
