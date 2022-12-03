package util

import "os"

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
