package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/dihedron/go-archetype/metadata"
	"github.com/jessevdk/go-flags"
)

// If the application only has the root command, use this approach;
// if will provide a version option to print the application version
// and metadata and then pass control to a single command struct.
func main_without_subcommands() {

	defer cleanup()

	if len(os.Args) == 2 && (os.Args[1] == "version" || os.Args[1] == "--version") {
		metadata.Print(os.Stdout)
		os.Exit(0)
	} else if len(os.Args) == 3 && os.Args[1] == "version" && (os.Args[2] == "--verbose" || os.Args[2] == "-v") {
		metadata.PrintFull(os.Stdout)
		os.Exit(0)
	}

	// these are the single command options
	var options struct {
		SettingA string `short:"a" long:"setting-a" optional:"true" default:"whatever"`
		SettingB string `short:"b" long:"setting-b" optional:"true" default:"whatever"`
	}
	args, err := flags.Parse(&options)
	if err != nil {
		slog.Error("error parsing command line", "error", err)
		fmt.Fprintf(os.Stderr, "Invalid command line: %v\n", err)
		os.Exit(1)
	}

	for _, arg := range args {
		// do somting with the free args
		fmt.Printf("Argument: %s\n", arg)
	}
}
