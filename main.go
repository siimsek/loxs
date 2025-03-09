package main

import (
	// ...existing imports...
	"embed"
	"os"
	"os/exec"
	"path/filepath"
)


//go:embed loxs.py
var pythonScript []byte

func main() {
	// Write the embedded Python script to a temporary file.
	tempFile := filepath.Join(os.TempDir(), "loxs.py")
	if err := os.WriteFile(tempFile, pythonScript, 0755); err != nil {
		// ...existing error handling...
		os.Exit(1)
	}

	// Execute the Python script from the temporary file.
	cmd := exec.Command("python3", tempFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		// ...existing error handling...
		os.Exit(1)
	}
}
