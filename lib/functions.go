package reload

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadTxt() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func Hex2decimal(hexa string) int {
	output, err := strconv.ParseInt(hexa, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	// x := string(output)
	return int(output)
}

func Bin2decimal(bin string) int {
	output, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	// x := string(output)
	return int(output)
}

// func before ()
