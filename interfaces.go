package main

import (
	"container/list"
)

type searchable interface {
	numStates() int
	complete() bool
	id() int
	reachable() []searchable
}

type searcher interface {
	search(searchable) *list.List
}
