package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	pkg := flag.String("pkg", "npm", "Package manager to use (npm or pnpm)")
	flag.Parse()

	fmt.Println("Generating .prettierrc...")

	var cmd *exec.Cmd

	switch *pkg {
	case "pnpm":
		cmd = exec.Command("pnpm", "add", "-D", "prettier")
	case "npm":
		cmd = exec.Command("npm", "install", "-D", "prettier")
	default:
		fmt.Println("Invalid package manager. Use npm or pnpm.")
		os.Exit(1)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Installing Prettier using %s...\n", *pkg)
	if err := cmd.Run(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	fmt.Println("Prettier installed successfully.")
}
