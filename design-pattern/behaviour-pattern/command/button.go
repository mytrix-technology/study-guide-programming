package main

type Button struct {
	command Command
}

// Invoker
func (b *Button) press() {
	b.command.execute()
}
