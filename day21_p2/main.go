package main

import (
	"fmt"
	"os"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day21.txt")
	utils.Check(err, "error opening input file")
	defer f.Close()

	lines := utils.ReadLines(f)

	var startA, startB int
	_, err = fmt.Sscanf(lines[0], "Player 1 starting position: %d", &startA)
	utils.Check(err, "unable to parse input")
	_, err = fmt.Sscanf(lines[1], "Player 2 starting position: %d", &startB)
	utils.Check(err, "unable to parse input")

	r := solve(startA, startB)
	fmt.Println(r)
}

func solve(startA int, startB int) int64 {
	statesA := make(map[State]int64)
	statesB := make(map[State]int64)
	var keepgoingA, keepgoingB bool

	statesA[State{0, startA, 0}] = 1
	statesB[State{0, startB, 0}] = 1
	for turn := 1; turn < 20; turn++ {
		keepgoingA = step(statesA, turn)
		keepgoingB = step(statesB, turn)

		if keepgoingA == false && keepgoingB == false {
			break
		}
	}

	var winsA, winsB int64
	for kA, vA := range statesA {
		if kA.Score >= 21 {
			for kB, vB := range statesB {
				if kB.Score < 21 && kB.Turns == kA.Turns-1 {
					// A wins
					winsA += vA * vB
				}
			}
		} else {
			for kB, vB := range statesB {
				if kB.Score >= 21 && kB.Turns == kA.Turns {
					// B wins
					winsB += vA * vB
				}
			}
		}
	}

	if winsA > winsB {
		return winsA
	} else {
		return winsB
	}
}

var universeMultiplier = map[int]int64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

type State struct {
	Score    int
	Position int
	Turns    int
}

func step(s map[State]int64, turn int) bool {
	keepgoing := false
	keys := make([]State, 0)

	// Find the states that will evolve this turn
	for k := range s {
		if k.Score < 21 && k.Turns == turn-1 {
			keys = append(keys, k)
		}
	}

	for _, k := range keys {
		for i := 3; i <= 9; i++ {
			pos := (k.Position + i) % 10
			var score int
			if pos == 0 {
				score = k.Score + 10
			} else {
				score = k.Score + pos
			}
			if score < 21 {
				keepgoing = true
			}

			s[State{score, pos, turn}] += s[k] * universeMultiplier[i]
		}
	}

	return keepgoing
}
