package commandone

import (
	"fmt"

	"github.com/dihedron/go-archetype/command/base"
)

type CommandOne struct {
	base.Command

	// Peers              []cluster.Peer `short:"t" long:"target" description:"The address of the node in the cluster to send the request to." required:"yes"`
	// HealthCheckService string         `short:"h" long:"health-check" description:"Which gRPC service to health check when searching for the leader." optional:"yes" default:"quis.RaftLeader"`
}

func (cmd *CommandOne) Execute(args []string) error {
	fmt.Printf("Executing CommandOne command")
	return nil
}
