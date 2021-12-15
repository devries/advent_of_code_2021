package main

import (
	"container/heap"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day15.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	grid := parseGrid(lines)
	total := solveGrid(grid)

	return total
}

type Grid struct {
	Points map[utils.Point]int
	X      int
	Y      int
}

func parseGrid(lines []string) Grid {
	xsize := len([]rune(lines[0]))
	ysize := len(lines)
	r := Grid{make(map[utils.Point]int), xsize * 5, ysize * 5}

	for j := 0; j < r.Y; j++ {
		for i := 0; i < r.X; i++ {
			ln := lines[j%ysize]
			c := []rune(ln)[i%xsize]
			scoreIncrement := i/xsize + j/ysize
			r.Points[utils.Point{X: i, Y: j}] = (int(c-'0')-1+scoreIncrement)%9 + 1
		}
	}

	return r
}

// Define a priority Queue and associated container.Heap interfaces
type QueueItem struct {
	Point utils.Point
	Score int
	Index int
}

type PriorityQueue []*QueueItem

func (q PriorityQueue) Len() int { return len(q) }

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].Score < q[j].Score
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].Index = i
	q[j].Index = j
}

func (q *PriorityQueue) Push(x interface{}) {
	n := len(*q)
	point := x.(*QueueItem)
	point.Index = n
	*q = append(*q, point)
}

func (q *PriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*q = old[0 : n-1]
	return item
}

func (q *PriorityQueue) update(item *QueueItem, score int) {
	item.Score = score
	heap.Fix(q, item.Index)
}

func solveGrid(grid Grid) int {
	queue := make(PriorityQueue, grid.X*grid.Y)
	allpoints := make(map[utils.Point]*QueueItem)

	idx := 0
	for j := 0; j < grid.Y; j++ {
		for i := 0; i < grid.X; i++ {
			pt := utils.Point{X: i, Y: j}
			var qi = QueueItem{pt, math.MaxInt, idx}
			if i == 0 && j == 0 {
				qi.Score = 0
			}
			queue[idx] = &qi
			allpoints[pt] = &qi
			idx++
		}
	}

	heap.Init(&queue)

	for queue.Len() > 0 {
		item := heap.Pop(&queue).(*QueueItem)

		for _, d := range utils.Directions {
			next := item.Point.Add(d)
			value, ok := grid.Points[item.Point.Add(d)]
			if ok {
				qi := allpoints[next]
				newscore := item.Score + value
				if newscore < qi.Score {
					queue.update(qi, newscore)
				}
			}
		}
	}

	return (*allpoints[utils.Point{X: grid.X - 1, Y: grid.Y - 1}]).Score
}
