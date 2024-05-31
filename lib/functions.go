package reload

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFile(filename string) string {
	var content string = ""
	file, err := os.Open(filename) // open the file for reading
	if err != nil {
		fmt.Println("File reading error", err)
	}
	defer file.Close() // closes the file after everything is done
	// fmt.Println(file.Stat())

	// OS version
	contentNEW, err := os.ReadFile(filename)
	fmt.Println(string(contentNEW))
	content = string(contentNEW)
	if err != nil {
		fmt.Println("File writing error", err)
	}

	// BUFIO scanner text (no split)
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println("BUFIO_SCAN:", scanner.Text())
	// 	content = scanner.Text()
	// }

	// BUFIO scanner words (split by words)
	// scanner := bufio.NewScanner(file)
	// // Set the split function for the scanning operation.
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	fmt.Println("BUFIO_SCAN:", scanner.Text())
	// 	content = scanner.Text()
	// }

	// BUFIO ERROR
	// if err := scanner.Err(); err != nil {
	// 	fmt.Println(err)
	// }

	return string(content)
}

func WriteFile(content, output string) error {
	data := []byte(content)
	err := os.WriteFile(output, data, 0644)
	if err != nil {
		fmt.Println("File writing error", err)
		return err
	}
	fmt.Println("Data successfully written to file")
	return err
}

func Hex2decimal(hexa string) string {
	output, err := strconv.ParseInt(hexa, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.Itoa(int(output))
}

func Bin2decimal(bin string) string {
	output, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.Itoa(int(output))
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// func before ()
