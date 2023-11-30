package main

import (
	"fmt"
	"graf"
)

func main() {
	graph, err := graf.FromFile("./data/example.gr", "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = graf.ExportFile(graph, "example.gr"); err != nil {
		fmt.Println(err.Error())
		return
	}

}
