package auxiliary

import (
	"bytes"
	"container/heap"
	"errors"
	"fmt"
	"graf"
	"log"
	"math"
	"os"
	"strings"
)

/* Auxiliary Graph for Engineering Large Networks
 * Geometric containers for efficient shortest path computation. (Wagner et al., 2005)
 */

var XCContentHead = "x %s %s"
var XCContentTail = " %s"

var XCBoilerplate = []string{
	"c Geometric Container (Aux Graph)\n",
	"c Made with Graf (Graph Algorithms Library in Go)\n",
	"c https://github.com/abhishtagatya/graf\n",
}

type ContainerTuple struct {
	U string
	V string
}

func LoadContainer(fileName string, prefix string) (map[ContainerTuple][]string, error) {
	aux := make(map[ContainerTuple][]string)

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(file)
	for {
		lineScan, err := buf.ReadString('\n')
		if err != nil {
			break
		}

		if strings.HasPrefix(lineScan, prefix) {
			lineVertex := strings.Split(lineScan, " ")

			vTup := ContainerTuple{U: lineVertex[1], V: lineVertex[2]}

			for _, lv := range lineVertex[3:] {
				aux[vTup] = append(aux[vTup], lv)
			}
		}
	}

	return aux, nil
}

func ExportContainer(aux map[ContainerTuple][]string, fileName string) error {

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	// Write Boilerplate
	for _, sb := range XCBoilerplate {
		if _, err = f.WriteString(sb); err != nil {
			return err
		}
	}

	if _, err = f.WriteString(graf.XFSpace); err != nil {
		return err
	}

	// Write Content
	for k, v := range aux {
		_, err = f.WriteString(fmt.Sprintf(XCContentHead, k.U, k.V))
		for _, vx := range v {
			_, err = f.WriteString(fmt.Sprintf(XCContentTail, vx))
		}
		_, err = f.WriteString("\n")
	}

	return nil
}

func ContainsVertex(list []string, v string) bool {
	for _, i := range list {
		if v == i {
			return true
		}
	}

	return false
}

func ComputeContainers(graph *graf.Graph) map[ContainerTuple][]string {
	aMap := make(map[string]ContainerTuple)
	auxContainer := make(map[ContainerTuple][]string)

	for sid := range graph.Vertices {
		sv := graph.Vertices[sid]

		distanceMap := map[string]float64{sid: 0}

		queue := graf.BlankQueue()
		heap.Push(&queue, &graf.QueueItem{
			Value:    sv,
			Priority: 0,
		})

		for !queue.IsEmpty() {
			cq := heap.Pop(&queue).(*graf.QueueItem)
			cv := cq.Value.(graf.Vertex)

			if cv != sv {
				if !ContainsVertex(auxContainer[aMap[cv.Id]], cv.Id) {
					auxContainer[aMap[cv.Id]] = append(auxContainer[aMap[cv.Id]], cv.Id)
				}
			}

			for _, edge := range graph.Edges[cv] {
				newDist := cq.Priority + edge.Weight
				if dist, ok := distanceMap[edge.ConnectedId]; !ok || newDist < dist {
					distanceMap[edge.ConnectedId] = newDist
					heap.Push(&queue, &graf.QueueItem{
						Value:    *edge.ConnectedVertex,
						Priority: newDist,
					})

					if cv == sv {
						aMap[edge.ConnectedId] = ContainerTuple{U: sv.Id, V: edge.ConnectedId}
					} else {
						aMap[edge.ConnectedId] = aMap[cv.Id]
					}
				}
			}
		}
	}

	return auxContainer
}

// DijkstraGeometricPrune Traverses the graph using Dijkstra's Algorithm by restrictions of Geometric Containers
func DijkstraGeometricPrune(graph graf.Graph, s string, e string, aux map[ContainerTuple][]string) (*graf.AlgorithmReport, error) {
	var sv graf.Vertex
	var ev graf.Vertex
	var ok bool

	if sv, ok = graph.Vertices[s]; !ok {
		return nil, errors.New(fmt.Sprintf("Starting Vertex: %s is not in Graph.", s))
	}
	if ev, ok = graph.Vertices[e]; !ok {
		return nil, errors.New(fmt.Sprintf("Ending Vertex: %s is not in Graph.", e))
	}

	report := graf.AlgorithmReport{
		StartVertex:    &sv,
		EndVertex:      &ev,
		Distance:       math.Inf(1),
		DistanceMap:    map[string]float64{s: 0},
		PredecessorMap: map[string]*graf.Vertex{s: nil},
		VisitMap:       map[string]bool{s: false},
	}

	queue := graf.BlankQueue()
	heap.Push(&queue, &graf.QueueItem{
		Value:    sv,
		Priority: 0,
	})

	for !queue.IsEmpty() {
		cq := heap.Pop(&queue).(*graf.QueueItem)
		cv := cq.Value.(graf.Vertex)

		if cv == ev {
			report.Distance = report.DistanceMap[ev.Id]

			pv := &ev
			for pv != nil {
				report.PredecessorChain = append(report.PredecessorChain, *pv)
				tv := report.PredecessorMap[pv.Id]
				pv = tv
			}

			return &report, nil
		}

		if report.VisitMap[cv.Id] {
			continue
		}

		report.VisitMap[cv.Id] = true
		for _, edge := range graph.Edges[cv] {

			container, ok := aux[ContainerTuple{U: cv.Id, V: edge.ConnectedId}]
			if !ok || !ContainsVertex(container, ev.Id) {
				continue
			}

			newDist := cq.Priority + edge.Weight
			if dist, ok := report.DistanceMap[edge.ConnectedId]; !ok || newDist < dist {
				report.DistanceMap[edge.ConnectedId] = newDist
				report.PredecessorMap[edge.ConnectedId] = &cv
				heap.Push(&queue, &graf.QueueItem{
					Value:    *edge.ConnectedVertex,
					Priority: newDist,
				})
			}
		}
	}

	return &report, nil
}
