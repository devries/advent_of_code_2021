package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day24.txt")
	utils.Check(err, "error opening input file")
	defer f.Close()

	solve(f)
}

func solve(r io.Reader) {
	lines := utils.ReadLines(r)

	alu := ALU{make(map[string]Stack), make([]Stack, 0)}
	alu.registers["w"] = Stack{"0"}
	alu.registers["x"] = Stack{"0"}
	alu.registers["y"] = Stack{"0"}
	alu.registers["z"] = Stack{"0"}
	alu = Execute(alu, lines)
	for k, v := range alu.registers {
		fmt.Printf("%s: %v\nDepends: %v\n\n", k, v, findDeps(v))
	}
	for i, s := range alu.conditions {
		fmt.Printf("cond:%d: %v\nDepends: %v\n\n", i, s, findDeps(s))
	}

	for i := int64(1); i < 10; i++ {
		sol, ok := Evaluate(alu, []int64{i}, "cond:1")
		if ok {
			fmt.Printf("inp:0 = %d then cond:1 = %d\n", i, sol)
		} else {
			fmt.Printf("inp:0 = %d then cond:1 is undefined\n", i)
		}
	}
}

func findDeps(ops []string) []string {
	deps := make(map[string]bool)

	for _, op := range ops {
		if strings.HasPrefix(op, "inp:") {
			deps[op] = true
		}
	}

	ret := make([]string, 0)
	for i := 0; i < 14; i++ {
		val := fmt.Sprintf("inp:%d", i)
		if deps[val] {
			ret = append(ret, val)
		}
	}
	return ret
}

// Evaluate expression item (can be register or condition)
func Evaluate(a ALU, input []int64, item string) (int64, bool) {
	var instructions []string

	if strings.HasPrefix(item, "cond:") {
		// Evaluating condition
		var cond int
		_, err := fmt.Sscanf(item, "cond:%d", &cond)
		utils.Check(err, "error parsing condition number")

		instructions = a.conditions[cond]
	} else {
		instructions = a.registers[item]
	}

	s := make(IntStack, 0)
	for _, v := range instructions {
		switch {
		case v == "add":
			pb := s.Pop()
			pa := s.Pop()
			r := pa + pb
			s.Push(r)

		case v == "mul":
			pb := s.Pop()
			pa := s.Pop()
			r := pa * pb
			s.Push(r)

		case v == "div":
			pb := s.Pop()
			if pb == 0 {
				return 0, false
			}
			pa := s.Pop()
			r := pa / pb
			s.Push(r)

		case v == "mod":
			pb := s.Pop()
			pa := s.Pop()
			if pa < 0 || pb <= 0 {
				return 0, false
			}
			r := pa % pb
			s.Push(r)

		case v == "eql":
			pb := s.Pop()
			pa := s.Pop()
			if pa == pb {
				s.Push(1)
			} else {
				s.Push(0)
			}

		case strings.HasPrefix(v, "inp:"):
			var num int
			_, err := fmt.Sscanf(v, "inp:%d", &num)
			utils.Check(err, "error parsing input number")

			s.Push(input[num])

		case strings.HasPrefix(v, "cond:"):
			var num int
			_, err := fmt.Sscanf(item, "cond:%d", &num)
			utils.Check(err, "error parsing condition number")
			r, ok := Evaluate(a, input, v)
			if ok {
				s.Push(r)
			} else {
				return 0, false
			}

		default:
			r, err := strconv.ParseInt(v, 10, 64)
			utils.Check(err, "error parsing number")

			s.Push(r)
		}
	}

	return s.Pop(), true
}

type ALU struct {
	registers  map[string]Stack
	conditions []Stack
}

func (a ALU) getValue(s string) Stack {
	switch s {
	case "w", "x", "y", "z":
		return a.registers[s]
	default:
		return Stack{s}
	}
}

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() string {
	old := *s
	v := old[len(old)-1]
	*s = old[:len(old)-1]

	return v
}

type IntStack []int64

func (s *IntStack) Push(v int64) {
	*s = append(*s, v)
}

func (s *IntStack) Pop() int64 {
	old := *s
	v := old[len(old)-1]
	*s = old[:len(old)-1]

	return v
}

func Execute(a ALU, lines []string) ALU {
	input_counter := 0
	condition_counter := 0

	for _, ln := range lines {
		parts := strings.Fields(ln)

		switch parts[0] {
		case "inp":
			a.registers[parts[1]] = Stack{fmt.Sprintf("inp:%d", input_counter)}
			input_counter++

		case "add":
			if parts[2] == "0" {
				break
			}
			reg := a.registers[parts[1]]
			b := a.getValue(parts[2])
			for _, v := range b {
				reg.Push(v)
			}
			reg.Push("add")
			a.registers[parts[1]] = reg

		case "mul":
			if parts[2] == "0" {
				a.registers[parts[1]] = Stack{"0"}
			} else if parts[1] == "1" {
				break
			} else {
				reg := a.registers[parts[1]]
				b := a.getValue(parts[2])
				for _, v := range b {
					reg.Push(v)
				}
				reg.Push("mul")
				a.registers[parts[1]] = reg
			}

		case "div":
			if parts[2] == "1" {
				break
			}
			reg := a.registers[parts[1]]
			b := a.getValue(parts[2])
			for _, v := range b {
				reg.Push(v)
			}
			reg.Push("div")
			a.registers[parts[1]] = reg

		case "mod":
			if parts[2] == "1" {
				break
			}
			reg := a.registers[parts[1]]
			b := a.getValue(parts[2])
			for _, v := range b {
				reg.Push(v)
			}
			reg.Push("mod")
			a.registers[parts[1]] = reg

		case "eql":
			c := Stack{}
			for _, v := range a.getValue(parts[1]) {
				c.Push(v)
			}
			for _, v := range a.getValue(parts[2]) {
				c.Push(v)
			}
			c.Push("eql")

			a.registers[parts[1]] = Stack{fmt.Sprintf("cond:%d", condition_counter)}
			condition_counter++
			a.conditions = append(a.conditions, c)

		}
	}
	return a
}
