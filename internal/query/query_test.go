package query

import (
	"frontend_go/internal/model"
	"frontend_go/internal/store"
	"path/filepath"
	"testing"
)

func TestFindRecords(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	_ = s.PutRecord(model.NewRecord("r", "f", "Alice", "view"))
	q := New(s)
	rs, e := q.FindRecords("alice")
	if e != nil || len(rs) != 1 {
		t.Fatal(e, len(rs))
	}
}
