package main

import "github.com/asaskevich/govalidator"

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type User struct {
	Username string `valid:"alpha,required"`
	Password string `valid:"alpha,required"`
}

func main() {
	user := User{Username: "something", Password: ""}
	valid, err := govalidator.ValidateStruct(user)

	if !valid {
		usernameError := govalidator.ErrorByField(err, "Username")
		passwordError := govalidator.ErrorByField(err, "Password")
		if usernameError != "" {
			println("Username Error :: " + usernameError)
		}
		if passwordError != "" {
			println("Password Error :: " + passwordError)
		}
	}
}
