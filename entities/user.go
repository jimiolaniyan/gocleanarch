package entities

// User is a simple data structure for a user. It belongs in the Entities layer.
type User struct {
	Entity
	username string
}

func NewUser(username string) *User {
	return &User{username: username}
}

func (u *User) Username() string {
	return u.username
}
