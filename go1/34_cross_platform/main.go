package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Current User
	fmt.Printf("Hi %s (id: %s)\n", user.Name, user.Uid)
	fmt.Println("Username:", user.Username)
	fmt.Println("Home Directory:", user.HomeDir)

	// 'real' user under sudo: https://stackoverflow.com/q/29733575/402585
	fmt.Println("Real User:", os.Getenv("SUDO_USER"))
	fmt.Println("ID:", os.Geteuid())
}
