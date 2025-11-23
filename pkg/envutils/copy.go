package envutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CopyTemplate(templateFile, outputFile string) error {
	src, err := os.Open(templateFile)
	if err != nil {
		return fmt.Errorf("template not found: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("unable to create env file: %w", err)
	}
	defer dst.Close()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.Contains(line, "=") {
			_, err := dst.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
