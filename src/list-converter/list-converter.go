package listconverter

import (
	"bufio"
	"io"
	"os"
	"strings"
)

////////////////PRIVATE METHODS////////////////////

func iterateNewlineFile(f *os.File, callback func(string)) error {
	//set up reader
	reader := bufio.NewReader(f)

	for {
		text, err := reader.ReadString('\n')
		//stop if we're at the end of the file
		if err == io.EOF {
			break
			// if there's an error, return
		} else if err != nil {
			return err
		}

		callback(strings.Replace(text, "\n", "", -1))
	}
	return nil
}

//////////////PUBLIC METHODS/////////////////

//ConvertToMap takes a newline delineated file, and outputs a map of whether each lineitem is present in the file
func ConvertToMap(f *os.File) (map[string]bool, error) {
	m := make(map[string]bool)
	cb := func(text string) {
		runeLength := len([]rune(text))

		if runeLength >= 8 && runeLength < 64 {
			m[text] = true
		}
	}

	err := iterateNewlineFile(f, cb)
	return m, err
}

//ConvertToSlice takes a newline delineated file, and outputs a slice of strings of the lines in that file
func ConvertToSlice(f *os.File) ([]string, error) {
	//this will be returned
	var sl []string
	cb := func(text string) {
		sl = append(sl, text)
	}

	err := iterateNewlineFile(f, cb)

	return sl, err
}
