package main

import (
	"fmt"
	"graf"
	"graf/impl/astarh"
	"os"
	"time"
)

func main() {

	if len(os.Args) <= 3 {
		return
	}

	fileName := os.Args[1]
	heuristicFile := os.Args[2]
	source := os.Args[3]
	target := os.Args[4]

	graph, err := graf.LoadFile(fileName, "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	heuristic, err := astarh.LoadCoordinate(heuristicFile, "v")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	startTime := time.Now()
	report, err := astarh.AStarHaversine(*graph, source, target, heuristic)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	result := fmt.Sprintf("Distance\t: %f\nVisits\t: %d\nElapsed Timet\t: %s", report.Distance, len(report.VisitMap), elapsedTime)

	fmt.Println("A* Haversine : ")
	fmt.Println(result)
}
