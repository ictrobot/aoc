package cmd

import (
	"fmt"
	"github.com/ictrobot/aoc/internal/api"
	"github.com/ictrobot/aoc/internal/solution"
	"os"
	"runtime/pprof"
	"time"
)

const (
	pprofSeconds  = 30
	pprofFilename = "aoc.prof"
)

func (o *Options) Run() {
	if o.Profile {
		f, err := os.Create(pprofFilename)
		if err != nil {
			panic(fmt.Errorf("could not create profile file: %w", err))
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			panic(fmt.Errorf("starting profile: %w", err))
		}
		defer pprof.StopCPUProfile()

		start := time.Now()
		for time.Since(start).Seconds() < pprofSeconds {
			o.runSolutions()
		}

		return
	}

	o.runSolutions()
}

func (o *Options) runSolutions() {
	if o.Year != 0 && o.Day != 0 {
		o.runSolution(o.Year, o.Day)
		return
	}

	var years []int
	if o.Year == 0 {
		years = solution.Years()
	} else {
		years = []int{o.Year}
	}

	for i, y := range years {
		if i > 0 {
			fmt.Println()
		}
		for j, d := range solution.Days(y) {
			if j > 0 {
				fmt.Println()
			}
			fmt.Printf("%d day %02d:\n", y, d)

			o.runSolution(y, d)
		}
	}
}

func (o *Options) runSolution(year, day int) {
	s := solution.For(year, day)
	if o.UseExampleInput {
		s2, twoExamples := s.(solution.TwoExamples)

		if o.Part1 || (o.Part2 && !twoExamples) {
			t := time.Now()
			s.ParseExample()
			o.printTiming(year, day, t, "example parsing")
			o.printParsed(year, day, s)
		}

		if o.Part1 {
			t := time.Now()
			p1 := s.Part1()
			o.printTiming(year, day, t, "example part 1")
			fmt.Println(p1)
		}

		if o.Part2 {
			if twoExamples {
				t := time.Now()
				s2.ParseExample2()
				o.printTiming(year, day, t, "example 2 parsing")
				o.printParsed(year, day, s)
			}

			t := time.Now()
			p2 := s.Part2()
			o.printTiming(year, day, t, "example part 2")
			fmt.Println(p2)
		}
	} else {
		input, err := api.GetInput(year, day)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "getting input for %d day %02d: %v\n", year, day, err)
			os.Exit(1)
		}

		t := time.Now()
		s.Parse(input)
		o.printTiming(year, day, t, "parsing")
		o.printParsed(year, day, s)

		if o.Part1 {
			t = time.Now()
			p1 := s.Part1()
			o.printTiming(year, day, t, "part 1")
			fmt.Println(p1)
		}

		if o.Part2 {
			t = time.Now()
			p2 := s.Part2()
			o.printTiming(year, day, t, "part 2")
			fmt.Println(p2)
		}
	}
}

func (o *Options) printTiming(year, day int, t time.Time, msg string) {
	if !o.PrintTiming {
		return
	}

	_, _ = fmt.Fprintf(
		os.Stderr,
		"%d day %02d %s: %.3fms\n",
		year,
		day,
		msg,
		float64(time.Since(t).Nanoseconds())/1000000.0,
	)
}

func (o *Options) printParsed(year, day int, s solution.Solution) {
	if !o.PrintParsed {
		return
	}

	_, _ = fmt.Fprintf(
		os.Stderr,
		"%d day %02d parsed: %+v\n",
		year,
		day,
		s,
	)
}
