package listconverter

import (
	"fmt"
	"os"
	"testing"
)

//TestConvertToMap testes the ConvertToSlice method given example input from the root of the repo
func TestConvertToMap(t *testing.T) {
	input, _ := os.Open("../../example-input.txt")
	output, _ := ConvertToMap(input)
	fmt.Println(output)

	if output != nil {
		t.Error("Map test failed")
	}
}

//TestConvertToSlice testes the ConvertToSlice method given example input from the root of the repo
func TestConvertToSlice(t *testing.T) {
	input, _ := os.Open("../../example-input.txt")
	output, _ := ConvertToSlice(input)

	fmt.Println(output)

	if output != nil {
		t.Error("Slice test failed")
	}
}
