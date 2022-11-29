package gobe

import "encoding/json"

// Turn any model into JSON
func ToJSON(model any) string {
	res, err := json.MarshalIndent(model, "", " ")
	if err != nil {
		return ""
	}
	return string(res)
}
