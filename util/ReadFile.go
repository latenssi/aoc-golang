package util

import (
	"os"
	"regexp"
	"strings"
)

func ReadFile(path string, target *string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	s := string(f)

	if target != nil {
		*target = s
	}

	return s, nil
}

func RemoveEmptyLines(input string) (string, error) {
	regex, err := regexp.Compile("\n\n")
	if err != nil {
		return "", err
	}
	s := regex.ReplaceAllString(input, "\n")
	s = strings.TrimSpace(s)
	return s, nil
}

func Lines(input string) []string {
	return strings.Split(input, "\n")
}
