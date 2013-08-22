package main

import(
	"container/list"
)

type depthFirstAll struct {}

func (d *depthFirstAll) search(s searchable) *list.List {
	success := list.New()
	idepthFirstAll(all, make([]searchable, 0), make([]bool, all.numStates()), success)
	return success
}

func idepthFirstAll(s searchable, path []searchable, visited []bool, success *list.List) {
	if visited[s.id()] == true {
		return
	}
	visited[s.id()] = true
	path = append(path, s)
	if s.complete() {
		sPath := make([]searchable, len(path))
		copy(sPath, path)
		success.PushFront(sPath)
	} else {
		nxt := s.reachable()
		for _, ns := range nxt {
			idepthFirstAll(ns, path, visited, success)
		}
	}
	visited[s.id()] = false
}

type depthFirstFirst struct {}

func (d *depthFirstFirst) search(s searchable) *list.List {
	success := list.New()
	idepthFirstFirst(all, make([]searchable, 0), make([]bool, all.numStates()), success)
	return success
}

func idepthFirstFirst(s searchable, path []searchable, visited []bool, success *list.List) bool {
	if visited[s.id()] == true {
		return false
	}
	visited[s.id()] = true
	path = append(path, s)
	if s.complete() {
		sPath := make([]searchable, len(path))
		copy(sPath, path)
		success.PushFront(sPath)
		return true
	} else {
		nextStates := s.reachable()
		for _, ns := range nextStates {
			if idepthFirstFirst(ns, path, visited, success) {
				return true
			}
		}
	}
	visited[s.id()] = false
	return false
}

type depthFirstShortest struct {}

func (d *depthFirstShortest) search(s searchable) *list.List {
	i := 0
	for {
		i++
		success := list.New()
		idepthFirstShortest(all, make([]searchable, 0), make([]bool, all.numStates()), success, i)
		if success.Len() > 0 {
			return success
		}
	}
	panic("Unreachable")
}

func idepthFirstShortest(s searchable, path []searchable, visited []bool, success *list.List, maxDepth int) bool {
	if visited[s.id()] == true  || len(path) > maxDepth {
		return false
	}
	visited[s.id()] = true
	path = append(path, s)
	if s.complete() {
		sPath := make([]searchable, len(path))
		copy(sPath, path)
		success.PushFront(sPath)
		return true
	} else {
		nextStates := s.reachable()
		for _, ns := range nextStates {
			if idepthFirstShortest(ns, path, visited, success, maxDepth) {
				return true
			}
		}
	}
	visited[s.id()] = false
	return false
}

type breadthFirstShortest struct {}

func (b *breadthFirstShortest) search(s searchable) *list.List {
	paths := list.New()
	paths.PushBack([]searchable{s})
	result := ibreadthFirstShortest(paths)
	success := list.New()
	success.PushBack(result)
	return success
}

func ibreadthFirstShortest(paths *list.List) []searchable {
	front := paths.Front()
	paths.Remove(front)
	path := front.Value.([]searchable)
	s := path[len(path)-1]
	if s.complete() {
		return path
	} else {
		nextStates := s.reachable()
		for _, ns := range nextStates {
			sPath := make([]searchable, len(path)+1)
			copy(sPath, path)
			sPath[len(sPath)-1] = ns
			paths.PushBack(sPath)
		}
	}
	return ibreadthFirstShortest(paths)
}
