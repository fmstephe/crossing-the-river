package main

import(
	"fmt"
)

const (
	cabbage = location(1)
	goat = location(2)
	tiger = location(4)
	man = location(8)
	all = cabbage | goat | tiger | man
	none = location(0)
)

type location uint32

func (s location) complete() bool {
	end := s.end
	return has(end, cabbage) && has(end, goat) && has(end, tiger) && has(end, man)
}

func (l location) legal() bool {
	if  l.has(man) {
		return true
	}
	if l.has(cabbage) && l.has(goat) {
		return false
	}
	if l.has(goat) && l.has(tiger) {
		return false
	}
	return true
}

func (l location) has(item uint8) bool {
	return (side & item) != 0
}

func (l location) add(item uint8) location {
	return side | item
}

func (l location) remove(item uint8) location {
	return side ^ item
}

func (l location) complement() location {
	return l ^ all
}

func allPairs(items []location) []location {
	pairs := make([]location, 0)
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			pairs = append(pairs, items[i].add(items[j]))
		}
	}
	return pairs
}

func (l location) all() []location {
	items := make([]uint8, 0)
	if has(side, cabbage) {
		items = append(items, cabbage)
	}
	if has(side, goat) {
		items = append(items, goat)
	}
	if has(side, tiger) {
		items = append(items, tiger)
	}
	return items
}

func complementAll(s []location) []location {
	newS := make([]location, len(s))
	for i := range s {
		newS[i] = s[i].complement()
	}
	return newS
}

var visited [16]bool

func main() {
	search(all, make([]location, 0))
}

func search(l location, path []location) {
	visited[l.start] = true
	path = append(path, l)
	if l.complete() {
		println(fmt.Sprintf("Complete %v", path))
	} else {
		states := generateNext(l)
		for _, s := range states {
			search(s, path)
		}
	}
	visited[l.start] = false
}

func generateNext(s location) []location {
	if has(s.start, man) {
		return next(s.start, s.end)
	}
	return reverse(next(s.end, s.start))
}

func next(manSide, otherSide uint8) []location {
	states := make([]location, 0)
	items := all(manSide)
	println(len(items))
	pairs := allPairs(items)
	println(len(pairs))
	// Move single items
	for _, item := range items {
		oldManSide := remove(remove(manSide, item), man)
		newManSide := add(add(otherSide, item), man)
		if legal(oldManSide) && legal(newManSide) {
			states = append(states, location{oldManSide, newManSide})
		}
	}
	// Move paired items
	for _, pair := range pairs {
		oldManSide := remove(remove(remove(manSide, pair[0]), pair[1]), man)
		newManSide := add(add(add(otherSide, pair[0]), pair[1]), man)
		if legal(oldManSide) && legal(newManSide) {
			states = append(states, location{oldManSide, newManSide})
		}
	}
	return states
}
