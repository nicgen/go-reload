package main

import (
	"fmt"
	"os"
	reload "reload/lib" // my external functions
)

func main() {
	var input, output string
	arguments := os.Args[1:]
	args_len := len(arguments)
	if args_len == 2 {
		input = arguments[0]
		output = arguments[1]
	} else if args_len > 2 {
		fmt.Println("Executable arguments error: too many arguments (input output)")
		os.Exit(1)
	} else {
		fmt.Println("Executable arguments error: not emough arguments (input output)")
		os.Exit(1) // https://golang.org/pkg/os/#Exit
	}
	reload.Launch(input, output)
}
