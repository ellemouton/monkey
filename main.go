package main

import (
	"fmt"
	"log"
	"os"
	user "os/user"

	"github.com/ellemouton/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming "+
		"language!\n Feel free to type in commands\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
