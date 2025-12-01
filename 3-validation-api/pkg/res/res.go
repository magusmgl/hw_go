package res

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, data any, stausCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(stausCode)
	json.NewEncoder(w).Encode(data)
}
