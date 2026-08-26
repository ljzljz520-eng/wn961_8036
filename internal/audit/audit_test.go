package audit

import (
	"frontend_go/internal/store"
	"path/filepath"
	"testing"
)

func TestAuditLog(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	l := New(s)
	if e := l.Log("Record", "r", "review", "u"); e != nil {
		t.Fatal(e)
	}
	a, e := l.Read("r-review")
	if e != nil || !l.Verify(a) {
		t.Fatal(e)
	}
}
