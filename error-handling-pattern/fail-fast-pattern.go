package main

import (
	"fmt"
	"log"
)

func main() {
	var userDatabase = map[string]string{
		"john": "secret",
	}

	username := "john"
	password := "secr3t"

	// get data user and validate
	if ok := userDatabase[username]; len(ok) == 0 {
		log.Fatal(`User or password is wrong`)
	}

	// validate password
	pass := userDatabase[username]
	if password != pass {
		log.Fatal(`User or password is wrong`)
	}

	// login is success
	fmt.Println(`Successfully login to system`)
}
