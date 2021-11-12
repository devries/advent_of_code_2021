package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestGcd(t *testing.T) {
	var tests = []struct {
		a      int64
		b      int64
		answer int64
	}{
		{1, 1, 1},
		{10, 5, 5},
		{12, 16, 4},
	}

	for _, test := range tests {
		result := Gcd(test.a, test.b)
		if result != test.answer {
			t.Errorf("For values %d and %d calculated %d, expected %d", test.a, test.b, result, test.answer)
		}
	}
}

func TestLcm(t *testing.T) {
	var tests = []struct {
		a      int64
		b      int64
		answer int64
	}{
		{1, 1, 1},
		{10, 5, 10},
		{12, 16, 48},
	}

	for _, test := range tests {
		result := Lcm(test.a, test.b)
		if result != test.answer {
			t.Errorf("For values %d and %d calculated %d, expected %d", test.a, test.b, result, test.answer)
		}
	}
}

func TestCountBits(t *testing.T) {
	var tests = []struct {
		n    uint32
		bits int
	}{
		{0b0, 0},
		{0b10, 1},
		{0b1011010110, 6},
	}

	for _, test := range tests {
		result := CountBits(test.n)
		if result != test.bits {
			t.Errorf("For bitfield %b and calculated %d bits, expected %d bits", test.n, result, test.bits)
		}
	}
}

func TestPoint(t *testing.T) {
	p := Point{0, 0}

	p2 := p.Add(North)
	p3 := p.Add(East)

	p4 := North.Add(South)

	if p2 != North {
		t.Errorf("0,0 + North should be North")
	}

	if p3 != East {
		t.Errorf("0,0 + East should be East")
	}

	if p4 != p {
		t.Errorf("North + South should be 0,0")
	}
}

func TestPermutations(t *testing.T) {
	a := []int{1, 2, 3}

	perms := [][]int{
		{1, 2, 3},
		{2, 1, 3},
		{3, 1, 2},
		{1, 3, 2},
		{2, 3, 1},
		{3, 2, 1},
	}

	for idx := range Permutations(len(a)) {
		for i, r2 := range perms {
			if a[idx[0]] == r2[0] && a[idx[1]] == r2[1] && a[idx[2]] == r2[2] {
				perms[i] = perms[len(perms)-1]
				perms[len(perms)-1] = nil
				perms = perms[:len(perms)-1]
				break
			}
		}
	}

	if len(perms) > 0 {
		t.Errorf("Not all permutations of %v were found", a)
	}
}

func TestCombinations(t *testing.T) {
	a := []string{"red", "green", "blue"}

	combs := [][]string{
		{"red", "green"},
		{"red", "blue"},
		{"green", "blue"},
	}

	for idx := range Combinations(len(a), 2) {
		fmt.Println(idx)
		for i, r := range combs {
			d := make(map[string]bool)
			for _, r2 := range r {
				d[r2] = true
			}

			if d[a[idx[0]]] && d[a[idx[1]]] {
				combs[i] = combs[len(combs)-1]
				combs = combs[:len(combs)-1]
				break
			}
		}
	}

	if len(combs) > 0 {
		t.Errorf("Not all combinations of %v were found", a)
	}
}

func TestCheck(t *testing.T) {
	Check(nil, "test no error") // This should not panic

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The Check function did not panic")
		}
	}()

	err := fmt.Errorf("generic error")

	Check(err, "test error")
}

func TestReadLines(t *testing.T) {
	test := struct {
		Input  string
		Output []string
	}{"one\ntwo\nthree\n", []string{"one", "two", "three"}}

	r := strings.NewReader(test.Input)
	result := ReadLines(r)

	for i, v := range result {
		if v != test.Output[i] {
			t.Errorf("Expected %v, got %v", test.Output, result)
		}
	}
}
