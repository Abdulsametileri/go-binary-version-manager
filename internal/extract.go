package internal

import (
	"fmt"
	"regexp"
)

func ExtractLibName(packageStr string) (string, error) {
	// This pattern looks for the last part of the path before any subdirectories and
	// the "@" symbol or the end of the string
	pattern := `github\.com/[\w-]+/([\w-]+)(?:/[\w-]+)*@[\w.-]+$`
	re := regexp.MustCompile(pattern)

	// Find the matches
	matches := re.FindStringSubmatch(packageStr)
	if len(matches) < 2 {
		return "", fmt.Errorf("could not parse library name from %s", packageStr)
	}
	return matches[1], nil
}
