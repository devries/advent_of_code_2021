package main

import (
	"strings"
	"testing"
)

var testInputSmall = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var testInputMedium = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var testInputLarge = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

func TestSolution(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInputSmall, 10},
		{testInputMedium, 19},
		{testInputLarge, 226},
	}

	for i, test := range tests {
		r := strings.NewReader(test.input)

		t.Logf("Test %d:", i+1)
		result := solve(r, testing.Verbose())

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
