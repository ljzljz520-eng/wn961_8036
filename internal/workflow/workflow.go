package workflow

import (
	"context"
	"fmt"
	"frontend_go/internal/model"
	"frontend_go/internal/service"
)

type Coordinator struct{ Svc *service.Service }

func New(s *service.Service) *Coordinator { return &Coordinator{Svc: s} }
func (c *Coordinator) Intake(ctx context.Context, p model.Profile, f model.TrainingFile) error {
	if e := c.Svc.RegisterProfile(ctx, p); e != nil {
		return e
	}
	if e := c.Svc.RegisterFile(ctx, f); e != nil {
		return e
	}
	return nil
}
func (c *Coordinator) Process(ctx context.Context, r model.Record) error {
	if r.Action == "" {
		return fmt.Errorf("action required")
	}
	return c.Svc.MarkViewed(ctx, r)
}
func (c *Coordinator) Archive(ctx context.Context, id string) error { return c.Svc.Archive(ctx, id) }
func (c *Coordinator) Lookup(ctx context.Context, id string) (model.Record, error) {
	return c.Svc.Store.GetRecord(id)
}
func (c *Coordinator) Submit(ctx context.Context, r model.Record) error {
	if e := c.Process(ctx, r); e != nil {
		return e
	}
	return c.Svc.Store.PutAudit(model.NewAudit(r.ID+"-submit", "Record", r.ID, "submit", r.StudentID))
}
func (c *Coordinator) Review(ctx context.Context, id, actor string) error {
	r, e := c.Lookup(ctx, id)
	if e != nil {
		return e
	}
	if r.Archived {
		return fmt.Errorf("archived")
	}
	return c.Svc.Store.PutAudit(model.NewAudit(id+"-review", "Record", id, "review", actor))
}
func (c *Coordinator) Notify(ctx context.Context, id string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return c.Svc.Store.PutEvent(model.NewEvent(id+"-notify", id, "notify", "system"))
}
func (c *Coordinator) Track(ctx context.Context, id string) bool {
	_, e := c.Lookup(ctx, id)
	return e == nil
}
