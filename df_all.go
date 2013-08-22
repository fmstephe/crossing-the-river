package main

import(
	"fmt"
)

func depthFirstAll(s searchable, path []searchable, visited []bool) [][]searchable {
	successes := make([][]searchable, 0)
	if visited[s.id()] == true {
		return successes
	}
	visited[s.id()] = true
	path = append(path, s)
	if s.complete() {
		return append(successes, path)
	} else {
		nextStates := s.reachable()
		for _, ns := range nextStates {
			ss := depthFirstAll(ns, path, visited)
			successes = append(successes, ss...)
		}
	}
	visited[s.id()] = false
	return successes
}

func depthFirstFirst(s searchable, path []searchable, visited []bool) {
	if visited[s.id()] == true {
		return
	}
	visited[s.id()] = true
	path = append(path, s)
	if s.complete() {
		println(fmt.Sprintf("%v", path))
	} else {
		nextStates := s.reachable()
		for _, ns := range nextStates {
			depthFirstFirst(ns, path, visited)
		}
	}
	visited[s.id()] = false
}

func breadthFirstShortest(s searchable, path []searchable, visited []bool) {
	if visited[s.id()] == true {
		return
	}
	visited[s.id()] = true
	path = append(path, s)
	if s.complete() {
		println(fmt.Sprintf("%v", path))
	} else {
		nextStates := s.reachable()
		for _, ns := range nextStates {
			breadthFirstShortest(ns, path, visited)
		}
	}
	visited[s.id()] = false
}
