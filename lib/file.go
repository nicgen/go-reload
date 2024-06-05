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
	// fmt.Println("File content:", string(file_content))
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
	// fmt.Println("Data successfully written to file")
	return err
}

// WriteResult format a slice to a string (punctuation)
func WriteResult(str []string) string {
	// todo if punc else add space
	var toggle bool
	// fmt.Println(toggle)
	res := ""
	for _, word := range str {
		// fmt.Println("word: ", word, "len: ", len(word))
		// fmt.Println(word)
		if IsItPunc(word) {
			if word == "'" {
				// fmt.Println("ACCENT")
				res = res[:len(res)-1]
				if toggle == false {
					toggle = true
					// fmt.Println("WORD: ", word)
					// if res[len(res)-1:] == " " {
					// 	// fmt.Println(res)
					// 	// fmt.Println(res[len(res)-3:])
					// 	word = "'__"
					// }
					word = " '"
					// res += word
				} else {
					toggle = false
					// if res[len(res)-1:] == " " {
					// 	// fmt.Println("LOL")
					// }
					word = "' "
					// word = "'"
				}
				// res += word
				// fmt.Println(res)
				// continue
			} else {
				res = res[:len(res)-1]
				// res += word + " "
				word = word + " "
				// continue
			}
			res += word
		} else if word == ":'" {
			// fmt.Println("YOP")
			word = ": '"
			res += word
			// } else if word == "'," {
			// 	// fmt.Println("YOP")
			// 	word = "'>>>>,"
			// 	res += word
		} else {
			// fmt.Println("WRITE_RES:", res)
			res += word + " "
		}
	}
	if res[len(res)-1:] == " " {
		// fmt.Println("YOP", res[len(res)-1:])
		res = res[:len(res)-1] + "\n"
		// fmt.Println("ONE", res)
		return res
	}
	res = res + "\n"
	// fmt.Println("TWO", res)
	return res
}
