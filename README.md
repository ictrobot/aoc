aoc
===

Advent Of Code (aoc) solutions written in Go.
The solutions are written quickly and not always following best practices (e.g. most errors panic).

This is primarily for my 2023 solutions.
Prior to 2023, I wrote self-contained solutions in TypeScript which only used the Node.js standard library and no other
library of shared helper functions.
Some of these solutions have been ported to Go to try and identify helpful utility functions (see `internal/util`).

## Downloading inputs

By default, the code will attempt to automatically download inputs from adventofcode.com using a session token stored
in the AOC_TOKEN session variable, the `~/.aoc_token` file or prompting for it on stdin (and saving it to
`~/.aoc_token`).
To avoid this, you can use the example inputs `--example` or manually populate the input files (the relevant path should
be printed).

Support for downloading inputs can also be removed using the `nohttp` build tag (e.g. `go build -tags=nohttp .`)

Automatic submission of solutions is not supported.

## Development

`go run ./internal/generate $year $day` will generate placeholder solution files for a given year & day and add
the solution to the list in `./internal/solution/solution_gen.go` automatically.
Not supplying a day will generate files for all released puzzles that year.
It is safe to rerun - existing solution files will not be replaced.

`go generate ./...` can be used to update `./internal/solution/solution_gen.go` without creating new solutions.

I would not recommend using this repository as a template as many functions aren't documented, and it was not designed
as one.
However, it is possible to do so by deleting my solutions using `rm -r ./internal/aoc*` followed by resetting the
solution list using `go generate ./...`.
My utility functions can also be deleted using `rm -r ./internal/util`.
