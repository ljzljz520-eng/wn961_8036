package frontend_go

import (
	"context"
	"frontend_go/internal/model"
	"frontend_go/internal/service"
	"frontend_go/internal/store"
	"frontend_go/internal/workflow"
	"path/filepath"
	"testing"
	"time"
)

func TestBusinessChain22(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	svc := service.New(s)
	c := workflow.New(svc)
	r := model.NewRecord("r", "f", "s", "view")
	done := make(chan error, 2)
	go func() { done <- svc.ConcurrentMark(context.Background(), r) }()
	go func() { done <- svc.MarkViewed(context.Background(), r) }()
	for i := 0; i < 2; i++ {
		select {
		case <-done:
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("concurrent mark did not complete")
		}
	}
	_ = c
}
