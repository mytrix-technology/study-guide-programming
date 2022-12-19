package main

import "fmt"

// First In First Out
type Fifo struct {
}

// Concrete Strategy
func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}
