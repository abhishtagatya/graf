package main

import (
	"fmt"
	"graf"
	"graf/impl/floydwarshall"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		return
	}

	fileName := os.Args[1]

	graph, err := graf.LoadFile(fileName, "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	allPairs, err := floydwarshall.FloydWarshall(*graph)
	fmt.Println(allPairs)
	return
}
