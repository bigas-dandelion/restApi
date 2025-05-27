package res

import (
	"encoding/json"
	"net/http"
)

// any -> interface{}
func Json(w http.ResponseWriter, data interface{}, statusCdoe int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCdoe)
	json.NewEncoder(w).Encode(data)
}
