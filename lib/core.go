package reload

import (
	"regexp"
)

// comment
func Launch(in, out string) {
	// ! TMP, populate the vars, next use the bash file for executing the tests
	// VARS
	input, output := in, out
	var result []string
	regex_formula := `(\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\))|([.|,|?|!|:|;]{1,})|(['])|([\w\d\(\)']{1,}\b)`
	// regex formulas
	// regex_formula := `(?P<word>(?:[\w\d]+)|(?:[.|,|?|!|:|;|']+)|(?P<tag>\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\)))`
	// regex_formula := `(?P<word>[\w\d]+)|(?P<punc>[.|,|?|!|:|;|']+)|(?P<tag>\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\))`
	// var isTag, r = regexp.MustCompile(regex_tag)

	// * Read input file
	file_content := ReadFile(input)
	// fmt.Println("FILE: ", input, "CONTENT: ", file_content)

	// SPLIT the content
	r, _ := regexp.Compile(regex_formula)
	match := r.FindAllString(file_content, -1)
	// fmt.Println("RGX match: ", match)
	for i, item := range match {
		// fmt.Println(item[:1])
		// fmt.Println(item)
		// ? check if string is a itemalid tag (bin,hex,up,low..)
		if string(item[0]) == "(" && len(item) >= 4 {
			// fmt.Println("RGX found: ", item)
			switch item {
			case "(bin)":
				// fmt.Println(item, " > binary")
				result[len(result)-1] = BinaryToDec(result[len(result)-1])
				// fmt.Println(BinaryToDec(item))
			case "(hex)":
				// fmt.Println(item, " > hex")
				result[len(result)-1] = HexToDec(result[len(result)-1])
				// fmt.Println(HexToDec(item))
			// up, low, cap tag
			default:
				tag, num := IsItTag(item)
				// if numbered tag
				if tag != "" {
					if num > 1 {
						// fmt.Println("TAG: ", tag, "//NUMBER", num)
						for i := len(result) - num; i < len(result); i++ {
							// result[i] = ToUpper(result[i])
							switch tag {
							case "up":
								result[i] = ToUpper(result[i])
							case "low":
								result[i] = ToLower(result[i])
							case "cap":
								result[i] = Capitalize(result[i])
							}
						}
						// continue
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
					// not aa valid tag
				} else {
					// item is a word
					// fmt.Println("RES2: ", result, " ITEM:", item)
					result = append(result, item)
				}
			}
		} else {
			// ? item is a word
			// item = item
			// fmt.Println("RES_: ", result, " ITEM:", item)
			if item == "a" || item == "A" {
				// fmt.Println("letter:", match[i+1][:1])
				if IsVowel(match[i+1][:1]) {
					item = item + "n"
				} else {
					item = item
				}
				// fmt.Println("letter:", rune(match[i+1][:1]))
				// letter := runes(match[i+1][:1])
			}
			// fmt.Println("RES: ", result)
			result = append(result, item)
			// todo if punctuation
		}
	}

	// * Write file
	// fmt.Println("RES: ", WriteResult(result))
	WriteFile(WriteResult(result), output)
}
