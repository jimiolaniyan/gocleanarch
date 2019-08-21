package gocleanarch

import "fmt"

// User is a simple data structure for a user. It belongs in the Entities layer.
type User struct {
	id string
	Username string
}

func (u *User) IsSame(user *User) bool {
	fmt.Println(u.id, user.id)
	return u.id == user.id
}

func (u *User) SetId(id string) {
	u.id = id
}

func (u *User) GetId() string {
	return u.id
}