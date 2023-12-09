package graf

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var XFSpace = "c \n"
var XFContent = "a %s %s %f\n"

var XFBoilerplate = []string{
	"c Made with Graf (Graph Algorithms Library in Go)\n",
	"c https://github.com/abhishtagatya/graf\n",
}

var XFMeta = []string{
	"p sp %d %d\n",
	"c graph contains %d nodes and %d edges\n",
}

// LoadFile Read Graph from (.gr) File
func LoadFile(fileName string, prefix string) (*Graph, error) {

	graph := BlankGraph()

	readFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lineScan := fileScanner.Text()
		if strings.HasPrefix(lineScan, prefix) {
			lineVertex := strings.Split(lineScan, " ")

			var v1, v2 Vertex
			var weight float64

			v1 = graph.AddVertex(lineVertex[1])
			v2 = graph.AddVertex(lineVertex[2])

			if weight, err = strconv.ParseFloat(lineVertex[3], 32); err == nil {
				graph.AddEdge(v1, v2, weight)
			}
		}
	}

	err = readFile.Close()
	if err != nil {
		return nil, err
	}

	return &graph, nil
}

func ExportFile(graph *Graph, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	// Write Boilerplate
	for _, sb := range XFBoilerplate {
		if _, err = f.WriteString(sb); err != nil {
			return err
		}
	}

	if _, err = f.WriteString(XFSpace); err != nil {
		return err
	}

	// Write Meta
	for _, sm := range XFMeta {
		_, err = f.WriteString(
			fmt.Sprintf(sm, len(graph.Vertices), CountEdge(graph)),
		)
		if err != nil {
			return err
		}
	}

	if _, err = f.WriteString(XFSpace); err != nil {
		return err
	}

	// Write Content
	for cv, edges := range graph.Edges {
		for _, uv := range edges {
			_, err = f.WriteString(
				fmt.Sprintf(XFContent, cv.Id, uv.ConnectedId, uv.Weight),
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
