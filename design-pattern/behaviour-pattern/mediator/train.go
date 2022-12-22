package main

// Component Interface
type Train interface {
	arrive()
	depart()
	permitArrival()
}
