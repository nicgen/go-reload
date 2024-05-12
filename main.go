package main

import (
	"fmt"
	reload "reload/lib"
	"strconv"
)

func main() {
	hex := reload.Hex2decimal("1A")
	fmt.Printf("type: %T\n", hex)
	fmt.Print(hex)
	bin := reload.Bin2decimal("101")
	fmt.Printf("type: %T\n", bin)
	fmt.Print(bin)
	fmt.Println("-------\n")
	fmt.Printf("type: %T\n", strconv.Itoa(bin))
	fmt.Println(strconv.Itoa(bin))
}
