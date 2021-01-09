package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/aewens/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, this is the REPL for Monkey\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
