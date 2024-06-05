package reload

import (
	"fmt"
	"os"
)

// ReadFile reads a file from a string (name of the file)
func ReadFile(filename string) string {
	var content string = ""
	// OS version
	file_content, err := os.ReadFile(filename)
	content = string(file_content)
	if err != nil {
		fmt.Println("File writing error", err)
	}
	return string(content)
}

// WriteFile writes a content to a file (content, filename)
func WriteFile(content, output string) error {
	// todo, create a test for this function
	data := []byte(content)
	err := os.WriteFile(output, data, 0644)
	if err != nil {
		fmt.Println("File writing error", err)
		return err
	}
	return err
}

// WriteResult format a slice to a string (punctuation)
func WriteResult(str []string) string {
	// todo if punc else add space
	var toggle bool
	res := ""
	for _, word := range str {
		if IsItPunc(word) {
			if word == "'" {
				// todo check space before
				res = res[:len(res)-1]
				if toggle == false {
					toggle = true
					word = " '"
				} else {
					toggle = false
					word = "' "
				}
			} else {
				res = res[:len(res)-1]
				word = word + " "
			}
			res += word
		} else if word == ":'" {
			word = ": '"
			res += word
		} else {
			res += word + " "
		}
	}
	if res[len(res)-1:] == " " {
		res = res[:len(res)-1] + "\n"
		return res
	}
	res = res + "\n"
	return res
}
