package helpers

import "encoding/json"

type JMap map[string]any

func (j JMap) Json() string {
	val, _ := json.Marshal(j)
	return string(val)
}
