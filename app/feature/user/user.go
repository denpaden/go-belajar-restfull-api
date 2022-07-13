package user

type User struct {
	Username string
	Name     string
}

func NewUser(username string, name string) *User {
	return &User{Username: username, Name: name}
}
