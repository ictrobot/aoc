package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func generateDay(baseDir string, year, day int) error {
	dir := filepath.Join(baseDir, "internal", fmt.Sprintf("aoc%d", year), fmt.Sprintf("day%02d", day))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	dayGo := filepath.Join(dir, fmt.Sprintf("day%02d.go", day))
	if _, err := os.Stat(dayGo); errors.Is(err, os.ErrNotExist) {
		if err := generateDayGo(dayGo, year, day); err != nil {
			return err
		}
	}

	dayTest := filepath.Join(dir, fmt.Sprintf("day%02d_test.go", day))
	if _, err := os.Stat(dayTest); errors.Is(err, os.ErrNotExist) {
		if err := generateDayTest(dayTest, year, day); err != nil {
			return fmt.Errorf("creating test file: %w", err)
		}
	}

	dayExample := filepath.Join(dir, "example")
	f, err := os.OpenFile(dayExample, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("creating example file: %w", err)
	}
	if err := f.Close(); err != nil {
		return err
	}

	// create input file ahead of time, so file can be already open in editor before it releases
	dayInput := filepath.Join(baseDir, "inputs", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d", day))
	f, err = os.OpenFile(dayInput, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("creating day input: %w", err)
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func generateDayGo(path string, year, day int) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	if err := dayTemplate.Execute(f, struct{ Year, Day int }{year, day}); err != nil {
		return err
	}

	return nil
}

func generateDayTest(path string, year, day int) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	if err := dayTestTemplate.Execute(f, struct{ Year, Day int }{year, day}); err != nil {
		return err
	}

	return nil
}

var dayTemplate = template.Must(template.New("").Parse(`package day{{ printf "%02d" .Day }}

import _ "embed"

//go:embed example
var Example string

type Day{{ printf "%02d" .Day }} struct {
}

func (d *Day{{ printf "%02d" .Day }}) Parse(input string) {
	
}

func (d *Day{{ printf "%02d" .Day }}) ParseExample() {
	d.Parse(Example)
}

func (d *Day{{ printf "%02d" .Day }}) Part1() any {
	return 0
}

func (d *Day{{ printf "%02d" .Day }}) Part2() any {
	return 0
}
`))

var dayTestTemplate = template.Must(template.New("").Parse(`package day{{ printf "%02d" .Day }}

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 0
const Part2 = 0

func TestDay{{ printf "%02d" .Day }}_ParseExample(t *testing.T) {
	d1 := Day{{ printf "%02d" .Day }}{}
	d1.ParseExample()

	d2 := Day{{ printf "%02d" .Day }}{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay{{ printf "%02d" .Day }}_ParseExample(b *testing.B) {
	d := Day{{ printf "%02d" .Day }}{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay{{ printf "%02d" .Day }}_Part1(t *testing.T) {
	d := Day{{ printf "%02d" .Day }}{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay{{ printf "%02d" .Day }}_Part1(b *testing.B) {
	d := Day{{ printf "%02d" .Day }}{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay{{ printf "%02d" .Day }}_Part2(t *testing.T) {
	d := Day{{ printf "%02d" .Day }}{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay{{ printf "%02d" .Day }}_Part2(b *testing.B) {
	d := Day{{ printf "%02d" .Day }}{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
`))
