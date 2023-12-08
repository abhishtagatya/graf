package imp

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"graf"
	"math"
	"os"
	"strconv"
	"strings"
)

type VertexCoordinate struct {
	Latitude  float64
	Longitude float64
}

func LoadCoordinate(fileName string, prefix string) (map[string]VertexCoordinate, error) {
	coordMap := make(map[string]VertexCoordinate)

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

			var longFix string
			var latFix string

			var longitude float64
			var latitude float64

			if strings.HasPrefix(lineVertex[2], "-") {
				longFix = lineVertex[2][:3] + "." + lineVertex[2][3:]
			} else {
				longFix = lineVertex[2][:2] + "." + lineVertex[2][2:]
			}

			if longitude, err = strconv.ParseFloat(longFix, 32); err != nil {
				return nil, err
			}

			if strings.HasPrefix(lineVertex[3], "-") {
				latFix = lineVertex[3][:3] + "." + lineVertex[3][3:]
			} else {
				latFix = lineVertex[3][:2] + "." + lineVertex[3][2:]
			}

			if latitude, err = strconv.ParseFloat(latFix, 32); err != nil {
				return nil, err
			}

			coordMap[lineVertex[1]] = VertexCoordinate{
				Latitude:  latitude,
				Longitude: longitude,
			}
		}
	}

	return coordMap, nil
}

func HaversineDistance(s, v VertexCoordinate) float64 {
	const earthRadius = 6371 // Earth's radius in kilometers

	// Convert degrees to radians
	lat1Rad := s.Latitude * math.Pi / 180
	lon1Rad := s.Longitude * math.Pi / 180
	lat2Rad := v.Latitude * math.Pi / 180
	lon2Rad := v.Longitude * math.Pi / 180

	// Calculate differences
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Haversine formula
	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate distance
	distance := earthRadius * c

	return distance * 10000
}

func AStarHaversine(graph graf.Graph, s string, e string, heuristic map[string]VertexCoordinate) (*HeuristicAlgorithmReport, error) {
	var sv graf.Vertex
	var ev graf.Vertex
	var ok bool

	if sv, ok = graph.Vertices[s]; !ok {
		return nil, errors.New(fmt.Sprintf("Starting Vertex: %s is not in Graph.", s))
	}
	if ev, ok = graph.Vertices[e]; !ok {
		return nil, errors.New(fmt.Sprintf("Ending Vertex: %s is not in Graph.", e))
	}

	maxTrueDistance := HaversineDistance(heuristic[sv.Id], heuristic[ev.Id])
	fmt.Println(maxTrueDistance)

	report := HeuristicAlgorithmReport{
		StartVertex:    &sv,
		EndVertex:      &ev,
		Distance:       math.Inf(1),
		DistanceMap:    map[string]float64{s: 0},
		HeuristicMap:   map[string]float64{s: maxTrueDistance},
		PredecessorMap: map[string]*graf.Vertex{s: nil},
		VisitMap:       map[string]bool{s: false},
	}

	queue := BlankHeuristicQueue()
	heap.Push(&queue, &HeuristicQueueItem{
		Value:    sv,
		Weight:   0,
		Priority: 0,
	})

	for !queue.IsEmpty() {
		cq := heap.Pop(&queue).(*HeuristicQueueItem)
		cv := cq.Value.(graf.Vertex)

		if cv == ev {
			report.Distance = report.DistanceMap[ev.Id]

			pv := &ev
			for pv != nil {
				report.PredecessorChain = append(report.PredecessorChain, *pv)
				tv := report.PredecessorMap[pv.Id]
				pv = tv
			}

			fmt.Println(report.DistanceMap[ev.Id], report.HeuristicMap[ev.Id])

			return &report, nil
		}

		if report.VisitMap[cv.Id] {
			continue
		}

		report.VisitMap[cv.Id] = true
		for _, edge := range graph.Edges[cv] {
			newDist := cq.Weight + edge.Weight
			newHeur := newDist + HaversineDistance(heuristic[edge.ConnectedId], heuristic[ev.Id])

			if heur, ok := report.HeuristicMap[edge.ConnectedId]; !ok || newHeur < heur {
				report.DistanceMap[edge.ConnectedId] = newDist
				report.HeuristicMap[edge.ConnectedId] = newHeur
				report.PredecessorMap[edge.ConnectedId] = &cv
				heap.Push(&queue, &HeuristicQueueItem{
					Value:    *edge.ConnectedVertex,
					Weight:   newDist,
					Priority: newHeur,
				})
			}
		}
	}

	return &report, nil
}
