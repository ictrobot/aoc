package day20

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/numbers"
	"github.com/ictrobot/aoc/internal/util/parse"
	"slices"
	"strings"
)

//go:embed example
var Example string

type Day20 struct {
	modules  map[string]*module
	stateLen int
	example  bool
}

type module struct {
	name       string
	class      moduleClass
	inputs     []*module
	outputs    []*module
	stateIndex int
}

type moduleClass uint8

const (
	broadcast moduleClass = iota
	flipFlop
	conjunction
)

type pulse struct {
	high     bool
	src, dst *module
}

func (d *Day20) Parse(input string) {
	lines := parse.Grid(parse.Lines, parse.Whitespace, strings.NewReplacer("->", "", ",", "").Replace(input))

	d.modules = make(map[string]*module)
	d.stateLen = 0
	d.example = false

	// first iterate over both sides to ensure we make *module instances
	// for modules that only appear on the rhs
	for _, line := range lines {
		for i, name := range line {
			var class moduleClass
			if i == 0 && name[0] == '%' {
				class = flipFlop

				name = name[1:]
				line[0] = name
			} else if i == 0 && name[0] == '&' {
				class = conjunction

				name = name[1:]
				line[0] = name
			}

			if m, exists := d.modules[name]; !exists {
				d.modules[name] = &module{name: name, class: class}
			} else if i == 0 {
				m.class = class
			}
		}
	}

	// then iterate over each line again to populate the input & output slices
	// with pointers
	for _, line := range lines {
		sender := d.modules[line[0]]
		for _, name := range line[1:] {
			receiver := d.modules[name]
			sender.outputs = append(sender.outputs, receiver)
			receiver.inputs = append(receiver.inputs, sender)
		}
	}

	// then allocate state offsets once we know each module's input count
	for _, line := range lines {
		m := d.modules[line[0]]
		if m.class == flipFlop {
			m.stateIndex = d.stateLen
			d.stateLen++
		} else if m.class == conjunction {
			m.stateIndex = d.stateLen
			d.stateLen += len(m.inputs)
		}
	}
}

func (d *Day20) ParseExample() {
	d.Parse(Example)
	d.example = true
}

func (d *Day20) Part1() any {
	state := make([]bool, d.stateLen)

	var highCount, lowCount int64
	for iteration := 1; iteration <= 1000; iteration++ {
		d.simulateButtonPress(state, func(p pulse) {
			if p.high {
				highCount++
			} else {
				lowCount++
			}
		})
	}

	return lowCount * highCount
}

func (d *Day20) Part2() any {
	if d.example {
		return "no example for part 2"
	}

	// rx has one input "m", which is a conjunction. For rx to get a low pulse
	// "m" must have all high inputs. Calculate the LCM of the cycle lengths
	// for each of the inputs to "m" sending a high pulse to find the first
	// iteration where it would be possible for all the inputs to high at once
	m := d.modules["rx"].inputs[0]
	cycleLengths := make([]int64, len(m.inputs))
	cyclesFound := 0

	state := make([]bool, d.stateLen)
	for iteration := int64(1); cyclesFound < len(m.inputs); iteration++ {
		d.simulateButtonPress(state, func(p pulse) {
			if !p.high || p.dst != m {
				return
			}

			idx := slices.Index(m.inputs, p.src)
			if cycleLengths[idx] > 0 {
				return
			}

			cycleLengths[idx] = iteration
			cyclesFound++
		})
	}

	lcm := int64(1)
	for _, cycleLen := range cycleLengths {
		lcm = numbers.LCM(lcm, cycleLen)
	}
	return lcm
}

func (d *Day20) simulateButtonPress(state []bool, callback func(pulse)) {
	queue := []pulse{{high: false, dst: d.modules["broadcaster"]}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		callback(p)

		sendHigh := p.high
		switch p.dst.class {
		case flipFlop:
			if p.high {
				// flip flops ignore high pulses
				continue
			}

			// invert state and send new state
			newState := !state[p.dst.stateIndex]
			state[p.dst.stateIndex] = newState
			sendHigh = newState

		case conjunction:
			// update state for the input module
			modState := state[p.dst.stateIndex : p.dst.stateIndex+len(p.dst.inputs)]
			modState[slices.Index(p.dst.inputs, p.src)] = p.high

			// if state is all high, then low, else high
			sendHigh = false
			for _, b := range modState {
				if !b {
					sendHigh = true
					break
				}
			}

		default:
			// forward signal unchanged
		}

		for _, o := range p.dst.outputs {
			queue = append(queue, pulse{src: p.dst, dst: o, high: sendHigh})
		}
	}
}
