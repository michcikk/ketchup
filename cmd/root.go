package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ketchup",
	Short: "python RPA CLI tool",
	Long: `ketchup is a command-line tool for creating, running, and managing 
RPA automation scripts written in Python`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add global flags here if needed
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}
