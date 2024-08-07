package faker

import (
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/brianvoe/gofakeit/v7"
	"github.com/nochzato/ticketopia-user-service/pkg/hashpass"
)

type User struct {
	FullName string `fake:"{name}"`
	Username string `fake:"{username}"`
	Password string `fake:"{password:true,false,true,true,false,6}"`
	Email    string `fake:"{email}"`
}

func RandomUser() (user *User, err error) {
	err = gofakeit.Struct(&user)
	return
}

func RandomHashedPassword() (string, error) {
	password := gofakeit.Password(true, false, true, true, false, 6)
	return hashpass.Hash(password)
}

func RandomFullName() string {
	return gofakeit.Name()
}

func RandomEmail() string {
	return gofakeit.Email()
}
