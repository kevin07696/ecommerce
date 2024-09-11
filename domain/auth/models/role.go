package models

import (
	"strings"
)

type Role string

func (m Models) NewRole(role string) (Role, error) {
	role = strings.TrimSpace(role)
	if role == "" {
		return "", ErrEmptyRole
	}

	validRoles := map[string]bool{
		"vendor":    true,
		"shopper":   true,
		"developer": true,
	}

	if !validRoles[role] {
		return "", ErrInvalidRole
	}

	return Role(role), nil
}
