package envutils

import (
	"os"
	"regexp"
)

func ReplaceTemplateVars(content string) string {
	re := regexp.MustCompile(`\{\{\s*(\w+)\s*\}\}`)

	return re.ReplaceAllStringFunc(content, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return ""
		}
		return os.Getenv(submatches[1])
	})
}
