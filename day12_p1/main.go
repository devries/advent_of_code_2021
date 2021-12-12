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
	seen := make(map[string]bool)

	go func() {
		mazeSolveRecursor(maze, &path, seen, ch)
		close(ch)
	}()

	return ch
}

func mazeSolveRecursor(maze map[string][]string, path *[]string, seen map[string]bool, ch chan<- []string) {
	current := (*path)[len(*path)-1]

	if current == strings.ToLower(current) {
		seen[current] = true
	}

	if current == "end" {
		// path complete! Send a copy back on the channel
		complete := make([]string, len(*path))
		copy(complete, *path)
		ch <- complete
		seen[current] = false
		return
	}

	nextpos := len(*path)
	*path = append(*path, "") // place holder for next step

	nextsteps := maze[current]

	for _, room := range nextsteps {
		if !seen[room] {
			(*path)[nextpos] = room
			mazeSolveRecursor(maze, path, seen, ch)
		}
	}

	*path = (*path)[:nextpos]
	seen[current] = false

	return
}
