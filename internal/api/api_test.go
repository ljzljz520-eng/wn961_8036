package api

import (
	"frontend_go/internal/service"
	"frontend_go/internal/store"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestHTTPMark(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r := httptest.NewRequest("POST", "/", strings.NewReader(`{"ID":"r","FileID":"f","StudentID":"s","Action":"view"}`))
	w := httptest.NewRecorder()
	New(service.New(s)).ServeHTTP(w, r)
	if w.Code != 201 {
		t.Fatal(w.Code)
	}
}
