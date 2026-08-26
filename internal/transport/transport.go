package transport

import (
	"encoding/json"
	"frontend_go/internal/model"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func ReadRecord(r *http.Request) (model.Record, error) {
	var v model.Record
	e := json.NewDecoder(r.Body).Decode(&v)
	return v, e
}
func MethodAllowed(r *http.Request, methods ...string) bool {
	for _, m := range methods {
		if r.Method == m {
			return true
		}
	}
	return false
}
func Query(r *http.Request, key string) string { return r.URL.Query().Get(key) }
func Error(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, map[string]string{"error": msg})
}
