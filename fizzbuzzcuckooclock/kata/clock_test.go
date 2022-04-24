package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/fizzbuzzcuckooclock/kata"
	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	inputWants := map[string]string{
		"13:34": "tick",
		"21:00": "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo",
		"11:15": "Fizz Buzz",
		"03:03": "Fizz",
		"14:30": "Cuckoo",
		"08:55": "Buzz",
		"00:00": "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo",
		"12:00": "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo",
	}
	for input, want := range inputWants {
		got := kata.FizzBuzzCuckooClock(input)
		assert.Equal(t, want, got)
	}
}
