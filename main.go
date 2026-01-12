package main

import (
	"github.com/michcikk/ketchup/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}