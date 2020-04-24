package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]
	len := 0

	for i := range args {

		len = i + 1 // counting the length of a slice

	}

	if len >= 1 { // if length is bigger than 1

		for i := 0; i <= len-1; i++ {

			filename := args[i]

			file, err := ioutil.ReadFile(string(filename))

			if err != nil {

				for _, char := range err.Error() {

					z01.PrintRune(char)
				}
				z01.PrintRune(10)

			}
			for _, char := range string(file) {

				z01.PrintRune(char)

			}
		}
	} else if len < 1 {

		result := io.TeeReader(os.Stdin, os.Stdout)
		ioutil.ReadAll(result)

	}
}
