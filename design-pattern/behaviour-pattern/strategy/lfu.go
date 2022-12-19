package main

import "fmt"

// Least Frequently Used
type Lfu struct {
}

// Concrete Strategy
func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}
