package utility

import "encoding/json"

func Jsonout(v any) string {
	b, _ := json.MarshalIndent(v, " ", "  ")
	return string(b)
}
