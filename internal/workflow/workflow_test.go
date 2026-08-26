package workflow

import (
	"context"
	"frontend_go/internal/model"
	"frontend_go/internal/service"
	"frontend_go/internal/store"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := New(service.New(s))
	if e := c.Intake(context.Background(), model.NewProfile("p", "student"), model.TrainingFile{ID: "f", Title: "brakes", Category: "pdf"}); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := New(service.New(s))
	r := model.NewRecord("r", "f", "s", "view")
	if e := c.Process(context.Background(), r); e != nil {
		t.Fatal(e)
	}
	if e := c.Archive(context.Background(), "r"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := New(service.New(s))
	r := model.NewRecord("r", "f", "s", "view")
	if e := c.Submit(context.Background(), r); e != nil {
		t.Fatal(e)
	}
	if e := c.Review(context.Background(), "r", "reviewer"); e != nil {
		t.Fatal(e)
	}
	if e := c.Notify(context.Background(), "r"); e != nil {
		t.Fatal(e)
	}
	if !c.Track(context.Background(), "r") {
		t.Fatal()
	}
}
