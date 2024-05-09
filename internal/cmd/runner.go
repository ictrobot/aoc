package cmd

import (
	"fmt"
	"github.com/ictrobot/aoc-go/internal/api"
	"github.com/ictrobot/aoc-go/internal/solution"
	"os"
	"runtime/pprof"
	"time"
)

const (
	pprofSeconds       = 30
	pprofMinIterations = 100
	pprofFilename      = "aoc.pprof"
)

func (o *Options) Run() {
	o.runSolutions()

	if o.Profile {
		// running the solutions once before starting profiling ensures we
		// don't include e.g. prompting for session token or downloading inputs

		// disable printing to avoid printing dominating profiles for fast solutions
		_, _ = fmt.Fprintln(os.Stderr, "\noutput suppressed for future iterations")
		profileOpt := *o
		profileOpt.Silent = true

		f, err := os.Create(pprofFilename)
		if err != nil {
			panic(fmt.Errorf("could not create profile file: %w", err))
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			panic(fmt.Errorf("starting profile: %w", err))
		}
		defer pprof.StopCPUProfile()

		iterations := 0
		start := time.Now()
		for iterations < pprofMinIterations || time.Since(start).Seconds() < pprofSeconds {
			profileOpt.runSolutions()
			iterations++
		}

		duration := time.Since(start)
		_, _ = fmt.Fprintf(os.Stderr,
			"completed %d iterations in %.2fs (avg %.3fms)\n",
			iterations,
			duration.Seconds(),
			duration.Seconds()*1000/float64(iterations),
		)
	}
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
		if i > 0 && !o.Silent {
			fmt.Println()
		}
		for j, d := range solution.Days(y) {
			if j > 0 && !o.Silent {
				fmt.Println()
			}
			if !o.Silent {
				fmt.Printf("%d day %02d:\n", y, d)
			}

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
			o.printResult(p1)
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
			o.printResult(p2)
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
			o.printResult(p1)
		}

		if o.Part2 {
			t = time.Now()
			p2 := s.Part2()
			o.printTiming(year, day, t, "part 2")
			o.printResult(p2)
		}
	}
}

func (o *Options) printTiming(year, day int, t time.Time, msg string) {
	if o.Silent || !o.PrintTiming {
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
	if o.Silent || !o.PrintParsed {
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

func (o *Options) printResult(r any) {
	if o.Silent {
		return
	}

	_, _ = fmt.Println(r)
}
