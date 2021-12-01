package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Scaffold generates a Solution directory structure for a day.
func Scaffold(templateDir string, outputDir string, variables interface{}) error {
	ensureDirectory(outputDir)

	walkErr := filepath.Walk(
		templateDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				ensureDirectory(path)

				return nil
			}

			dir := filepath.Join(
				outputDir,
				strings.TrimPrefix(
					strings.TrimSuffix(path, filepath.Base(path)),
					templateDir,
				),
			)
			ensureDirectory(dir)

			targetFile := strings.TrimSuffix(filepath.Join(dir, filepath.Base(path)), ".tmpl")

			outputFile, err := os.Create(targetFile)
			if err != nil {
				return fmt.Errorf("unable to create output file: %w", err)
			}

			defer outputFile.Close()

			tmpl, err := template.ParseFiles(path)
			if err != nil {
				return fmt.Errorf("unable to parse template file: %w", err)
			}

			tmplErr := tmpl.Execute(outputFile, variables)
			if tmplErr != nil {
				return fmt.Errorf("template execution failed: %w", tmplErr)
			}

			return nil
		},
	)

	if walkErr != nil {
		return fmt.Errorf("something went wrong with scaffolding: %w", walkErr)
	}

	return nil
}
