package reload

import (
	"fmt"
	"strconv"
	"strings"
)

// HexToDec converts a Hexadecimal to a decimal
func HexToDec(hexa string) string {
	output, err := strconv.ParseInt(hexa, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.Itoa(int(output))
}

// BinaryToDec converts a binary to a decimal
func BinaryToDec(bin string) string {
	output, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(bin, output, strconv.Itoa(int(output)))
	return strconv.Itoa(int(output))
}

// ToLower converts a string to lowercase
func ToLower(str string) string {
	return strings.ToLower(str)
}

// ToUpper converts a string to uppercase
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

// Capitalize capitalizes the first letter of a string
func Capitalize(str string) string {
	// return strings.Title(str)
	// return str
	if str == "" {
		return str
	}
	// Split the string into words
	letters := strings.Split(str, "")
	// fmt.Println(letters)
	// Capitalize the first letter of each word
	for i, letter := range letters {
		// words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		// i == 0 > words[i] = strings.ToUpper(string(word[i]))
		if i == 0 {
			letters[i] = strings.ToUpper(string(letter))
		} else {
			letters[i] = strings.ToLower(string(letter))
		}
	}
	// Join the capitalized words back into a string
	// return strings.Join(words, " ")
	// fmt.Println(letters)
	return strings.Join(letters, "")
}
