package hopeserver

import (
	"encoding/json"
)

// ReadJSON is used to parse JSON
func ReadJSON(data string, rules []map[string]string) []string {

	results := rules
	json.Unmarshal([]byte(data), &results)

	return results
}
