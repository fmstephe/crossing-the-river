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
