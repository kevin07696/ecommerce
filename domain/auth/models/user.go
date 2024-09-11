package models

type User struct {
	Username Username `bson:"username,omitempty"`
	Role     Role     `bson:"role,omitempty"`
	Email    Email    `bson:"inline"`
}

func (m Models) NewUser(username Username, email Email, role Role) User {
	return User{
		Username: username,
		Email:    email,
		Role:     role,
	}
}
