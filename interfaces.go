package main

type searchable interface {
	numStates() int
	complete() bool
	id() int
	reachable() []searchable
}
