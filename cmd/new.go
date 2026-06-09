package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type NewCmd struct {
	ProjectName string `arg:"" name:"project-name" help:"Project name to create."`
}

func (c *NewCmd) Run() error {
	if err := createNewProject(c.ProjectName); err != nil {
		return fmt.Errorf("creating project: %w", err)
	}
	fmt.Printf("✓ Successfully created project '%s'\n", c.ProjectName)
	fmt.Printf("\nNext steps:\n")
	fmt.Printf("  cd %s\n", c.ProjectName)
	fmt.Printf("  docker-compose -f deploy/docker-compose.yml up -d\n")
	fmt.Printf("  go run main.go\n")
	return nil
}

func createNewProject(projectName string) error {
	// Check if directory already exists
	if _, err := os.Stat(projectName); !os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' already exists", projectName)
	}

	// Clone ares-layout repository
	fmt.Printf("Cloning ares-layout template...\n")
	cmd := exec.Command("git", "clone", "https://github.com/xushuhui/ares-layout.git", projectName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone template: %w", err)
	}

	// Remove .git directory
	gitDir := filepath.Join(projectName, ".git")
	if err := os.RemoveAll(gitDir); err != nil {
		return fmt.Errorf("failed to remove .git directory: %w", err)
	}

	// Update go.mod with new module name
	goModPath := filepath.Join(projectName, "go.mod")
	_, err := os.ReadFile(goModPath)
	if err != nil {
		return fmt.Errorf("failed to read go.mod: %w", err)
	}

	// Replace module name
	newContent := []byte(fmt.Sprintf("module %s\n\ngo 1.23\n\nrequire (\n\tgithub.com/go-redis/redis/v8 v8.11.5\n\tgithub.com/lib/pq v1.10.9\n\tgithub.com/xushuhui/ares v0.1.0\n)\n", projectName))
	if err := os.WriteFile(goModPath, newContent, 0644); err != nil {
		return fmt.Errorf("failed to update go.mod: %w", err)
	}

	// Initialize new git repository
	cmd = exec.Command("git", "init")
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		// Not critical, just warn
		fmt.Printf("Warning: failed to initialize git repository\n")
	}

	return nil
}
