package main

import(
)

const (
	cabbage = items(1)
	goat = items(2)
	tiger = items(4)
	man = items(8)
	all = cabbage | goat | tiger | man
	none = items(0)
)

type items uint32

func (i items) numStates() int {
	return 16
}

func (i items) complete() bool {
	return i == none
}

func (i items) id() int {
	return int(i)
}

func (i items) legal() bool {
	if  i.has(man) {
		return true
	}
	if i.has(cabbage) && i.has(goat) {
		return false
	}
	if i.has(goat) && i.has(tiger) {
		return false
	}
	return true
}

func (i items) has(h items) bool {
	return (i & h) == h
}

func (i items) add(a items) items {
	return i | a
}

func (i items) remove(r items) items {
	return i ^ r
}

func (i items) complement() items {
	return i ^ all
}

func (i items) String() string {
	other := i.complement()
	return "{"+i.iString()+","+other.iString()+"}"
}

func (i items) iString() string {
	c := ""
	g := ""
	t := ""
	m := ""
	if i.has(cabbage) {
		c = "c,"
	}
	if i.has(goat) {
		g = "g,"
	}
	if i.has(tiger) {
		t = "t,"
	}
	if i.has(man) {
		m = "m,"
	}
	return "("+c+g+t+m+")"
}

func complementAll(is []searchable) []searchable {
	newIs := make([]searchable, len(is))
	for i := range is {
		newIs[i] = (is[i].(items)).complement()
	}
	return newIs
}

func (i items) reachable() []searchable {
	if i.has(man) {
		return next(i)
	}
	i = i.complement()
	return complementAll(next(i))
}

func next(i items) []searchable {
	all := i.all()
	pairs := allPairs(all)
	all = append(all, pairs...)
	nxt := make([]searchable, 0)
	for _, ai := range all {
		newI := i.remove(ai).remove(man)
		if newI.legal() {
			nxt = append(nxt, newI)
		}
	}
	return nxt
}

func (i items) all() []items {
	is := make([]items, 0)
	if i.has(cabbage) {
		is = append(is, cabbage)
	}
	if i.has(goat) {
		is = append(is, goat)
	}
	if i.has(tiger) {
		is = append(is, tiger)
	}
	is = append(is, none)
	return is
}

func allPairs(is []items) []items {
	pairs := make([]items, 0)
	for i := 0; i < len(is); i++ {
		for j := i + 1; j < len(is); j++ {
			pairs = append(pairs, is[i].add(is[j]))
		}
	}
	return pairs
}

