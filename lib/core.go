package reload

import (
	"regexp"
)

// main function that launches all corrections
func Launch(in, out string) {
	// VARS
	input, output := in, out
	var result []string
	regex_formula := `(\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\))|([.|,|?|!|:|;]{1,})|(['])|([\w\d\(\)']{1,}\b)`
	file_content := ReadFile(input)

	// SPLIT the content
	r, _ := regexp.Compile(regex_formula)
	match := r.FindAllString(file_content, -1)
	for i, item := range match {
		// ? check if string is a itemalid tag (bin,hex,up,low..)
		if string(item[0]) == "(" && len(item) >= 4 {
			switch item {
			case "(bin)":
				result[len(result)-1] = BinaryToDec(result[len(result)-1])
			case "(hex)":
				result[len(result)-1] = HexToDec(result[len(result)-1])
			default:
				tag, num := IsItTag(item)
				if tag != "" {
					if num > 1 { // fmt.Println("TAG: ", tag, "//NUMBER", num)
						for i := len(result) - num; i < len(result); i++ {
							switch tag {
							case "up":
								result[i] = ToUpper(result[i])
							case "low":
								result[i] = ToLower(result[i])
							case "cap":
								result[i] = Capitalize(result[i])
							}
						}
					} else {
						switch tag {
						case "up":
							result[len(result)-1] = ToUpper(result[len(result)-1])
						case "low":
							result[len(result)-1] = ToLower(result[len(result)-1])
						case "cap":
							result[len(result)-1] = Capitalize(result[len(result)-1])
						}
					}
				} else {
					result = append(result, item)
				}
			}
		} else {
			if item == "a" || item == "A" {
				if IsVowel(match[i+1][:1]) {
					item = item + "n"
				} else {
					item = item
				}
			}
			result = append(result, item)
			// todo if punctuation
		}
	}
	WriteFile(WriteResult(result), output)
}
