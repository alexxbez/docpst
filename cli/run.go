package cli

import (
	"fmt"
	"os"
)

func Run(argv []string) {
	if len(argv) == 3 {
		if argv[1] == "new" {
			create(argv[2])
		} else {
			fmt.Println("Not a valid command")
			os.Exit(1)
		}
	} else if len(argv) == 2 {
		if argv[1] == "watch" {
			watch()
		} else {
			fmt.Println("Not a valid command")
			os.Exit(1)
		}
	} else {
		fmt.Println("Usage: docs [command] [arg]")
		os.Exit(1)
	}
}
