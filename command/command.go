package command

import (
	"github.com/dihedron/go-archetype/command/commandone"
	"github.com/dihedron/go-archetype/command/commandtwo"
	"github.com/dihedron/go-archetype/command/version"
)

// Commands is the set of root command groups.
type Commands struct {
	// CommandONe runs the CommandOne command
	//lint:ignore SA5008 go-flags uses multiple tags to define aliases and choices
	CommandOne commandone.CommandOne `command:"one" alias:"1" alias:"o" description:"Run some command."`
	// CommandTwo runs the CommandTwo command
	//lint:ignore SA5008 go-flags uses multiple tags to define aliases and choices
	CommandTwo commandtwo.CommandTwo `command:"two" alias:"2" alias:"t" description:"Run some command."`
	// Version prints the application version information and exits.
	//lint:ignore SA5008 go-flags uses multiple tags to define aliases and choices
	Version version.Version `command:"version" alias:"ver" alias:"v" description:"Show the command version and exit."`
}
