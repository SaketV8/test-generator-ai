package promptgen

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Data struct for the template
type FileData struct {
	Path    string
	Content string
}

// folder to exclude, like prompts folderr
var excludedDirs = map[string]bool{
	".git":       true,
	"tmp":        true,
	"prompts":    true,
	"database":   true,
	"router":     true,
	"middleware": true,
	"utils":      true,
	"test":       true,
}

// read all files from folder
func ReadFolder(folder string) ([]FileData, error) {
	var files []FileData

	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			// skipping directories which I want
			if excludedDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		files = append(files, FileData{
			// Path:    strings.TrimSuffix(d.Name(), filepath.Ext(d.Name())),
			Path:    filepath.ToSlash(path),
			Content: strings.TrimSpace(string(data)),
		})

		return nil
	})

	return files, err
}

// read all files from folder
func ReadTestFolder(folder string) ([]FileData, error) {
	var files []FileData

	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		files = append(files, FileData{
			// Path:    strings.TrimSuffix(d.Name(), filepath.Ext(d.Name())),
			Path:    filepath.ToSlash(path),
			Content: strings.TrimSpace(string(data)),
		})

		return nil
	})

	return files, err
}

// ===================================================================
// Template based prompts generator (-_-)
// ===================================================================
const textTemplate = `
Complete files here:

{{range .}}
--------------------
{{.Path}}
--------------------
{{.Content}}
{{end}}
`

// template output to path/file.yaml
func WriteTemplateToFile(data []FileData, path string) error {
	outFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outFile.Close()

	tmpl, err := template.New("code-content").Parse(textTemplate)
	if err != nil {
		return err
	}

	return tmpl.Execute(outFile, data)
}

func WriteTemplateToString(data []FileData) (string, error) {
	var builder strings.Builder

	tmpl, err := template.New("code-content").Parse(textTemplate)
	if err != nil {
		return "", err
	}

	if err := tmpl.Execute(&builder, data); err != nil {
		return "", err
	}

	return builder.String(), nil
}
