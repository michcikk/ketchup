package executor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/michcikk/ketchup/internal/venv"
)

type Executor struct {
	ProjectPath string
	VenvManager *venv.Manager
}

func NewExecutor(projectPath string) *Executor {
	return &Executor{
		ProjectPath: projectPath,
		VenvManager: venv.NewManager(projectPath),
	}
}

// Setup prepares the environment for execution
func (e *Executor) Setup() error {
	// Create venv if it doesn't exist
	if !e.VenvManager.Exists() {
		if err := e.VenvManager.Create(); err != nil {
			return err
		}
	}
	
	// Install dependencies
	requirementsPath := filepath.Join(e.ProjectPath, "requirements.txt")
	if err := e.VenvManager.InstallDependencies(requirementsPath); err != nil {
		return err
	}
	
	return nil
}

// Run executes a Python script
func (e *Executor) Run(scriptPath string) error {
	// Ensure setup is complete
	if err := e.Setup(); err != nil {
		return err
	}
	
	// Get absolute script path
	absScriptPath, err := filepath.Abs(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	
	fmt.Printf("[ketchup] Running %s...\n", filepath.Base(scriptPath))
	
	// Execute the script
	cmd := exec.Command(e.VenvManager.GetPythonPath(), absScriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Dir = e.ProjectPath
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("script execution failed: %w", err)
	}

	fmt.Println("[ketchup] ✓ Completed successfully")
	return nil
}