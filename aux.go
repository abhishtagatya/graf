package graf

import (
	"container/heap"
	"slices"
)

/* Fully Auxiliary Examples to Specific Problem Implementations. */

type AuxTuple struct {
	U string
	V string
}

func ComputeContainers(graph *Graph) map[AuxTuple][]string {
	aMap := make(map[string]AuxTuple)
	auxContainer := make(map[AuxTuple][]string)

	for sid := range graph.Vertices {
		sv := graph.Vertices[sid]

		distanceMap := map[string]float64{sid: 0}

		queue := BlankQueue()
		heap.Push(&queue, &QueueItem{
			Value:    sv,
			Priority: 0,
		})

		for !queue.IsEmpty() {
			cq := heap.Pop(&queue).(*QueueItem)
			cv := cq.Value.(Vertex)

			if cv != sv {
				if !slices.Contains(auxContainer[aMap[cv.Id]], cv.Id) {
					auxContainer[aMap[cv.Id]] = append(auxContainer[aMap[cv.Id]], cv.Id)
				}
			}

			for _, edge := range graph.Edges[cv] {
				newDist := cq.Priority + edge.Weight
				if dist, ok := distanceMap[edge.ConnectedId]; !ok || newDist < dist {
					distanceMap[edge.ConnectedId] = newDist
					heap.Push(&queue, &QueueItem{
						Value:    *edge.ConnectedVertex,
						Priority: newDist,
					})

					if cv == sv {
						aMap[edge.ConnectedId] = AuxTuple{U: sv.Id, V: edge.ConnectedId}
					} else {
						aMap[edge.ConnectedId] = aMap[cv.Id]
					}
				}
			}

		}
	}

	return auxContainer
}

//type VertexCoordinateAux struct {
//	Latitude  float64 `json:"latitude"`
//	Longitude float64 `json:"longitude"`
//}
//
//func MapVertexCoordinateAux(fileName string, prefix string) (map[string]VertexCoordinateAux, error) {
//
//	mapVCA := make(map[string]VertexCoordinateAux)
//
//	readFile, err := os.Open(fileName)
//	if err != nil {
//		return mapVCA, err
//	}
//
//	fileScanner := bufio.NewScanner(readFile)
//	fileScanner.Split(bufio.ScanLines)
//
//	for fileScanner.Scan() {
//		lineScan := fileScanner.Text()
//		if strings.HasPrefix(lineScan, prefix) {
//			lineVertex := strings.Split(lineScan, " ")
//
//			var longF float64
//			var latF float64
//
//			var longFix string
//			var latFix string
//
//			if strings.HasPrefix(lineVertex[2], "-") {
//				longFix = lineVertex[2][:3] + "." + lineVertex[2][3:]
//			} else {
//				longFix = lineVertex[2][:2] + "." + lineVertex[2][2:]
//			}
//
//			if strings.HasPrefix(lineVertex[3], "-") {
//				latFix = lineVertex[3][:3] + "." + lineVertex[3][3:]
//			} else {
//				latFix = lineVertex[3][:2] + "." + lineVertex[3][2:]
//			}
//
//			if longF, err = strconv.ParseFloat(longFix, 32); err != nil {
//				return mapVCA, err
//			}
//
//			if latF, err = strconv.ParseFloat(latFix, 32); err != nil {
//				return mapVCA, err
//			}
//
//			mapVCA[lineVertex[1]] = VertexCoordinateAux{
//				Latitude:  latF,
//				Longitude: longF,
//			}
//		}
//	}
//
//	return mapVCA, nil
//}
//
//func toRadians(degrees float64) float64 {
//	return degrees * (math.Pi / 180)
//}
//
//func toDegrees(radians float64) float64 {
//	return radians * (180 / math.Pi)
//}
//
//func calculateBearing(s, v VertexCoordinateAux) float64 {
//	deltaLong := v.Longitude - s.Longitude
//	x := math.Cos(toRadians(v.Latitude)) * math.Sin(toRadians(deltaLong))
//	y := math.Cos(toRadians(s.Latitude))*math.Sin(toRadians(v.Latitude)) - math.Sin(toRadians(s.Latitude))*math.Cos(toRadians(v.Latitude))*math.Cos(toRadians(deltaLong))
//
//	initBearing := toDegrees(math.Atan2(x, y))
//	bearing := math.Mod(initBearing+360, 360)
//	return bearing
//}
//
//func getDirection(bearing float64) string {
//	directions := [8]string{"North", "Northeast", "East", "Southeast", "South", "Southwest", "West", "Northwest"}
//	index := int(math.Round(bearing/45)) % 8
//
//	return directions[index]
//}
//
//func GetDirection(s, v VertexCoordinateAux) string {
//	return getDirection(calculateBearing(s, v))
//}
//
//func VCAAllowed(currDirection, globalDirection string) bool {
//	validDirection := map[string][]string{
//		"North": {
//			"North", "Northeast", "East", "Northwest", "West",
//		},
//		"Northeast": {
//			"Northeast", "East", "Southeast", "North", "Northwest",
//		},
//		"East": {
//			"East", "Southeast", "South", "Northeast", "North",
//		},
//		"Southeast": {
//			"Southeast", "South", "Southwest", "East", "Northeast",
//		},
//		"South": {
//			"South", "Southwest", "West", "Southeast", "East",
//		},
//		"Southwest": {
//			"Southwest", "West", "Northwest", "South", "Southeast",
//		},
//		"West": {
//			"West", "Northwest", "North", "Southwest", "South",
//		},
//		"Northwest": {
//			"Northwest", "North", "Northeast", "West", "Southwest",
//		},
//	}
//
//	return slices.Contains(validDirection[globalDirection], currDirection)
//}
