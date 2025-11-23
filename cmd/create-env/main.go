package main

import (
	"fmt"
	"os"

	"github.com/murtazapatel89100/go-utils/pkg/envutils"
)

func main() {
	err := envutils.CopyTemplate("env.template", ".env")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(".env created from env.template")
}
