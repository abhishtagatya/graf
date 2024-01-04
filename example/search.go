package main

import (
	"fmt"
	"github.com/abhishtagatya/graf"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 2 {
		return
	}

	fileName := os.Args[1]
	source := os.Args[2]

	graph, err := graf.LoadFile(fileName, "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	startTime := time.Now()
	report, err := graf.BreadthSearch(*graph, source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	result := fmt.Sprintf("Distance\t: %f\nVisits\t: %d\nElapsed Timet\t: %s", report.Distance, len(report.VisitMap), elapsedTime)

	fmt.Println("BFS : ")
	fmt.Println(result)

	for _, v := range report.PredecessorChain {
		fmt.Printf("%s ", v.Id)
	}
	fmt.Println()
	fmt.Println(report.StartVertex.Id, report.EndVertex.Id)
}
