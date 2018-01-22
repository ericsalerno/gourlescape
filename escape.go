package gourlescape

import "strings"

// Escape structure for escaping URLs
type Escape struct {
	Replacer   string
	RemoveTail bool
}

// URL escape a string to a url style
func (e Escape) URL(input string) string {

	if e.Replacer == "" {
		e.Replacer = "-"
	}

	output := ""
	input = strings.TrimSpace(input)
	runes := []rune(input)

	for _, character := range runes {
		if (character >= 97 && character <= 122) || (character >= 48 && character <= 57) {
			//Lowercase leets and numbers
			output += string(character)
		} else if character >= 65 && character <= 90 {
			//Capital letters, lowercase them
			output += string(character + 32)
		} else {
			output += e.Replacer
		}
	}

	if e.RemoveTail {
		output = strings.Trim(output, e.Replacer)
	}

	return output
}
