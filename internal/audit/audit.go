package audit

import (
	"fmt"
	"frontend_go/internal/model"
	"frontend_go/internal/store"
)

type Logger struct{ Store *store.Store }

func New(s *store.Store) *Logger { return &Logger{Store: s} }
func (l *Logger) Log(entity, id, op, actor string) error {
	a := model.NewAudit(fmt.Sprintf("%s-%s", id, op), entity, id, op, actor)
	return l.Store.PutAudit(a)
}
func (l *Logger) Read(id string) (model.Audit, error) { return l.Store.GetAudit(id) }
func (l *Logger) Verify(a model.Audit) bool           { return a.Valid() && a.Operation != "" }
