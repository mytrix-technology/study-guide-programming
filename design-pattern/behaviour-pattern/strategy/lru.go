package main

import "fmt"

// Least Recently Used
type Lru struct {
}

// Concrete Strategy
func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}
