package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func FormatJSON(input []byte) ([]byte, error) {
	var prettyJSON bytes.Buffer

	err := json.Indent(&prettyJSON, input, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	return prettyJSON.Bytes(), nil
}
