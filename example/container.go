package main

import (
	"fmt"
	"graf"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		return
	}

	arg := os.Args[1]
	if arg == "" {
		return
	}

	graph, err := graf.FromFile(arg, "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	auxContainer := graf.ComputeContainers(graph)

	fmt.Println("Saving auxContainer...")
	err = graf.ExportAuxContainer(auxContainer, fmt.Sprintf("%s.aux", arg))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
