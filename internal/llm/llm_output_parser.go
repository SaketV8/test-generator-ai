package llm

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/SaketV8/test-generator-ai/internal/models"
)

// func ParseLLmOutPut(llmOutput string) []models.FileToGenerate {
func ParseLLmOutPut(llmOutput string) []models.FileData {
	// --- Parsing Stage with the Corrected Regex ---
	// utils.PrintBoxedText("--- Parsing LLM Output ---")
	filesData := []models.FileData{}

	re := regexp.MustCompile(`\[START OF FILE: (.*?)]\s*([\s\S]*?)\s*\[END OF FILE: (.*?)]`)
	matches := re.FindAllStringSubmatch(llmOutput, -1)

	if len(matches) == 0 {
		fmt.Println("Could not find any file content blocks in the LLM output.")
		return nil
	}

	for _, match := range matches {
		startPath := strings.TrimSpace(match[1])
		content := strings.TrimSpace(match[2])
		endPath := strings.TrimSpace(match[3])

		// **MANUAL CHECK:** We now perform the check in Go code.
		if startPath != endPath {
			fmt.Printf("Mismatched tags found: START tag was '%s' but END tag was '%s'. Skipping this block.\n", startPath, endPath)
			continue // Skip this malformed block
		}

		file := models.FileData{
			Path:    startPath,
			Content: content,
		}
		filesData = append(filesData, file)
		fmt.Printf("Queued file for writing: %s\n", file.Path)
	}

	fmt.Println("--------------------------")

	return filesData
}
