package main

import "sync"

func main() {
	// entry point
}

type Counter struct {
	mu    sync.Mutex
	value int
}

/*
// alternately
type Counter struct {
	sync.Mutex
	value int
}
// usage - c.Lock()
*/

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
