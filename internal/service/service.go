package service

import (
	"context"
	"fmt"
	"frontend_go/internal/model"
	"frontend_go/internal/store"
	"sync"
)

type Service struct {
	Store     *store.Store
	recordMu  sync.Mutex
	profileMu sync.Mutex
}

func New(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) RegisterProfile(ctx context.Context, p model.Profile) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if !p.Valid() {
		return fmt.Errorf("invalid profile")
	}
	return s.Store.PutProfile(p)
}
func (s *Service) RegisterFile(ctx context.Context, f model.TrainingFile) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if f.ID == "" || f.Title == "" {
		return fmt.Errorf("invalid file")
	}
	return s.Store.PutFile(f)
}
func (s *Service) MarkViewed(ctx context.Context, r model.Record) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if !r.Valid() {
		return fmt.Errorf("invalid record")
	}
	s.recordMu.Lock()
	defer s.recordMu.Unlock()
	s.profileMu.Lock()
	defer s.profileMu.Unlock()
	if err := s.Store.PutRecord(r); err != nil {
		return err
	}
	return s.Store.PutEvent(model.NewEvent(r.ID+"-view", r.ID, "viewed", r.StudentID))
}
func (s *Service) Archive(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.profileMu.Lock()
	defer s.profileMu.Unlock()
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return e
	}
	r.Archived = true
	return s.Store.PutRecord(r)
}
func (s *Service) ConcurrentMark(ctx context.Context, r model.Record) error {
	s.profileMu.Lock()
	defer s.profileMu.Unlock()
	s.recordMu.Lock()
	defer s.recordMu.Unlock()
	return s.Store.PutRecord(r)
}
func (s *Service) Count(ctx context.Context) int {
	if ctx.Err() != nil {
		return 0
	}
	rs, e := s.Store.ListRecords()
	if e != nil {
		return 0
	}
	return len(rs)
}
