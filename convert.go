package execute_library

import (
	"strings"
	"unicode"
)

func Convert(text string) string{
	if IsUpper(text)==true{
		return strings.ToLower(text)
	}else if IsLower(text)==true {
		return strings.ToUpper(text)
	}else{
		return "The string does not match the input type"
	}
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}