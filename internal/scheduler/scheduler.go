package scheduler

import (
	"context"
	"sync"
	"time"
)

type Job func(context.Context) error
type Scheduler struct {
	mu   sync.Mutex
	jobs []Job
}

func New() *Scheduler          { return &Scheduler{jobs: []Job{}} }
func (s *Scheduler) Add(j Job) { s.mu.Lock(); s.jobs = append(s.jobs, j); s.mu.Unlock() }
func (s *Scheduler) Run(ctx context.Context) []error {
	s.mu.Lock()
	jobs := append([]Job{}, s.jobs...)
	s.mu.Unlock()
	out := []error{}
	for _, j := range jobs {
		if e := j(ctx); e != nil {
			out = append(out, e)
		}
	}
	return out
}
func (s *Scheduler) RunAsync(ctx context.Context) chan error {
	ch := make(chan error, 1)
	go func() {
		e := s.Run(ctx)
		if len(e) > 0 {
			ch <- e[0]
		} else {
			ch <- nil
		}
	}()
	return ch
}
func Wait(ctx context.Context, d time.Duration) bool {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-t.C:
		return true
	}
}
