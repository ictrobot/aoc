package day13

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/ictrobot/aoc/internal/util/parse"
	"reflect"
	"slices"
)

//go:embed example
var Example string

type Day13 struct {
	packets []interface{}
}

func (d *Day13) Parse(input string) {
	d.packets = nil
	for i, l := range parse.Lines(input) {
		if l == "" {
			continue
		}
		var p []interface{}
		if err := json.Unmarshal([]byte(l), &p); err != nil {
			panic(fmt.Errorf("reading line %d: %w", i, err))
		}
		d.packets = append(d.packets, p)
	}
}

func (d *Day13) ParseExample() {
	d.Parse(Example)
}

func (d *Day13) Part1() any {
	var sum int
	for i := 0; i < len(d.packets)-1; i += 2 {
		if o := compare(d.packets[i], d.packets[i+1]); o == correct {
			sum += 1 + (i / 2)
		}
	}
	return sum
}

func (d *Day13) Part2() any {
	divider1 := []interface{}{[]interface{}{2.0}}
	divider2 := []interface{}{[]interface{}{6.0}}
	packets := append([]interface{}{divider1, divider2}, d.packets...)

	slices.SortFunc(packets, func(a, b interface{}) int {
		return int(compare(a, b))
	})

	var pos1, pos2 int
	for i, s := range packets {
		if reflect.DeepEqual(s, divider1) {
			pos1 = i + 1
		}
		if reflect.DeepEqual(s, divider2) {
			pos2 = i + 1
		}
	}

	return pos1 * pos2
}

type order int

const (
	correct   order = -1
	equal     order = 0
	incorrect order = 1
)

func compare(left, right interface{}) order {
	lNum, lIsNum := left.(float64)
	rNum, rIsNum := right.(float64)

	lList, lIsList := left.([]interface{})
	rList, rIsList := right.([]interface{})

	if lIsNum && rIsNum {
		if lNum < rNum {
			return correct
		} else if lNum > rNum {
			return incorrect
		}
		return equal
	} else if lIsList && rIsList {
		for i := 0; i < min(len(lList), len(rList)); i++ {
			if o := compare(lList[i], rList[i]); o != equal {
				return o
			}
		}

		if len(lList) < len(rList) {
			return correct
		} else if len(rList) < len(lList) {
			return incorrect
		}
		return equal
	} else if lIsList {
		return compare(lList, []interface{}{right})
	} else if rIsList {
		return compare([]interface{}{left}, rList)
	} else {
		panic(fmt.Errorf("unknown types: %#v & %#v", left, right))
	}
}
