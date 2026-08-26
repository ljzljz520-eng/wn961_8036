package api

import (
	"encoding/json"
	"frontend_go/internal/model"
	"frontend_go/internal/service"
	"net/http"
)

type Handler struct{ Svc *service.Service }

func New(s *service.Service) *Handler { return &Handler{Svc: s} }
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var x model.Record
		if json.NewDecoder(r.Body).Decode(&x) != nil {
			http.Error(w, "bad", 400)
			return
		}
		if e := h.Svc.MarkViewed(r.Context(), x); e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		w.WriteHeader(201)
		return
	}
	w.WriteHeader(200)
}
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("ok"))
}
