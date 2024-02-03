package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseJSONRequestBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}
