package main

import (
	"fmt"
	"graf"
	"graf/impl/minoverlay"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) <= 2 {
		return
	}

	fileName := os.Args[1]
	nodes := os.Args[2]

	graph, err := graf.LoadFile(fileName, "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sub := strings.Split(nodes, " ")
	fmt.Println(sub)

	startTime := time.Now()
	aux, err := minoverlay.MinOverlay(*graph, 0.0, sub)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	result := fmt.Sprintf("Aux\t: %v\nNodes\t: %d\nElapsed Timet\t: %s", aux, len(sub), elapsedTime/time.Duration(len(sub)))

	fmt.Println("Min Overlay : ")
	fmt.Println(result)

}
