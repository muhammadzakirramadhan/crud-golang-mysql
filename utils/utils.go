package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, data interface{}, status int) {
	toByte, err := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "Error!", http.StatusBadRequest)
	}

	w.WriteHeader(status)
	w.Write([]byte(toByte))
}
