package JSON

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "applicaton/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func Read(r *http.Request, data any) error {
	reqPayload := r.Body
	json.NewDecoder(reqPayload).DisallowUnknownFields()
	return json.NewDecoder(reqPayload).Decode(data)
}
