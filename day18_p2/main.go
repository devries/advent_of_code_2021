package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/devries/advent_of_code_2021/utils"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Parse()
	f, err := os.Open("../inputs/day18.txt")
	utils.Check(err, "error opening input file")
	defer f.Close()

	r := solve(f, utils.Verbose)
	fmt.Println(r)
}

func solve(r io.Reader, verbose bool) int {
	lines := utils.ReadLines(r)

	elements := make([]Element, len(lines))

	for i, line := range lines {
		elements[i], _ = parseSnailNumber([]rune(line), 0)
	}

	maxValue := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i != j {
				test := add(elements[i], elements[j])
				v := test.Value()
				if v > maxValue {
					maxValue = v
				}
				if verbose {
					fmt.Printf("%s + %s = %s (%d)\n", elements[i], elements[j], test, v)
				}
			}
		}
	}

	return maxValue
}

type Element interface {
	Value() int
	String() string
	Copy() Element
}

type Number struct {
	val int
}

func (n *Number) Value() int {
	return n.val
}

func (n *Number) String() string {
	return fmt.Sprintf("%d", n.val)
}

func (n *Number) Copy() Element {
	return &Number{n.val}
}

type Pair struct {
	left  Element
	right Element
}

func (p *Pair) Value() int {
	return 3*p.left.Value() + 2*p.right.Value()
}

func (p *Pair) String() string {
	return fmt.Sprintf("[%s,%s]", p.left.String(), p.right.String())
}

func (p *Pair) Copy() Element {
	return &Pair{p.left.Copy(), p.right.Copy()}
}

func parseSnailNumber(line []rune, start int) (Element, int) {

	switch line[start] {
	case '[':
		// Pair
		var p Pair
		pos := start + 1
		p.left, pos = parseSnailNumber(line, pos)
		if line[pos] != ',' {
			panic(fmt.Errorf("Expected a comma for second part of pair, got %c", line[pos]))
		}
		pos++
		p.right, pos = parseSnailNumber(line, pos)
		lastrune := line[pos]
		if lastrune != ']' {
			panic(fmt.Errorf("Pair does not end with closing bracket"))
		}
		return &p, pos + 1

	default:
		// Number
		for i := start + 1; true; i++ {
			if line[i] == ',' || line[i] == ']' {
				v, err := strconv.Atoi(string(line[start:i]))
				utils.Check(err, fmt.Sprintf("Unable to parse the number %s", string(line[start:i])))
				return &Number{v}, i
			}
		}
	}

	return nil, 0
}

type ChainLink struct {
	e          Element
	depth      int
	parent     *Pair
	parentSide string
}

// Return all elements in order up to depth 4
func scanElements(e Element, depth int, parent *Pair, side string) []ChainLink {
	switch v := e.(type) {
	case *Pair:
		if depth == 4 {
			return []ChainLink{{v, 4, parent, side}}
		}
		ret := make([]ChainLink, 0)
		leftElements := scanElements(v.left, depth+1, v, "left")
		ret = append(ret, leftElements...)
		rightElements := scanElements(v.right, depth+1, v, "right")
		ret = append(ret, rightElements...)
		return ret
	case *Number:
		cl := ChainLink{v, depth, parent, side}
		return []ChainLink{cl}
	}

	return []ChainLink{}
}

// Explode first value, return true if explosion happens
func explode(e Element) bool {
	chain := scanElements(e, 0, nil, "")

	for i, c := range chain {
		p, ok := c.e.(*Pair)
		if ok {
			leftNumber := p.left.Value()
			rightNumber := p.right.Value()
			switch c.parentSide {
			case "left":
				c.parent.left = &Number{0}
			case "right":
				c.parent.right = &Number{0}
			}
			if i > 0 {
				prev := chain[i-1]
				switch v := prev.e.(type) {
				case *Number:
					v.val += leftNumber
				case *Pair:
					n := v.right.(*Number)
					n.val += leftNumber
				}
			}
			if i < len(chain)-1 {
				next := chain[i+1]
				switch v := next.e.(type) {
				case *Number:
					v.val += rightNumber
				case *Pair:
					n := v.left.(*Number)
					n.val += rightNumber
				}
			}
			return true
		}
	}
	return false
}

func split(e Element) bool {
	chain := scanElements(e, 0, nil, "")

	for _, c := range chain {
		n, ok := c.e.(*Number)
		if ok {
			if n.val > 9 {
				a := n.val / 2
				b := a
				if n.val%2 == 1 {
					b += 1
				}

				pair := &Pair{&Number{a}, &Number{b}}

				switch c.parentSide {
				case "left":
					c.parent.left = pair
				case "right":
					c.parent.right = pair
				}
				return true
			}
		}
	}

	return false
}

func add(a, b Element) Element {
	ret := &Pair{a.Copy(), b.Copy()}

	for {
		explodeTest := explode(ret)
		if explodeTest {
			continue
		}
		splitTest := split(ret)
		if splitTest {
			continue
		}

		break
	}

	return ret
}
