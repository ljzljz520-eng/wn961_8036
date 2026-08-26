package query

import (
	"frontend_go/internal/model"
	"frontend_go/internal/store"
	"strings"
)

type Engine struct{ Store *store.Store }

func New(s *store.Store) *Engine { return &Engine{Store: s} }
func (e *Engine) FindRecords(term string) ([]model.Record, error) {
	rs, err := e.Store.ListRecords()
	if err != nil {
		return nil, err
	}
	out := make([]model.Record, 0)
	for _, r := range rs {
		if strings.Contains(strings.ToLower(r.ID+" "+r.StudentID+" "+r.Action), strings.ToLower(term)) {
			out = append(out, r)
		}
	}
	return out, nil
}
func (e *Engine) FindFile(id string) (model.TrainingFile, error) { return e.Store.GetFile(id) }
func (e *Engine) StudentHistory(id string) ([]model.Record, error) {
	rs, err := e.Store.ListRecords()
	if err != nil {
		return nil, err
	}
	out := []model.Record{}
	for _, r := range rs {
		if r.StudentID == id {
			out = append(out, r)
		}
	}
	return out, nil
}
func (e *Engine) Recent(limit int) ([]model.Record, error) {
	rs, err := e.Store.ListRecords()
	if err != nil {
		return nil, err
	}
	if limit < 0 {
		limit = 0
	}
	if limit > len(rs) {
		limit = len(rs)
	}
	return rs[len(rs)-limit:], nil
}
