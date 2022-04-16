package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/jadencase/kata"
	"github.com/stretchr/testify/assert"
)

func TestJadenCa(t *testing.T) {
	assert.Equal(t, "Most Trees Are Blue", kata.ToJadenCase("most trees are blue"))
}

/*
   Expect(ToJadenCase("most trees are blue")).To(Equal("Most Trees Are Blue"))
   Expect(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own.")).To(Equal("All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."))
   Expect(ToJadenCase("When I die. then you will realize")).To(Equal("When I Die. Then You Will Realize"))
   Expect(ToJadenCase("Jonah Hill is a genius")).To(Equal("Jonah Hill Is A Genius"))
   Expect(ToJadenCase("Dying is mainstream")).To(Equal("Dying Is Mainstream"))

*/
