package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day19.txt")
	utils.Check(err, "error opening input file")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	beaconSets := parseInput(lines)

	unprocessedIndex := 1
	scanners := []Point{{0, 0, 0}}

	for origin := 0; origin < len(beaconSets)-1; origin++ {
		ma := findMetrics(beaconSets[origin])
		for i := unprocessedIndex; i < len(beaconSets); i++ {
			mb := findMetrics(beaconSets[i])
			matches := findMatches(ma, mb)
			if countOverlapPoints(matches) >= 12 {
				transformMap := make(map[Transform]int)
				maxTransforms := 0
				for _, m := range matches {
					pa := m.MetricA.A
					pb := reorient[m.Orientation](m.MetricB.A)
					offset := Point{pb.X - pa.X, pb.Y - pa.Y, pb.Z - pa.Z}
					transforms := transformMap[Transform{m.Orientation, offset}] + 1
					transformMap[Transform{m.Orientation, offset}] = transforms
					if transforms > maxTransforms {
						maxTransforms = transforms
					}
				}

				if maxTransforms > 0 {
					for k, v := range transformMap {
						if v == maxTransforms {
							// Transform found
							scanners = append(scanners, k.Offset)
							for j := 0; j < len(beaconSets[i]); j++ {
								beaconSets[i][j] = k.Apply(beaconSets[i][j])
							}
							beaconSets[unprocessedIndex], beaconSets[i] = beaconSets[i], beaconSets[unprocessedIndex]
							unprocessedIndex++
						}
					}
				}
			}
		}
	}

	maxDistance := 0
	for i := 0; i < len(scanners)-1; i++ {
		for j := i + 1; j < len(scanners); j++ {
			a := scanners[i]
			b := scanners[j]
			d := abs(a.X-b.X) + abs(a.Y-b.Y) + abs(a.Z-b.Z)
			if d > maxDistance {
				maxDistance = d
			}
		}
	}

	return maxDistance
}

type Point struct {
	X int
	Y int
	Z int
}

func parseInput(lines []string) [][]Point {
	beaconSets := make([][]Point, 0)

	beacons := make([]Point, 0)
	for i, line := range lines {
		if strings.HasPrefix(line, "---") && i > 0 {
			beaconSets = append(beaconSets, beacons)
			beacons = make([]Point, 0)
			continue
		}
		var p Point
		n, _ := fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		if n == 3 {
			beacons = append(beacons, p)
		}
	}
	beaconSets = append(beaconSets, beacons)

	return beaconSets
}

var reorient = [](func(p Point) Point){
	func(p Point) Point { return Point{p.X, p.Y, p.Z} },
	func(p Point) Point { return Point{p.Y, -p.X, p.Z} },
	func(p Point) Point { return Point{-p.X, -p.Y, p.Z} },
	func(p Point) Point { return Point{-p.Y, p.X, p.Z} },
	func(p Point) Point { return Point{-p.X, p.Y, -p.Z} },
	func(p Point) Point { return Point{-p.Y, -p.X, -p.Z} },
	func(p Point) Point { return Point{p.X, -p.Y, -p.Z} },
	func(p Point) Point { return Point{p.Y, p.X, -p.Z} },
	func(p Point) Point { return Point{p.X, -p.Z, p.Y} },
	func(p Point) Point { return Point{-p.Z, -p.X, p.Y} },
	func(p Point) Point { return Point{-p.X, p.Z, p.Y} },
	func(p Point) Point { return Point{p.Z, p.X, p.Y} },
	func(p Point) Point { return Point{-p.X, -p.Z, -p.Y} },
	func(p Point) Point { return Point{-p.Z, p.X, -p.Y} },
	func(p Point) Point { return Point{p.X, p.Z, -p.Y} },
	func(p Point) Point { return Point{p.Z, -p.X, -p.Y} },
	func(p Point) Point { return Point{p.Y, p.Z, p.X} },
	func(p Point) Point { return Point{-p.Z, p.Y, p.X} },
	func(p Point) Point { return Point{-p.Y, -p.Z, p.X} },
	func(p Point) Point { return Point{p.Z, -p.Y, p.X} },
	func(p Point) Point { return Point{p.Y, -p.Z, -p.X} },
	func(p Point) Point { return Point{p.Z, p.Y, -p.X} },
	func(p Point) Point { return Point{-p.Y, p.Z, -p.X} },
	func(p Point) Point { return Point{-p.Z, -p.Y, -p.X} },
}

type ThreeDistance [3]int

func (t *ThreeDistance) Len() int           { return 3 }
func (t *ThreeDistance) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t *ThreeDistance) Less(i, j int) bool { return t[i] < t[j] }

type Metric struct {
	Distances ThreeDistance
	Distance  int
	DeltaX    int
	DeltaY    int
	DeltaZ    int
	A         Point
	B         Point
}

type Metrics []Metric

func (m Metrics) Len() int           { return len(m) }
func (m Metrics) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Metrics) Less(i, j int) bool { return m[i].Distance < m[j].Distance }

func (d ThreeDistance) String() string {
	return fmt.Sprintf("%d,%d,%d", d[0], d[1], d[2])
}

func findMetrics(beacons []Point) Metrics {
	res := make(Metrics, 0)

	for i := 0; i < len(beacons)-1; i++ {
		for j := i + 1; j < len(beacons); j++ {
			a := beacons[i]
			b := beacons[j]
			dx := b.X - a.X
			dy := b.Y - a.Y
			dz := b.Z - a.Z
			dxm := abs(dx)
			dym := abs(dy)
			dzm := abs(dz)
			m := Metric{ThreeDistance{dxm, dym, dzm}, dxm + dym + dzm, dx, dy, dz, a, b}
			sort.Sort(&m.Distances)

			res = append(res, m)
		}
	}

	sort.Sort(res)
	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Match struct {
	MetricA     Metric
	MetricB     Metric
	Orientation int
}

func findMatches(a Metrics, b Metrics) []Match {
	res := make([]Match, 0)

	for i, j := 0, 0; i < len(a) || j < len(b); {
		if i >= len(a) || j >= len(b) {
			break
		}
		ma := a[i]
		mb := b[j]

		if ma.Distance < mb.Distance {
			i++
			continue
		}
		if mb.Distance > ma.Distance {
			j++
			continue
		}

		// Same distance
		if ma.Distances == mb.Distances {
			// Same components
			pa := Point{ma.DeltaX, ma.DeltaY, ma.DeltaZ}
			pb := Point{mb.DeltaX, mb.DeltaY, mb.DeltaZ}
			for k, f := range reorient {
				if pa == f(pb) {
					res = append(res, Match{ma, mb, k})
				}
			}
		}
		if i < len(a)-1 && j < len(b)-1 {
			if a[i+1].Distance < b[j+1].Distance {
				i++
			} else {
				j++
			}
		} else if j < len(b)-1 {
			j++
		} else if i < len(a)-1 {
			i++
		} else {
			break
		}
	}

	return res
}

func countOverlapPoints(matches []Match) int {
	res := make(map[Point]bool)

	for _, m := range matches {
		res[m.MetricA.A] = true
		res[m.MetricA.B] = true
	}

	return len(res)
}

type Transform struct {
	Orientation int
	Offset      Point
}

func (t Transform) Apply(p Point) Point {
	prot := reorient[t.Orientation](p)
	pprime := Point{prot.X - t.Offset.X, prot.Y - t.Offset.Y, prot.Z - t.Offset.Z}

	return pprime
}
