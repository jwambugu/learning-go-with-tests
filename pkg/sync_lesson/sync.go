package sync_lesson

import "sync"

// Notes
// - Use channels when passing ownership of data
// - Use mutexes for managing state

// Counter will increment a number.
type Counter struct {
	value int
	mu    sync.Mutex
}

// NewCounter returns a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the count.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

// Value returns the current count.
func (c *Counter) Value() interface{} {
	return c.value
}
