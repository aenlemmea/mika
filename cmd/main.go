package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/aenlemmea/mika/console"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Hello %s.\n", user.Username)
	fmt.Printf("Type exit to exit. Use clear to clear the screen.\n")
	fmt.Printf("The keywords exit and clear are reserved for console use.\n")
	console.Console(os.Stdin, os.Stdout)
}
