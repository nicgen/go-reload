package main

import (
	"fmt"
	"os"
	reload "reload/lib"
	"strconv"
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
	fmt.Println("args:", os.Args)

	input = "sample.txt"
	output = "result.txt"

	fmt.Println("READ ZE FILE")
	reload.ReadTxt()

	fmt.Println("------------------------------------------")

	hex := reload.Hex2decimal("1A")
	fmt.Printf("type: %T\n", hex)
	fmt.Print(hex)
	bin := reload.Bin2decimal("101")
	fmt.Printf("type: %T\n", bin)
	fmt.Print(bin)

	fmt.Println("------------------------------------------")

	fmt.Printf("type: %T\n", strconv.Itoa(bin))
	fmt.Println(strconv.Itoa(bin))
}
