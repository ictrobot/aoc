package solution

//go:generate go run github.com/ictrobot/aoc/internal/generate

type Solution interface {
	Parse(s string)
	ParseExample()
	Part1() interface{}
	Part2() interface{}
}

type TwoExamples interface {
	Solution
	ParseExample2()
}
