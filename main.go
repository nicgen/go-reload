package main

import (
	"fmt"
	"os"
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

func main() {
	var input, output string
	// * Arguments
	// fmt.Println(os.Args)
	arguments := os.Args[1:]
	args_len := len(arguments)
	if args_len != 2 {
		fmt.Println("Executable arguments error: not emough arguments (input output)")
		os.Exit(1) // https://golang.org/pkg/os/#Exit
	} else {
		input = arguments[0]
		output = arguments[1]
	}
	reload.Launch(input, output)
}

// ! MAIN END

// func test() {
// 	fmt.Print("\n")
// 	// ? TEST lib/functions
// 	hex := reload.HexToDec("1A")
// 	fmt.Println("HEX: ", hex)
// 	fmt.Printf("type: %T\n", hex)
// 	bin := reload.BinaryToDec("101")
// 	fmt.Println("BIN: ", bin)
// 	fmt.Printf("type: %T\n", bin)
// 	fmt.Print("\n")
// 	// ? TEST: REMOVE A WORD in an array
// 	myWords := []string{"one", "two", "three", "four"}
// 	fmt.Println(myWords)
// 	myWords = reload.RemoveIndex(myWords, 2)
// 	fmt.Println(myWords) // [one two four]
// 	fmt.Println(reload.Capitalize("ShinDeiru"))
// 	// Capitalize("ShinDeiru")
// }
