package lifecycle

import (
	"context"
	"fmt"
	"frontend_go/internal/model"
	"sync"
)

type Machine struct {
	mu     sync.Mutex
	states map[string]string
}

func New() *Machine { return &Machine{states: map[string]string{}} }
func (m *Machine) Start(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.states[id] != "" {
		return false
	}
	m.states[id] = "received"
	return true
}
func (m *Machine) Advance(id, state string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.states[id]; !ok {
		return fmt.Errorf("unknown")
	}
	m.states[id] = state
	return nil
}
func (m *Machine) State(id string) string  { m.mu.Lock(); defer m.mu.Unlock(); return m.states[id] }
func (m *Machine) Complete(id string) bool { return m.Advance(id, "complete") == nil }
func Check(ctx context.Context, r model.Record) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if !r.Valid() {
		return fmt.Errorf("invalid")
	}
	return nil
}
