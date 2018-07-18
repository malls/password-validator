package listconverter

import (
	"os"
	"testing"
)

func getFile() *os.File {
	testFile, _ := os.Open("list-converter_test.txt")

	return testFile
}

//TestConvertToMap testes the ConvertToSlice method given example input from the root of the repo
func TestConvertToMap(t *testing.T) {
	input := getFile()
	output, _ := ConvertToMap(input, 2, 8)

	expected := map[string]bool{"first": true, "second": true, "third": true}

	if output["first"] != expected["first"] {
		t.Error("First present value map test failed")
	}

	if output["second"] != expected["second"] {
		t.Error("Second present value map test failed")
	}

	if output["a"] != expected["a"] {
		t.Error("Short map test failed")
	}

	if output["hundredth"] != expected["hundredth"] {
		t.Error("Short map test failed")
	}
}

//TestConvertToSlice testes the ConvertToSlice method given example input from the root of the repo
func TestConvertToSlice(t *testing.T) {
	input := getFile()
	output, _ := ConvertToSlice(input)

	expected := []string{"first", "second", "third", "a", "hundredth"}

	for i := range expected {
		if output[i] != expected[i] {
			t.Error("Slice test failed")
		}
	}
}
