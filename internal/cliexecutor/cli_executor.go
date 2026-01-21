package cliexecutor

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/SaketV8/test-generator-ai/internal/utils"
)

func RunGoTest(buildDir string, logFilePath string) error {
	// Create (or truncate) the log file
	utils.PrintBoxedText("Running Go Test")
	logFile, err := utils.CreateEmptyFile(logFilePath)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer logFile.Close()

	// Prepare make command
	// cmd := exec.Command("go test -v ./pkg/test/...")
	cmd := exec.Command("go", "test", "-v", "./pkg/test/...")
	cmd.Dir = buildDir

	// Send output to both stdout and the log file
	multiOut := io.MultiWriter(os.Stdout, logFile)
	cmd.Stdout = multiOut
	cmd.Stderr = multiOut

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(multiOut, "test run failed 🐳: %v\n", err)
		return fmt.Errorf("test run failed: %w", err)
	}

	fmt.Fprintln(multiOut, "\nTest run succeeded")
	utils.PrintBoxedText("Running Go Test completed")
	return nil
}
