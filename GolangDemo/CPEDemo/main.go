package main

import (
	"log"
	"regexp"
	"strings"
)

var (
	nonEscapedChar = regexp.MustCompile(`[a-zA-Z0-9\*\.\_\-]`)
)

func main() {
	cpe := `cpe:2.3:a:cgiirc:cgi\:irc:0.3.6:*:*:*:*:*:*:*`

	cpe2 := strings.ReplaceAll(cpe, "\\:", "\x00")

	comp := strings.Split(cpe2, ":")
	for i, part := range comp {
		part = strings.ReplaceAll(part, "\x00", ":")
		part = strings.ReplaceAll(part, "\\\\", "\x00")
		part = strings.ReplaceAll(part, "\\", "")
		part = strings.ReplaceAll(part, "\x00", `\`)
		comp[i] = part
	}

	log.Print("end")
}

func RemoveEscapeChar(s string) string {
	var result []rune

	for i, char := range s {
		log.Print(string(char))
		// Check if the character is not a letter or digit

		if i != len(s)-1 && char == '\\' && !nonEscapedChar.MatchString(string(s[i+1])) {
			log.Print("bindo")
			continue
		}
		result = append(result, char)
	}

	return string(result)
}

func GetFromName(CPEName string) {
	//strings.ReplaceAll()
	cpe := strings.ReplaceAll(CPEName, "\\:", "\x00")
	components := strings.Split(cpe, ":")
	for i, part := range components {
		components[i] = strings.ReplaceAll(part, "\x00", ":")
	}

	//Check if component valid
	if len(components) != 13 {
		log.Printf("found %v in %s", len(components), CPEName)
	}

}
