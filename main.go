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
	// res    []string
	// file   string
)

func main() {
	// ! TMP, populate the vars, next use the bash file for executing the tests
	regex_formula := `(?P<tag>\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\))|(?P<punc>[.|,|?|!|:|;|']+)|(?P<word>\b[\w\d\(\)']+\b)`
	regex_formula_tag := `\(\b(?P<tag>cap|up|low)(?:,\s?(?P<num>\d+))?\)`
	// regex_formula := `(?P<word>(?:[\w\d]+)|(?:[.|,|?|!|:|;|']+)|(?P<tag>\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\)))`
	// regex_formula := `(?P<word>[\w\d]+)|(?P<punc>[.|,|?|!|:|;|']+)|(?P<tag>\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\))`
	// var isTag, r = regexp.MustCompile(regex_tag)

	// ARGUMENTS
	arguments := os.Args[1:]
	args_len := len(arguments)
	// fmt.Println("args:", arguments)
	if args_len != 2 {
		fmt.Println("Hey dumbass, not emough arguments!")
		os.Exit(1) // https://golang.org/pkg/os/#Exit
	} else {
		input = arguments[0]
		output = arguments[1]
	}

	fmt.Println("IN:", input, "OUT:", output)

	// READ FILE
	file_content := reload.ReadFile(input)
	fmt.Println("FILE: ", input, "CONTENT: ", file_content)

	// WRITE FILE
	reload.WriteFile(file_content, output)

	// SPLIT tags
	t, _ := regexp.Compile(regex_formula_tag)
	match_tag := t.FindAllString(file_content, -1)
	fmt.Println("RGX match: ", match_tag)
	// re := regexp.MustCompile(regex_formula)
	groupNames := t.SubexpNames()
	for matchNum, match := range t.FindAllStringSubmatch(file_content, -1) {
		for groupIdx, group := range match {
			name := groupNames[groupIdx]
			if name == "" {
				name = "*"
			}
			fmt.Printf("#%d text: '%s', group: '%s'\n", matchNum, group, name)
		}
	}
	// TODO regex for tag function

	// SPLIT the content
	r, _ := regexp.Compile(regex_formula)
	match := r.FindAllString(file_content, -1)
	fmt.Println("RGX match: ", match)
	for _, item := range match {
		fmt.Println(item)
		// check if string is a itemalid tag (bin,hex,up,low..)
		if string(item[0]) == "(" && len(item) >= 4 {
			// fmt.Println("RGX found: ", item)
			switch item {
			case "(bin)":
				fmt.Println(item, " > binary")
			case "(hex)":
				fmt.Println(item, " > hex")
			// todo create regex function to determine if tag is correct https://regex101.com/r/3MOcx8/3
			case "(cap*":
				fmt.Println(item, " > capitalize")
			case "(low*":
				fmt.Println(item, " > lowercase")
			case "(up":
				fmt.Println(item, " > uppercase")
			default:
				// fmt.Println("I've never tried", item, "before")
				// if strings.Contains(item, ",") {
				// 	fmt.Println("TAG found: ", item)
				// }
				fmt.Println(item, " > not a valid tag")
			}
		}
	}

	// ? ALT, with groupname (way overkill)
	// re := regexp.MustCompile(regex_formula)
	// groupNames := re.SubexpNames()
	// for matchNum, match := range re.FindAllStringSubmatch(file_content, -1) {
	// 	for groupIdx, group := range match {
	// 		name := groupNames[groupIdx]
	// 		if name == "" {
	// 			name = "*"
	// 		}
	// 		fmt.Printf("#%d text: '%s', group: '%s'\n", matchNum, group, name)
	// 	}
	// }

	fmt.Print("\n")
	// ? TEST lib/fucntions
	hex := reload.Hex2decimal("1A")
	fmt.Println("HEX: ", hex)
	fmt.Printf("type: %T\n", hex)
	bin := reload.Bin2decimal("101")
	fmt.Println("BIN: ", bin)
	fmt.Printf("type: %T\n", bin)
	fmt.Print("\n")
	// ? TEST: REMOVE A WORD in an array
	myWords := []string{"one", "two", "three", "four"}
	fmt.Println(myWords)
	myWords = reload.RemoveIndex(myWords, 2)
	fmt.Println(myWords) // [one two four]
}
