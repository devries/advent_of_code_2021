package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/devries/advent_of_code_2021/utils"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Parse()
	f, err := os.Open("../inputs/day12.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f, utils.Verbose)
	fmt.Println(r)
}

func solve(r io.Reader, verbose bool) int {
	lines := utils.ReadLines(r)

	maze := parseMaze(lines)
	ch := mazeSolver(maze)

	total := 0
	for p := range ch {
		if verbose {
			fmt.Println(strings.Join(p, ","))
		}
		total += 1
	}

	return total
}

func parseMaze(lines []string) map[string][]string {
	maze := make(map[string][]string)

	for _, ln := range lines {
		parts := strings.Split(ln, "-")

		maze[parts[0]] = append(maze[parts[0]], parts[1])
		maze[parts[1]] = append(maze[parts[1]], parts[0])
	}

	return maze
}

// Generate maze solutions and export them on channel

func mazeSolver(maze map[string][]string) <-chan []string {
	ch := make(chan []string)
	path := []string{"start"}
	seen := make(map[string]int)

	go func() {
		mazeSolveRecursor(maze, &path, seen, 2, ch)
		close(ch)
	}()

	return ch
}

func mazeSolveRecursor(maze map[string][]string, path *[]string, seen map[string]int, maxseen int, ch chan<- []string) {
	current := (*path)[len(*path)-1]

	maxallowed := maxseen
	if current == strings.ToLower(current) {
		if current == "start" {
			maxallowed = 1
		}

		if seen[current] >= maxallowed {
			// do not continue path
			return
		}

		if seen[current]++; seen[current] == 2 {
			maxseen = 1
		}
	}

	if current == "end" {
		// path complete! Send a copy back on the channel
		complete := make([]string, len(*path))
		copy(complete, *path)
		ch <- complete
		seen[current]--
		return
	}

	nextpos := len(*path)
	*path = append(*path, "") // place holder for next step

	nextsteps := maze[current]

	for _, room := range nextsteps {
		(*path)[nextpos] = room
		mazeSolveRecursor(maze, path, seen, maxseen, ch)
	}

	*path = (*path)[:nextpos]
	seen[current]--

	return
}
