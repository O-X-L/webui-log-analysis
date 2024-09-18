package api

import (
	"encoding/json"
	"net/http"

	"github.com/O-X-L/webui-log-analysis/internal/u"
)

func errorResponse(w http.ResponseWriter, s int, e string) {
	w.WriteHeader(s)
	json.NewEncoder(w).Encode(map[string]string{"error": e})
}

/*
func badRequestResponse(w http.ResponseWriter, e string) {
	errorResponse(w, http.StatusBadRequest, e)
}
*/

func failureResponse(w http.ResponseWriter, e string) {
	errorResponse(w, http.StatusInternalServerError, e)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		u.LogError("api-test", err)
		failureResponse(w, "Failed to JSON-encode data")
	}
}
