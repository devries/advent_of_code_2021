package main

import (
	"container/heap"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day23.txt")
	utils.Check(err, "error opening input file")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	finalState := State{[4][2]rune{{'A', 'A'}, {'B', 'B'}, {'C', 'C'}, {'D', 'D'}}, [7]rune{}}
	s := parseInput(lines)

	seen := make(map[State]int)
	seen[s] = 0

	// Add initial StateEnergy to queue and iterate through
	queue := make(PriorityQueue, 1)
	queue[0] = StateEnergy{s, 0}

	heap.Init(&queue)

	for queue.Len() > 0 {
		current := heap.Pop(&queue).(StateEnergy)

		for _, se := range nextStates(current.Value, current.Energy) {
			s := se.Value
			e := se.Energy

			foundEnergy, ok := seen[s]
			if !ok || e < foundEnergy {
				seen[s] = e
				heap.Push(&queue, se)
			}
		}
	}

	return seen[finalState]
}

// Horizontal Order: 01A2B3C4D56
// Row 0:              0 1 2 3
// Row 1:              0 1 2 3
// End State:          A B C D
type State struct {
	Pods    [4][2]rune
	Hallway [7]rune
}

func (s State) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "\n#############\n")
	fmt.Fprintf(&sb, "#")
	for i := 0; i < 7; i++ {
		c := s.Hallway[i]
		if c == 0 {
			fmt.Fprintf(&sb, ".")
		} else {
			fmt.Fprintf(&sb, "%c", c)
		}
		if i > 0 && i < 5 {
			fmt.Fprintf(&sb, ".")
		}
	}
	fmt.Fprintf(&sb, "#\n###")

	for j := 0; j < 2; j++ {
		for i := 0; i < 4; i++ {
			c := s.Pods[i][j]
			if c == 0 {
				fmt.Fprintf(&sb, ".#")
			} else {
				fmt.Fprintf(&sb, "%c#", c)
			}
		}
		if j == 0 {
			fmt.Fprintf(&sb, "##\n  #")
		} else {
			fmt.Fprintf(&sb, "\n")
		}
	}
	fmt.Fprintf(&sb, "  #########\n")

	return sb.String()
}

type StateEnergy struct {
	Value  State
	Energy int
}

func parseInput(lines []string) State {
	s := State{}

	for r := 0; r < 2; r++ {
		values := []rune(lines[r+2])
		for p := 0; p < 4; p++ {
			c := values[2*p+3]
			s.Pods[p][r] = c
		}
	}

	return s
}

var energyPerStep = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

// 01X2X3X4X56#
//   0 1 2 3
var stepsFromHallwayToPodDoor = [][]int{
	{2, 4, 6, 8},
	{1, 3, 5, 7},
	{1, 1, 3, 5},
	{3, 1, 1, 3},
	{5, 3, 1, 1},
	{7, 5, 3, 1},
	{8, 6, 4, 2},
}

func nextStates(s State, startEnergy int) []StateEnergy {
	ret := make([]StateEnergy, 0)
	var energy int

	// Check everything in hallway to see if it can go home
outerHallway:
	for pos := 0; pos < 7; pos++ {
		c := s.Hallway[pos]
		if c == 0 {
			continue
		}

		// Letter found, check if the corresponding pod is ready
		podDesired := int(c - 'A')
		for row := 0; row < 2; row++ {
			if s.Pods[podDesired][row] != c && s.Pods[podDesired][row] != 0 {
				// Pod is not ready
				continue outerHallway
			}
		}

		// Check if the path to pod is open
		stopping := podDesired + 2
		if pos < stopping {
			for i := pos + 1; i < stopping; i++ {
				if s.Hallway[i] != 0 {
					continue outerHallway
				}
			}
		} else {
			for i := pos - 1; i >= stopping; i-- {
				if s.Hallway[i] != 0 {
					continue outerHallway
				}
			}
		}

		// It is possible to move piece home
		var endRow int
		if s.Pods[podDesired][1] == 0 {
			endRow = 1
		} else {
			endRow = 0
		}

		energy = startEnergy + (stepsFromHallwayToPodDoor[pos][podDesired]+endRow+1)*energyPerStep[c]
		nState := s
		nState.Hallway[pos] = 0
		nState.Pods[podDesired][endRow] = c
		ret = append(ret, StateEnergy{nState, energy})
	}

	if len(ret) > 0 {
		return ret
	}

	// Run through all pods to see if they can evolve
outerPods:
	for pod := 0; pod < 4; pod++ {
		endState := 'A' + rune(pod)
		for row := 0; row < 2; row++ {
			c := s.Pods[pod][row]
			if c == endState {
				skip := true
				// Check if all descending runes are correct
				for d := row + 1; d < 2; d++ {
					if s.Pods[pod][d] != endState {
						skip = false
					}
				}
				if skip {
					continue outerPods
				}
			}

			if c != 0 {
				// Will have to try moving available piece then end the pod

				// 2, 3, 4, 5
				stopping := pod + 2 // This is the split point in the hallway for this pod

				// Iterate from hallway position down and up until you hit another letter
				for i := stopping - 1; i >= 0; i-- {
					if s.Hallway[i] == 0 {
						nState := s
						nState.Pods[pod][row] = 0
						nState.Hallway[i] = c
						energy = startEnergy + (stepsFromHallwayToPodDoor[i][pod]+row+1)*energyPerStep[c]
						ret = append(ret, StateEnergy{nState, energy})
					} else {
						break
					}
				}

				for i := stopping; i < 7; i++ {
					if s.Hallway[i] == 0 {
						nState := s
						nState.Pods[pod][row] = 0
						nState.Hallway[i] = c
						energy = startEnergy + (stepsFromHallwayToPodDoor[i][pod]+row+1)*energyPerStep[c]
						ret = append(ret, StateEnergy{nState, energy})
					} else {
						break
					}
				}
				// Can't move the letter under this one
				continue outerPods
			}
		}
	}

	return ret
}

type PriorityQueue []StateEnergy

func (q PriorityQueue) Len() int { return len(q) }

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].Energy < q[j].Energy
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Push(x interface{}) {
	se := x.(StateEnergy)
	*q = append(*q, se)
}

func (q *PriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}
