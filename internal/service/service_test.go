package service

import (
	"context"
	"frontend_go/internal/model"
	"frontend_go/internal/store"
	"path/filepath"
	"testing"
)

func TestMarkViewed(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	if e := New(s).MarkViewed(context.Background(), model.NewRecord("r", "f", "s", "view")); e != nil {
		t.Fatal(e)
	}
}
