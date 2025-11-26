package main

import (
	"fmt"

	"github.com/orchard9/tui-diagrams/pkg/diagrams"
)

func main() {
	// Create an authentication sequence diagram
	seq := diagrams.NewSequenceDiagram()

	seq.AddActor("user", "User").
		AddActor("client", "Client").
		AddActor("server", "Server").
		AddActor("db", "Database")

	seq.AddMessage("user", "client", "Login", diagrams.MessageSync).
		AddMessage("client", "server", "POST /auth", diagrams.MessageSync).
		AddMessage("server", "db", "Query User", diagrams.MessageSync).
		AddMessage("db", "server", "User Data", diagrams.MessageReturn).
		AddMessage("server", "server", "Validate", diagrams.MessageSync).
		AddMessage("server", "client", "JWT Token", diagrams.MessageReturn).
		AddMessage("client", "user", "Success", diagrams.MessageReturn)

	fmt.Println("Authentication Sequence Diagram")
	fmt.Println("================================")
	fmt.Println(seq.Render())
}
