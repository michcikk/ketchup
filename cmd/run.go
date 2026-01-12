package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/michcikk/ketchup/internal/executor"
)

var runCmd = &cobra.Command{
	Use:   "run [script.py]",
	Short: "Run a Python automation script",
	Long:  `Execute a Python script in an isolated virtual environment`,
	Args:  cobra.ExactArgs(1),
	RunE:  runScript,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runScript(cmd *cobra.Command, args []string) error {
	scriptPath := args[0]
	
	// Check if script exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return fmt.Errorf("script not found: %s", scriptPath)
	}
	
	// Get project directory (where the script is located)
	projectPath := filepath.Dir(scriptPath)
	if projectPath == "." {
		projectPath, _ = os.Getwd()
	}
	
	// Create and run executor
	exec := executor.NewExecutor(projectPath)
	return exec.Run(scriptPath)
}