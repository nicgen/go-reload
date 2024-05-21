package main

import (
	"fmt"
	"os"
	"regexp"
	reload "reload/lib" // my external functions
)

/*
Usage:
$ go run . input.txt output.txt
or once compiled
$ reload input.txt output.txt
todo
- verifications
	- arguments, type, number
	- does input exists
	- is it a valid text file
- add Usage to README
*/

var (
	input  string
	output string
	// file   string
)

func main() {
	// ! TMP, populate the vars, next use the bash file for executing the tests
	input = "sample.txt"
	output = "result.txt"
	regex_formula := `(\b(?:[\w\d]+)|(?:[.|,|?|!|:|;|']+)|\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\))`

	// ARGUMENTS
	arguments := os.Args[1:]
	args_len := len(arguments)
	// fmt.Println("args:", arguments)
	if args_len != 2 {
		fmt.Println("Hey dumbass, not emough arguments!")
		os.Exit(1) // https://golang.org/pkg/os/#Exit
	}

	// READ FILE
	file_content := reload.ReadFile(input)
	fmt.Println("FILE: ", input, "CONTENT: ", file_content)

	// SPLIT the content
	r, _ := regexp.Compile(regex_formula)
	match := r.FindAllString(file_content, -1)
	for _, v := range match {
		// fmt.Println(string(v[0]))
		if string(v[0]) == "(" && len(v) >= 4 {
			fmt.Println(v)

		}
		// if regexp.FindString("an", "highlands is a part of Scotland") {

		// }
	}

	// TEST lib/fucntions
	hex := reload.Hex2decimal("1A")
	fmt.Println("HEX: ", hex)
	fmt.Printf("type: %T\n", hex)
	bin := reload.Bin2decimal("101")
	fmt.Println("BIN: ", bin)
	fmt.Printf("type: %T\n", bin)
	// RENMOVE A WORD
	myWords := []string{"one", "two", "three", "four"}
	fmt.Println(myWords)
	myWords = reload.RemoveIndex(myWords, 2)
	fmt.Println(myWords) // [one two four]
}
