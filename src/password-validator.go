package main

import (
	"fmt"
	"os"

	"github.com/malls/password-validator/src/ascii"
	"github.com/malls/password-validator/src/list-converter"
)

func main() {
	pipedFile := os.Stdin
	fi, err := pipedFile.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Size() > 0 {

		//the first argument will be the file with the bad passwords
		badPasswordList := os.Args[1]

		//read the bad password file, then convert to a map we can reference
		f, err := os.Open(badPasswordList)

		if err != nil {
			panic("Please provide a valid invalid password list file.")
		}

		badPasswords, err := listconverter.ConvertToMap(f)

		if err != nil {
			panic(err)
		}

		inputPasswords, err := listconverter.ConvertToSlice(pipedFile)

		if err != nil {
			panic(err)
		}

		//go through the range of passwords and see if it's a good one
		for _, pass := range inputPasswords {
			passLength := len([]rune(pass))

			if passLength < 8 {
				fmt.Println(pass, "-> Error: Too Short")
			} else if passLength > 64 {
				fmt.Println(pass, "-> Error: Too Long")
			} else if ascii.Invalid(pass) {
				fmt.Println(pass, "-> Error: Invalid Characters")
			} else if _, ok := badPasswords[pass]; ok {
				fmt.Println(pass, "-> Error: Too Common")
			}
		}

	} else {
		fmt.Println("Please provide a newline separated file of passwords.")
		return
	}

}
