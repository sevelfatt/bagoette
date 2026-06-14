package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeJSONFromRequestBody(r *http.Request, body interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return err
	}
	return nil
}

func RespondJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}