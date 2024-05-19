package reload

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(filename string) string {
	var content string = ""
	file, err := os.Open(filename) // open the file for reading
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()       // closes the file after everything is done
	fmt.Println(file.Stat()) // return

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		content = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return string(content)
}

// func Hex2decimal(hexa string) int {
// 	output, err := strconv.ParseInt(hexa, 16, 32)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// x := string(output)
// 	return int(output)
// }

func Hex2decimal(hexa string) (string, error) {
	output, err := strconv.ParseInt(hexa, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	// x := string(output)
	// return int(output)
	return strconv.Itoa(int(output)), nil
}

func Bin2decimal(bin string) string {
	output, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	// x := string(output)
	// return int(output)
	return strconv.Itoa(int(output))
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// func before ()
