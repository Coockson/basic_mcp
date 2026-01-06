package main

import (
	"basic_mcp/calculator"
	"basic_mcp/greeter"
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "basic-mcp", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: greeter.Description}, greeter.SayHi)
	mcp.AddTool(server, &mcp.Tool{Name: "calculate", Description: calculator.Description}, calculator.Calculate)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
