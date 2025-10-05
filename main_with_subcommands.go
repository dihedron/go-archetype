package main

import (
	"os"

	"github.com/dihedron/go-archetype/command"
	"github.com/jessevdk/go-flags"
)

// If the application has multiple subcommands, use this version of main()
// which will initialize the subcommands and parse the command line arguments.
// The version subcommand is initialised and parsed here. For the commands
// implementation see the command package.
func main_with_subcommands() {

	defer cleanup()

	options := command.Commands{}
	if _, err := flags.NewParser(&options, flags.Default).Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			//fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		default:
			//fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
}
