package models

import (
	"regexp"
	"strings"
)

type Username string

func (m Models) NewUsername(un string) (Username, error) {
	un = strings.TrimSpace(un)
	if un == "" {
		return "", ErrEmptyUsername
	}

	pattern := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{3,19}$`)
	isMatch := pattern.MatchString(un)
	if !isMatch {
		return "", ErrInvalidUsername
	}

	return Username(un), nil
}
