package day15

import (
	_ "embed"
	"strings"
)

//go:embed example
var Example string

type Day15 struct {
	steps []string
}

type lens struct {
	label       string
	focalLength uint8
}

func (d *Day15) Parse(input string) {
	d.steps = strings.Split(strings.TrimSpace(input), ",")
}

func (d *Day15) ParseExample() {
	d.Parse(Example)
}

func (d *Day15) Part1() any {
	sum := uint32(0)
	for _, step := range d.steps {
		sum += uint32(hash(step))
	}
	return sum
}

func (d *Day15) Part2() any {
	boxes := make([][]lens, 256)

	// preallocate each slice with cap=4
	flat := make([]lens, len(boxes)*4)
	for i := 0; i < len(boxes); i++ {
		boxes[i] = flat[i*4 : i*4 : (i+1)*4]
	}

steps:
	for _, step := range d.steps {
		if step[len(step)-1] == '-' {
			// remove lens from box
			label := step[:len(step)-1]
			box := hash(label)

			for i := range boxes[box] {
				if boxes[box][i].label == label {
					boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
					break
				}
			}
		} else {
			// replace lens with same label or append to box
			focal := step[len(step)-1] - '0'
			label := step[:len(step)-2]
			box := hash(label)

			for i := range boxes[box] {
				if boxes[box][i].label == label {
					boxes[box][i].focalLength = focal
					continue steps
				}
			}
			boxes[box] = append(boxes[box], lens{label: label, focalLength: focal})
		}
	}

	var sum int
	for i := range boxes {
		for j, l := range boxes[i] {
			sum += (1 + i) * (1 + j) * int(l.focalLength)
		}
	}
	return sum
}

func hash(s string) (h uint8) {
	for i := 0; i < len(s); i++ {
		h += s[i]
		h *= 17
		// h %= 256 not needed as go spec says unsigned int + - * are computed 2**n
	}
	return
}
