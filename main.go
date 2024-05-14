package main

import (
	"fmt"
	"os"
	"regexp"
	reload "reload/lib"
)

// ? cmd go run . input.txt output.txt
// and once compiled:
// ? reload input.txt output.txt

var (
	input  string
	output string
	file   string
)

func main() {
	arguments := os.Args[1:]
	fmt.Println("args:", arguments)
	fmt.Println("READZEFILE------------------------------------------")
	input = "sample.txt"
	output = "result.txt"

	mytext := reload.ReadFile(input)
	fmt.Println("FILE: ", input, "CONTENT: ", mytext)

	fmt.Println("REGEX------------------------------------------")

	myregex := `(\b(?:[\w\d]+)|(?:[.|,|?|!|:|;|']+)|\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\))`
	r, _ := regexp.Compile(myregex)

	match := r.FindAllString(mytext, -1)

	for _, v := range match {
		fmt.Println(v)
	}

	fmt.Println(match)

	fmt.Println("FUNCS------------------------------------------")

	hex := reload.Hex2decimal("1A")
	fmt.Println("HEX: ", hex)
	fmt.Printf("type: %T\n", hex)
	bin := reload.Bin2decimal("101")
	fmt.Println("BIN: ", bin)
	fmt.Printf("type: %T\n", bin)

	// fmt.Println("TYPE------------------------------------------")

	// fmt.Println("BIN/ITOA: ", strconv.Itoa(bin))
	// fmt.Printf("type: %T\n", strconv.Itoa(bin))
}
