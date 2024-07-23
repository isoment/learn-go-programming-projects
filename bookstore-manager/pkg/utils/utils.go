package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
Take the request json body and unmarshal it into x. This could be anything but in our case
will likely be a model.
*/
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
