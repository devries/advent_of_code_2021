package main

import (
	"math/big"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input  string
		answer *big.Int
	}{
		{"C200B40A82", big.NewInt(3)},
		{"04005AC33890", big.NewInt(54)},
		{"880086C3E88112", big.NewInt(7)},
		{"CE00C43D881120", big.NewInt(9)},
		{"D8005AC2A8F0", big.NewInt(1)},
		{"F600BC2D8F", big.NewInt(0)},
		{"9C005AC2F8F0", big.NewInt(0)},
		{"9C0141080250320F1802104A08", big.NewInt(1)},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := solve(r)

		if result.Cmp(test.answer) != 0 {
			t.Errorf("%s: expected %s, got %s", test.input, test.answer, result)
		}
	}
}
