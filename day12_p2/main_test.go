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
		{testInputSmall, 36},
		{testInputMedium, 103},
		{testInputLarge, 3509},
	}

	for i, test := range tests {
		r := strings.NewReader(test.input)

		t.Logf("Test %d", i+1)
		result := solve(r, testing.Verbose())

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}

var bmarkInput = `TR-start
xx-JT
xx-TR
hc-dd
ab-JT
hc-end
dd-JT
ab-dd
TR-ab
vh-xx
hc-JT
TR-vh
xx-start
hc-ME
vh-dd
JT-bm
end-ab
dd-xx
end-TR
hc-TR
start-vh`

func BenchmarkSolition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(bmarkInput)
		solve(r, false)
	}
}
