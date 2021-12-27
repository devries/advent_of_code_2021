package main

import (
	"strings"
	"testing"

	"github.com/devries/advent_of_code_2021/utils"
)

var testProgram = `inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2`

func TestExecEval(t *testing.T) {
	tests := []struct {
		program string
		inputs  []int64
		w       int64
		x       int64
		y       int64
		z       int64
	}{
		{testProgram, []int64{9}, 1, 0, 0, 1},
	}

	for _, test := range tests {
		r := strings.NewReader(test.program)
		lines := utils.ReadLines(r)
		alu := ALU{make(map[string]Stack), make([]Stack, 0)}
		alu.registers["w"] = Stack{"0"}
		alu.registers["x"] = Stack{"0"}
		alu.registers["y"] = Stack{"0"}
		alu.registers["z"] = Stack{"0"}

		alu = Execute(alu, lines)
		w, _ := Evaluate(alu, test.inputs, "w")
		x, _ := Evaluate(alu, test.inputs, "x")
		y, _ := Evaluate(alu, test.inputs, "y")
		z, _ := Evaluate(alu, test.inputs, "z")

		if w != test.w {
			t.Errorf("Expected w=%d, got %d", test.w, w)
		}
		if x != test.x {
			t.Errorf("Expected x=%d, got %d", test.x, x)
		}
		if y != test.y {
			t.Errorf("Expected y=%d, got %d", test.y, y)
		}
		if z != test.z {
			t.Errorf("Expected z=%d, got %d", test.z, z)
		}
	}
}
