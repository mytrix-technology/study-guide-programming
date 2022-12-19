package main

// Strategy Interface
type EvictionAlgo interface {
	evict(c *Cache)
}
