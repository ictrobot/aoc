package cmd

import (
	"fmt"
	"github.com/ictrobot/aoc/internal/solution"
	"os"
	"strconv"
)

type Options struct {
	Year, Day                int
	Part1, Part2             bool
	UseExampleInput          bool
	PrintTiming, PrintParsed bool
	Profile                  bool
}

func ParseOptions() *Options {
	o := Options{Part1: true, Part2: true}
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]

		// long options
		if len(arg) > 2 && arg[0] == '-' && arg[1] == '-' {
			switch arg[2:] {
			case "time", "timing":
				o.PrintTiming = true
			case "example", "examples":
				o.UseExampleInput = true
			case "parsed":
				o.PrintParsed = true
			case "pprof":
				o.Profile = true
			case "help":
				printHelp("")
			default:
				printHelp(fmt.Sprintf("Unknown long option `%s`", arg[2:]))
			}
			continue
		}

		// short options
		if len(arg) > 1 && arg[0] == '-' {
			for j := 1; j < len(arg); j++ {
				switch arg[j] {
				case 't':
					o.PrintTiming = true
				case 'e':
					o.UseExampleInput = true
				case 'p':
					o.PrintParsed = true
				case 'h':
					printHelp("")
				default:
					printHelp(fmt.Sprintf("Unknown short option `%c`", arg[j]))
				}
			}
			continue
		}

		if o.Year == 0 {
			y, err := strconv.ParseInt(arg, 10, 0)
			if err != nil || len(solution.Days(int(y))) == 0 {
				printHelp(fmt.Sprintf("Unknown year %d", y))
			}
			o.Year = int(y)
			continue
		}

		if o.Day == 0 {
			d, err := strconv.ParseInt(arg, 10, 0)
			if err != nil || solution.For(o.Year, int(d)) == nil {
				printHelp(fmt.Sprintf("Unknown day %d day %02d", o.Year, d))
			}
			o.Day = int(d)
			continue
		}

		if o.Part1 && o.Part2 {
			switch arg {
			case "1":
				o.Part2 = false
			case "2":
				o.Part1 = false
			default:
				printHelp(fmt.Sprintf("Unknown part `%s`", arg))
			}
			continue
		}

		printHelp(fmt.Sprintf("Unknown argument `%s`", arg))
	}

	return &o
}
