package main

import (
	"fmt"
	"github.com/abhishtagatya/graf"
	"github.com/abhishtagatya/graf/impl/geocontainer"
	"os"
	"time"
)

func main() {

	if len(os.Args) <= 3 {
		return
	}

	fileName := os.Args[1]
	containerFile := os.Args[2]
	source := os.Args[3]
	target := os.Args[4]

	graph, err := graf.LoadFile(fileName, "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	container, err := geocontainer.LoadGeoContainer(containerFile, "x")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	startTime := time.Now()
	report, err := geocontainer.DijkstraGeometricPrune(*graph, source, target, container)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	result := fmt.Sprintf("Distance\t: %f\nVisits\t: %d\nElapsed Timet\t: %s", report.Distance, len(report.VisitMap), elapsedTime)

	fmt.Println("Dijkstra Geo. Prune : ")
	fmt.Println(result)
}
