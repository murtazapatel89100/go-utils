package prettier

import (
	"encoding/json"
	"fmt"
	"os"
)

var prettierConfigTemplate = map[string]any{

	"semi":           false,
	"singleQuote":    true,
	"tabWidth":       2,
	"trailingComma":  "es5",
	"printWidth":     100,
	"bracketSpacing": true,
	"jsxSingleQuote": false,
	"arrowParens":    "always",
	"endOfLine":      "lf",
}

func GeneratePrettierConfig() error {
	data, err := json.MarshalIndent(prettierConfigTemplate, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(".prettierrc", data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write .prettierrc: %w", err)
	}
	return nil
}
