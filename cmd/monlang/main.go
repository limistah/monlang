package main

import (
	"fmt"
	"github.com/limistah/monlang/packages/repl"
	"os"
	"os/user"
)

const MONKEY_FACE = `
 _      ____  _      _     ____  _      _____
/ \__/|/  _ \/ \  /|/ \   /  _ \/ \  /|/  __/
| |\/||| / \|| |\ ||| |   | / \|| |\ ||| |  _
| |  ||| \_/|| | \||| |_/\| |-||| | \||| |_//
\_/  \|\____/\_/  \|\____/\_/ \|\_/  \|\____\
`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the monlang programming language!\n", user.Username)
	fmt.Printf("%s\n", MONKEY_FACE)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
