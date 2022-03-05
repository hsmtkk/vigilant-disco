package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/primenum/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := "(2**2)(3**3)(5)(7)(11**2)(17)"
	got := kata.PrimeFactors(7775460)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := "(2**2)(5)(3967)"
	got := kata.PrimeFactors(79340)
	assert.Equal(t, want, got)
}

func TestPrimes0(t *testing.T) {
	want := []int{2, 2, 5, 3967}
	got := kata.Primes(79340)
	assert.Equal(t, want, got)
}

func TestPrimes1(t *testing.T) {
	want := []int{2}
	got := kata.Primes(2)
	assert.Equal(t, want, got)
}

func TestPrimes2(t *testing.T) {
	want := []int{2, 2}
	got := kata.Primes(4)
	assert.Equal(t, want, got)
}

func TestPrimes3(t *testing.T) {
	want := []int{2, 3}
	got := kata.Primes(6)
	assert.Equal(t, want, got)
}

/*

import (
	. "codewarrior/kata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(n int, exp string) {
    var ans = PrimeFactors(n)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("ConvertFracts", func() {
   It("Basic tests", func() {
        dotest(7775460, "(2**2)(3**3)(5)(7)(11**2)(17)")
        dotest(79340, "")

   })
})
*/
