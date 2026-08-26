package metrics

import (
	"sync"
	"time"
)

type Counter struct {
	mu     sync.Mutex
	values map[string]int
}

func New() *Counter                       { return &Counter{values: map[string]int{}} }
func (c *Counter) Inc(name string)        { c.mu.Lock(); c.values[name]++; c.mu.Unlock() }
func (c *Counter) Add(name string, n int) { c.mu.Lock(); c.values[name] += n; c.mu.Unlock() }
func (c *Counter) Get(name string) int    { c.mu.Lock(); defer c.mu.Unlock(); return c.values[name] }
func (c *Counter) Snapshot() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	o := map[string]int{}
	for k, v := range c.values {
		o[k] = v
	}
	return o
}
func Since(t time.Time) time.Duration { return time.Since(t) }
