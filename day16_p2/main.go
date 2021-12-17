package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"

	"github.com/devries/advent_of_code_2021/utils"
	"github.com/spf13/pflag"
)

var operationMap = map[int64]string{
	0: "sum",
	1: "mul",
	2: "min",
	3: "max",
	4: "num",
	5: "gt",
	6: "lt",
	7: "eq",
}

func main() {
	pflag.Parse()
	f, err := os.Open("../inputs/day16.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f, utils.Verbose)
	fmt.Println(r)
}

type Packet interface {
	Version() int64
	Type() int64
	String() string
}

type PacketData struct {
	Value *big.Int
	Size  uint
}

func (p PacketData) Extract(start, length uint) *big.Int {
	r := new(big.Int)
	mask := new(big.Int)
	one := big.NewInt(1)

	mask.Sub(mask.Lsh(one, length), one)

	r.And(r.Rsh(p.Value, p.Size-start-length), mask)

	return r
}

type Literal struct {
	PacketVersion int64
	PacketType    int64
	Value         *big.Int
}

func (l Literal) Version() int64 {
	return l.PacketVersion
}

func (l Literal) Type() int64 {
	return l.PacketType
}

func (l Literal) String() string {
	return fmt.Sprintf("v:%d|t:%d(num)=%s", l.PacketVersion, l.PacketType, l.Value.String())
}

type Operator struct {
	PacketVersion int64
	PacketType    int64
	Packets       []Packet
}

func NewOperator(version int64, packetType int64) *Operator {
	r := Operator{version, packetType, make([]Packet, 0)}

	return &r
}

func (o *Operator) Version() int64 {
	return o.PacketVersion
}

func (o *Operator) Type() int64 {
	return o.PacketType
}

func (o *Operator) String() string {
	parts := make([]string, 0)
	pad := "\t"

	for _, v := range o.Packets {
		sub := v.String()
		subparts := strings.Split(sub, "\n")

		for _, p := range subparts {
			parts = append(parts, fmt.Sprintf("%s%s", pad, p))
		}
	}
	return fmt.Sprintf("v:%d|t:%d(%s)[\n%s\n]", o.PacketVersion, o.PacketType, operationMap[o.PacketType], strings.Join(parts, "\n"))
}

func solve(r io.Reader, verbose bool) *big.Int {
	lines := utils.ReadLines(r)
	pd := parseFile(lines)

	pk, _ := parsePacket(pd, 0)
	if verbose {
		fmt.Println(pk)
	}

	return calculate(pk)

}

func parseFile(lines []string) PacketData {
	v := lines[0]
	z := new(big.Int)

	z.SetString(v, 16)
	p := PacketData{z, uint(len(v) * 4)}

	return p
}

func parsePacket(pd PacketData, start uint) (Packet, uint) {
	version := pd.Extract(start, 3).Int64()
	start += 3
	packetType := pd.Extract(start, 3).Int64()
	start += 3

	if packetType == 4 {
		// Value Packet
		val := big.NewInt(0)
		for {
			tmp := pd.Extract(start, 5).Int64()
			start += 5
			if tmp >= 16 {
				val.Add(val, big.NewInt(tmp&15))
				val.Lsh(val, 4)
			} else {
				val.Add(val, big.NewInt(tmp))
				break
			}
		}
		return Literal{version, packetType, val}, start
	} else {
		// Operator Packet
		lengthType := pd.Extract(start, 1).Int64()
		start++

		oper := NewOperator(version, packetType)
		switch lengthType {
		case 0:
			// 15 bit number representing number of bits is next
			bits := pd.Extract(start, 15).Int64()
			start += 15
			end := start + uint(bits)

			for start < end {
				var p Packet
				p, start = parsePacket(pd, start)
				oper.Packets = append(oper.Packets, p)
			}

		case 1:
			// 11 bit number representing number of subpackets is next
			packets := pd.Extract(start, 11).Int64()
			start += 11

			for i := int64(0); i < packets; i++ {
				var p Packet
				p, start = parsePacket(pd, start)
				oper.Packets = append(oper.Packets, p)
			}
		}

		return oper, start
	}
}

func sumVersions(p Packet) int64 {
	total := int64(0)

	total += p.Version()

	if oper, ok := p.(*Operator); ok {
		for _, subp := range oper.Packets {
			total += sumVersions(subp)
		}
	}

	return total
}

func calculate(p Packet) *big.Int {

	switch v := p.(type) {
	case Literal:
		return v.Value
	case *Operator:
		switch v.Type() {
		case 0:
			// sum
			r := big.NewInt(0)
			for _, sub := range v.Packets {
				r.Add(r, calculate(sub))
			}
			return r

		case 1:
			// product
			r := big.NewInt(1)
			for _, sub := range v.Packets {
				r.Mul(r, calculate(sub))
			}
			return r

		case 2:
			// minimum
			r := new(big.Int)
			first := true
			for _, sub := range v.Packets {
				a := calculate(sub)
				if first {
					r = a
					first = false
				} else {
					if a.Cmp(r) < 0 {
						r = a
					}
				}
			}
			return r

		case 3:
			// maximum
			r := big.NewInt(0)
			for _, sub := range v.Packets {
				a := calculate(sub)
				if a.Cmp(r) > 0 {
					r = a
				}
			}
			return r

		case 5:
			// greater than
			if calculate(v.Packets[0]).Cmp(calculate(v.Packets[1])) > 0 {
				return big.NewInt(1)
			} else {
				return big.NewInt(0)
			}

		case 6:
			// less than
			if calculate(v.Packets[0]).Cmp(calculate(v.Packets[1])) < 0 {
				return big.NewInt(1)
			} else {
				return big.NewInt(0)
			}

		case 7:
			// equal to
			if calculate(v.Packets[0]).Cmp(calculate(v.Packets[1])) == 0 {
				return big.NewInt(1)
			} else {
				return big.NewInt(0)
			}
		}
	}

	return big.NewInt(0)
}
