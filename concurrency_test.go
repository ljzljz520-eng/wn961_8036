package frontend_go

import (
	"context"
	"frontend_go/internal/model"
	"frontend_go/internal/service"
	"frontend_go/internal/store"
	"path/filepath"
	"testing"
)

func TestConcurrentMarks(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	svc := service.New(s)
	ch := make(chan error, 2)
	for i := 0; i < 2; i++ {
		go func() { ch <- svc.MarkViewed(context.Background(), model.NewRecord("r", "f", "s", "view")) }()
	}
	for i := 0; i < 2; i++ {
		if e := <-ch; e != nil {
			t.Fatal(e)
		}
	}
}
