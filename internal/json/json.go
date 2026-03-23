package JSON

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "applicaton")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
