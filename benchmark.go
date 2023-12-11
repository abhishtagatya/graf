package graf

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var XBMeta = []string{
	"c benchmark contains %d tests\n",
}

var XBContent = "b %s %s\n"

type BenchmarkSet struct {
	Source string
	Target string
}

func LoadBenchmarkFile(fileName string, prefix string) ([]BenchmarkSet, error) {

	benchmarkSet := make([]BenchmarkSet, 0)

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

			benchmarkSet = append(benchmarkSet, BenchmarkSet{
				Source: lineVertex[1],
				Target: lineVertex[2],
			})
		}
	}

	return benchmarkSet, nil
}

func ExportBenchmarkFile(graph *Graph, fileName string, k int) error {
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
	for _, sm := range XBMeta {
		_, err = f.WriteString(
			fmt.Sprintf(sm, k),
		)
		if err != nil {
			return err
		}
	}

	if _, err = f.WriteString(XFSpace); err != nil {
		return err
	}

	// Write Content
	for i := 0; i < k; i++ {

		sv := getRandomVertex(graph.Vertices)
		uv := getRandomVertex(graph.Vertices)

		_, err = f.WriteString(
			fmt.Sprintf(XBContent, sv, uv),
		)
	}

	return nil
}

func getRandomVertex(m map[string]Vertex) string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	randomIndex := rand.Intn(len(keys))
	return keys[randomIndex]
}
