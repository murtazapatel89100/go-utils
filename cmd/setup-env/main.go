package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/murtazapatel89100/go-utils/pkg/envutils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: setup-env <app-directory>")
		os.Exit(1)
	}

	appDir := os.Args[1]
	templatePath := filepath.Join(appDir, "env.test")
	outputPath := filepath.Join(appDir, ".env")

	content, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Println("Could not read env.test")
		os.Exit(1)
	}

	result := envutils.ReplaceTemplateVars(string(content))

	err = os.WriteFile(outputPath, []byte(result), 0644)
	if err != nil {
		fmt.Println("Failed to write .env")
		os.Exit(1)
	}

	fmt.Println(".env created for pipeline:", appDir)
}
