package main

// Subject Interface
type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
