package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s ! This is the monkey programming language!\n", user.Username)
	fmt.Printf("Type your commands\n ")
	repl.Start(os.Stdin, os.Stdout)
}
