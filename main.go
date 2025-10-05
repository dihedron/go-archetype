package main

import (
	"os"
)

func main() {

	defer cleanup()

	if len(os.Args) > 1 && os.Args[1] == "with" {
		main_with_subcommands()
	} else if len(os.Args) > 1 && os.Args[1] == "without" {
		main_without_subcommands()
	}
}
