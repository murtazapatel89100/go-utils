package main

import (
	"fmt"
	"os"

	"github.com/murtazapatel89100/go-utils/pkg/formatter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jsonfmt <file.json>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	formatted, err := formatter.FormatJSON(content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(string(formatted))
}
