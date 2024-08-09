package main

import (
	"fmt"
	"github.com/limistah/monlang/packages/repl"
	"os"
	"os/user"
)

const MONKEY_FACE = `
  __  __   ____   _   _  _                 _   _   _____ 
 |  \/  | / __ \ | \ | || |         /\    | \ | | / ____|
 | \  / || |  | ||  \| || |        /  \   |  \| || |  __ 
 | |\/| || |  | || . ` || |       / /\ \  | . ` || | |_ |
 | |  | || |__| || |\  || |____  / ____ \ | |\  || |__| |
 |_|  |_| \____/ |_| \_||______|/_/    \_\|_| \_| \_____|
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
