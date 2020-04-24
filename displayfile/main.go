package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	args := os.Args[1:]
	len := 0

	for i := range args {

		len = i + 1

	}
	if len > 1 {

		fmt.Println("Too many arguments")

	} else if len < 1 {

		fmt.Println("File name missing")

	} else {

		filename := args[0]
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("%v", err.Error())
		}

		fmt.Println(string(file))

	}
}
