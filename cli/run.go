package cli

import (
// "fmt"
)

func Run(argv []string) {
	if len(argv) == 3 {
		if argv[1] == "new" {
			create(argv[2])
		}
	}
}
