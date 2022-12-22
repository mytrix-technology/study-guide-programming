package main

// Handler
type Department interface {
	execute(*Patient)
	setNext(Department)
}
