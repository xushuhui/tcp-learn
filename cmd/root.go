package cmd

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

type CLI struct {
	New      NewCmd      `cmd:"" help:"Create a new Ares project"`
	Generate GenerateCmd `cmd:"" help:"Generate code from templates"`
	Gen      GenCmd      `cmd:"" help:"Generate GORM models and query code"`
	OpenAPI  OpenAPICmd  `cmd:"" name:"openapi" help:"Generate OpenAPI specification"`
}

func Execute() error {
	var cli CLI
	ctx := kong.Parse(
		&cli,
		kong.Name("aresctl"),
		kong.Description("Aresctl is a powerful CLI tool that helps you develop applications with the Ares framework. It provides code generation, OpenAPI documentation, and more."),
	)
	return ctx.Run()
}

type OpenAPICmd struct {
}

func (c *OpenAPICmd) Run() error {
	if err := generateOpenAPI(); err != nil {
		return fmt.Errorf("generating OpenAPI: %w", err)
	}
	fmt.Println("✓ Generated openapi.yaml successfully")
	return nil
}

func generateOpenAPI() error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Default paths - can be made configurable via flags
	routeDir := cwd + "/internal/server"
	apiDir := cwd + "/api"
	outputFile := cwd + "/openapi.yaml"

	// Check if directories exist
	if _, err := os.Stat(routeDir); os.IsNotExist(err) {
		return fmt.Errorf("route directory not found: %s", routeDir)
	}
	if _, err := os.Stat(apiDir); os.IsNotExist(err) {
		return fmt.Errorf("api directory not found: %s", apiDir)
	}

	// Generate OpenAPI spec
	GenerateOpenAPI(routeDir, apiDir, outputFile)

	return nil
}
