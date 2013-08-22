package main

import (
	"fmt"
)

func main() {
	mySearcher := &breadthFirstShortest{}
	success := mySearcher.search(all)
	for e := success.Front(); e != nil; e = e.Next() {
		path := e.Value.([]searchable)
		println(fmt.Sprintf("%v", path))
	}
	println("Found", success.Len(), "unique solutions")
}
