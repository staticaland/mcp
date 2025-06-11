package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Hello World MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)

	// Add a simple hello tool
	helloTool := mcp.NewTool("hello",
		mcp.WithDescription("Greet the user with a hello message"),
		mcp.WithString("name",
			mcp.Description("Name to greet (optional)"),
		),
	)

	s.AddTool(helloTool, handleHello)

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func handleHello(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := req.GetString("name")
	if err != nil || name == "" {
		name = "World"
	}

	message := fmt.Sprintf("Hello, %s!", name)
	return mcp.NewToolResultText(message), nil
}