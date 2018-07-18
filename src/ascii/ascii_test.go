package ascii

import (
	"testing"
)

//TestInvalidASCII Given a bad character, see if it says the input is bad
func TestInvalidASCII(t *testing.T) {
	input := "EL NIÑO"
	output := Invalid(input)

	if output != true {
		t.Error("TestInvalidASCII test failed")
	}
}

//TestAsterisks Given a bad character, see if it says the input is bad
func TestAsterisks(t *testing.T) {
	input := "cens*red"
	output := Invalid(input)

	if output != true {
		t.Error("Asterisk test failed")
	}
}

//TestValid tests legitimate input
func TestValid(t *testing.T) {
	input := "abcdefg"
	output := Invalid(input)

	if output == true {
		t.Error("Enya test failed")
	}
}
