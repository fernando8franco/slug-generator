package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	godiacritics "gopkg.in/Regis24GmbH/go-diacritics.v2"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	slugs := []string{}

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		lowerCase := strings.ToLower(text)
		noDiacrits := godiacritics.Normalize(lowerCase)
		noSimbols := clearString(noDiacrits)
		slug := strings.ReplaceAll(noSimbols, " ", "-")

		slugs = append(slugs, slug)
	}

	fmt.Println(slugs)
}
