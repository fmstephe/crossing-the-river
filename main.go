package main

import (
	"fmt"
)

func main() {
	successes := depthFirstAll(all, make([]searchable, 0), make([]bool, all.numStates()))
	for _, ss := range successes {
		println(fmt.Sprintf("%v", ss))
	}
}
