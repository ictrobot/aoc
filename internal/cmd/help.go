package cmd

import (
	"fmt"
	"github.com/ictrobot/aoc/internal/solution"
	"io"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

func printHelp(error string) {
	w := os.Stdout
	if error != "" {
		w = os.Stderr
		_, _ = fmt.Fprintf(w, "%s\n\n", error)
	}

	_, _ = fmt.Fprintln(w, `Usage:
	aoc
		run all solutions

	aoc $year
		run all solutions for the given year

	aoc $year $day
		run the solution for the given year & day

	aoc $year $day $part
		run the solution for the given year, day & part`)

	if error != "" {
		_, _ = fmt.Fprintln(w, "\nSee --help for more information")

		os.Exit(2)
	}

	_, _ = fmt.Fprintf(w, `
Options:
	--example/-e
		use the example input, do not use previously downloaded inputs or
		attempt to download inputs using a session token (read from the
		AOC_TOKEN environment variable, ~/.aoc_token, or stdin)

	--timing/-t
		print how long it takes to parse the input & solve each part

	--parsed/-p
		print the parsed input

	--pprof
		run the specified solutions in a loop for %d seconds and write a cpu
		profile to %s

	--help/-h
		print help`, pprofSeconds, pprofFilename)

	printSolutions(w)
	printBuild(w)

	os.Exit(0)
}

func printSolutions(w io.Writer) {
	years := solution.Years()
	if len(years) == 0 {
		return
	}

	b := strings.Builder{}
	b.WriteString("\nSolutions:\n")

	for _, year := range years {
		b.WriteByte('\t')
		b.WriteString(strconv.Itoa(year))
		b.WriteString(": ")

		days := solution.Days(year)
		for i, day := range days {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(strconv.Itoa(day))
		}

		b.WriteByte('\n')
	}

	_, _ = fmt.Fprint(w, b.String())
}

func printBuild(w io.Writer) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	var revision, time, modified string
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			revision = s.Value
		case "vcs.time":
			time = s.Value
		case "vcs.modified":
			modified = s.Value
		}
	}

	s := "\n" + info.Path + " "

	if revision != "" {
		s += "(" + revision
		if time != "" {
			s += " " + time
		}
		if modified != "" && modified != "false" {
			s += " modified"
		}
		s += ") "
	} else if info.Main.Version != "" && info.Main.Version != "(devel)" {
		s += "(" + info.Main.Version + ") "
	}

	s += info.GoVersion + "\n"

	_, _ = fmt.Fprint(w, s)
}
