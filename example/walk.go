package main

import (
	"fmt"
	"graf"
	"graf/impl/walk"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) <= 3 {
		return
	}

	fileName := os.Args[1]
	source := os.Args[2]

	step, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	graph, err := graf.LoadFile(fileName, "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rWalk, err := walk.RandomWalk(*graph, source, step)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Random Walk: ")
	for i, v := range rWalk.PredecessorChain {
		if i != len(rWalk.PredecessorChain)-1 {
			fmt.Printf("%s -> ", v.Id)
			continue
		}
		fmt.Println(v.Id)
	}
	fmt.Println("-----")

	minWalk, err := walk.MinWalk(*graph, source, step)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Min Walk: ")
	for i, v := range minWalk.PredecessorChain {
		if i != len(minWalk.PredecessorChain)-1 {
			fmt.Printf("%s -> ", v.Id)
			continue
		}
		fmt.Println(v.Id)
	}
	fmt.Println("-----")

	maxWalk, err := walk.MaxWalk(*graph, source, step)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Max Walk: ")
	for i, v := range maxWalk.PredecessorChain {
		if i != len(maxWalk.PredecessorChain)-1 {
			fmt.Printf("%s -> ", v.Id)
			continue
		}
		fmt.Println(v.Id)
	}

}
