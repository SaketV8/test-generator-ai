package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// ========================================================
func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("failed to read file: %v", err)
		return "", err
	}
	return string(data), nil
}

// ========================================================

// ========================================================
func WriteFile(path string, content string) error {
	// Ensure the parent directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write the file
	return os.WriteFile(path, []byte(content), 0644)
}

// ========================================================

// ========================================================
// It creates parent directories if needed.
func CreateEmptyFile(path string) (*os.File, error) {
	// Ensure parent directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// Create or truncate the file
	return os.Create(path)
}

// ========================================================

// ========================================================
// CopyFile copies the contents from src to dst.
// It creates the parent directory of dst if needed.
func CopyFile(src, dst string) error {
	// Open source file
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// Ensure destination directory exists
	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Create destination file
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy contents
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	// Ensure contents are flushed
	return out.Sync()
}

// ========================================================

// ========================================================
// DeleteFile deletes the file at the given path and returns an error if it fails.
func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", path, err)
	}
	return nil
}

// ========================================================

// ========================================================
// RenameFile renames (or moves) a file from oldPath to newPath.
// It creates the destination directory if needed.
func RenameFile(oldPath, newPath string) error {
	// Ensure parent directory of newPath exists
	dir := filepath.Dir(newPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Rename the file
	return os.Rename(oldPath, newPath)
}

// ========================================================
