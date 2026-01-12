package venv

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Manager struct {
	VenvPath string
}

func NewManager(projectPath string) *Manager {
	return &Manager{
		VenvPath: filepath.Join(projectPath, ".venv"),
	}
}

// Exists checks if virtual environment exists
func (m *Manager) Exists() bool {
	_, err := os.Stat(m.VenvPath)
	return !os.IsNotExist(err)
}

// Create creates a new virtual environment
func (m *Manager) Create() error {
	fmt.Println("[ketchup] Creating virtual environment...")
	
	cmd := exec.Command("python3", "-m", "venv", m.VenvPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create venv: %w", err)
	}

	fmt.Println("[ketchup] ✓ Virtual environment created")
	return nil
}

// GetPythonPath returns the path to the Python executable in the venv
func (m *Manager) GetPythonPath() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(m.VenvPath, "Scripts", "python.exe")
	}
	return filepath.Join(m.VenvPath, "bin", "python")
}

// GetPipPath returns the path to pip in the venv
func (m *Manager) GetPipPath() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(m.VenvPath, "Scripts", "pip.exe")
	}
	return filepath.Join(m.VenvPath, "bin", "pip")
}

// InstallDependencies installs packages from requirements.txt
func (m *Manager) InstallDependencies(requirementsPath string) error {
	if _, err := os.Stat(requirementsPath); os.IsNotExist(err) {
		fmt.Println("[ketchup] No requirements.txt found, skipping dependencies")
		return nil
	}

	fmt.Println("[ketchup] Installing dependencies...")
	cmd := exec.Command(m.GetPipPath(), "install", "-r", requirementsPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install dependencies: %w", err)
	}

	fmt.Println("[ketchup] ✓ Dependencies installed")
	return nil
}