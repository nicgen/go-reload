package reload

import (
	"regexp"
	"strconv"
)

// IsItPunc checks if the word is a punctuation
func IsItPunc(str string) bool {
	if str[:1] == "." || str[:1] == "," || str[:1] == "?" || str[:1] == "!" || str[:1] == ":" || str[:1] == ";" || str[:1] == "'" {
		return true
	}
	return false
}

// function that returns tag + n
func IsItTag(str string) (tag string, num int) {
	// SPLIT tags
	regex_formula_tag := `(?P<name>cap|up|low)(?:,\s(?P<number>\d+))?`
	tag = ""
	num = 1
	t := regexp.MustCompile(regex_formula_tag)
	matches := t.FindStringSubmatch(str)
	if len(matches) != 0 {
		tag = matches[1]
		num, _ = strconv.Atoi(matches[2])
		// return
	}
	return
}

func IsVowel(l string) bool {
	switch l {
	case "a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H":
		// fmt.Println("letter: ", l)
		return true
	}
	return false
}
