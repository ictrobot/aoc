package day21

import (
	_ "embed"
	"fmt"
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/samber/lo"
	"maps"
	"math"
	"strings"
)

//go:embed example
var Example string

type Day21 struct {
	constants map[string]float64
	ops       map[string]op
}

type op struct {
	lhs string
	op  string
	rhs string
}

func (d *Day21) Parse(input string) {
	// remove punctuation the parse into grid of lines & fields
	lines := parse.Grid(parse.Lines, parse.Whitespace, strings.ReplaceAll(input, ":", ""))

	d.constants = make(map[string]float64)
	d.ops = make(map[string]op)

	for _, l := range lines {
		if len(l) == 2 {
			d.constants[l[0]] = parse.Float64(l[1])
		} else {
			d.ops[l[0]] = op{l[1], l[2], l[3]}
		}
	}
}

func (d *Day21) ParseExample() {
	d.Parse(Example)
}

func (d *Day21) Part1() any {
	results := evaluate(d.constants, d.ops)

	f := results["root"]
	if f > numbers.Float64MaxInt || f < -numbers.Float64MaxInt {
		panic("result too large, loss of precision occurred")
	}
	return int64(f)
}

func (d *Day21) Part2() any {
	// use NaN to represent unknown value / depends on human as it
	// propagates through adding/subtracting/multiplying/dividing
	constants := maps.Clone(d.constants)
	constants["humn"] = math.NaN()

	ops := maps.Clone(d.ops)
	ops["root"] = op{ops["root"].lhs, "=", ops["root"].rhs}

	results := evaluate(constants, ops)

	m := "root"
	target := math.NaN()
	for m != "humn" {
		op := ops[m]
		lhs := results[op.lhs]
		rhs := results[op.rhs]

		if math.IsNaN(lhs) && math.IsNaN(rhs) {
			panic("both sides depend on humn")
		} else if math.IsNaN(rhs) {
			// lhs fixed, rhs changed by humn

			switch op.op {
			case "+":
				target = target - lhs
			case "-":
				target = lhs - target
			case "*":
				target = target / lhs
			case "/":
				target = lhs / target
			case "=":
				target = lhs
			default:
				panic(fmt.Errorf("unknown op: %v", op.op))
			}

			m = op.rhs
		} else if math.IsNaN(lhs) {
			// rhs fixed, lhs changed by humn

			switch op.op {
			case "+":
				target = target - rhs
			case "-":
				target = target + rhs
			case "*":
				target = target / rhs
			case "/":
				target = target * rhs
			case "=":
				target = rhs
			default:
				panic(fmt.Errorf("unknown op: %v", op.op))
			}

			m = op.lhs
		} else {
			panic("neither side depends on humn")
		}
	}

	if target > numbers.Float64MaxInt || target < -numbers.Float64MaxInt {
		panic("result too large, loss of precision occurred")
	}
	return int64(target)
}

func evaluate(constants map[string]float64, ops map[string]op) map[string]float64 {
	results := maps.Clone(constants)
	queue := lo.Keys(ops)

	for len(queue) > 0 {
		m := queue[0]
		queue = queue[1:]

		op := ops[m]

		lhs, lhsOk := results[op.lhs]
		rhs, rhsOk := results[op.rhs]
		if !lhsOk || !rhsOk {
			queue = append(queue, m)
			continue
		}

		switch op.op {
		case "+":
			results[m] = lhs + rhs
		case "-":
			results[m] = lhs - rhs
		case "*":
			results[m] = lhs * rhs
		case "/":
			results[m] = lhs / rhs
		case "=":
			// for part 2, skip
		default:
			panic(fmt.Errorf("unknown op: %v", op.op))
		}
	}

	return results
}
