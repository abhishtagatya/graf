package main

import (
	"fmt"
	"github.com/abhishtagatya/graf"
	"github.com/abhishtagatya/graf/impl/dijkstra"
	"os"
	"time"
)

func main() {

	if len(os.Args) <= 3 {
		return
	}

	fileName := os.Args[1]
	source := os.Args[2]
	target := os.Args[3]

	graph, err := graf.LoadFile(fileName, "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	startTime := time.Now()
	report, err := dijkstra.SingleDijkstra(*graph, source, target)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	result := fmt.Sprintf("Distance\t: %f\nVisits\t: %d\nElapsed Timet\t: %s", report.Distance, len(report.VisitMap), elapsedTime)

	fmt.Println("Dijkstra : ")
	fmt.Println(result)
}
