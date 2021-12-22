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

func solve(startA int, startB int) int {
	// The dice rolls are 9N+6 from N=0... This means a %10 sequence of moves for
	// each player that repeats.

	movesA := []int{6, 4, 2, 0, 8}
	movesB := []int{5, 3, 1, 9, 7}

	pointsA := getPointSequence(startA, movesA)
	pointsB := getPointSequence(startB, movesB)

	turnsA := movesToWin(pointsA)
	turnsB := movesToWin(pointsB)

	var turns, scoreA, scoreB, rolls, lowScore int
	if turnsA < turnsB {
		turns = turnsA
		scoreA = score(turns, pointsA)
		scoreB = score(turns-1, pointsB)
		rolls = turns*6 - 3
		lowScore = scoreB
	} else {
		turns = turnsB
		scoreA = score(turns, pointsA)
		scoreB = score(turns, pointsB)
		rolls = turns * 6
		lowScore = scoreA
	}

	return rolls * lowScore
}

func getPointSequence(start int, moves []int) []int {
	// A point sequence derives from the move sequence, but may not be the same length

	moveSum := 0
	for _, v := range moves {
		moveSum += v
	}

	cycleOffset := moveSum % 10

	var repeats int
	if cycleOffset == 0 {
		repeats = 1
	} else {
		lcm := int(utils.Lcm(int64(cycleOffset), 10))
		repeats = lcm / cycleOffset
	}

	ret := make([]int, 0)
	pos := start
	for i := 0; i < repeats*len(moves); i++ {
		pos += moves[i%len(moves)]
		pos %= 10
		if pos == 0 {
			ret = append(ret, 10)
		} else {
			ret = append(ret, pos)
		}
	}

	return ret
}

func movesToWin(pointSequence []int) int {
	total := 0
	for _, i := range pointSequence {
		total += i
	}

	multiples := 1000 / total

	points := multiples * total
	moves := multiples * len(pointSequence)

	for i := 0; i < len(pointSequence); i++ {
		if points >= 1000 {
			return moves
		}
		moves += 1
		points += pointSequence[i]
	}
	// Should never get here
	return 0
}

func score(moves int, pointSequence []int) int {
	total := 0
	for _, i := range pointSequence {
		total += i
	}

	seqLen := len(pointSequence)
	repeats := moves / seqLen
	score := repeats * total

	for m := repeats * seqLen; m < moves; m++ {
		score += pointSequence[m%seqLen]
	}

	return score
}
