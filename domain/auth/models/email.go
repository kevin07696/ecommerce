package models

import (
	"regexp"
	"strings"
)

type Email struct {
	Local      string
	SubAddress string
	Domain     string
}

func (m Models) NewEmail(email string) (Email, error) {
	em := strings.TrimSpace(email)
	if em == "" {
		return Email{}, ErrEmptyEmail
	}

	pattern := regexp.MustCompile(`^([\w.]+)([+\-]?[\w.]*)@([\w]+[.][\w]+[.]?[\w]*)$`)
	isMatch := pattern.MatchString(em)
	if !isMatch {
		return Email{}, ErrInvalidEmail
	}
	submatches := pattern.FindAllStringSubmatch(em, 1)
	return Email{
		Local:      submatches[0][1],
		SubAddress: submatches[0][2],
		Domain:     submatches[0][3],
	}, nil
}

func (em Email) ToString() string {
	return em.Local + em.SubAddress + "@" + em.Domain
}
