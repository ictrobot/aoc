package parse

import (
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
)

type placeholderTest struct {
	_ Placeholder
}

type placeholderTestOptional struct {
	LHS []int
	_   Placeholder `match:"and"`
	RHS *int
	_   Placeholder `flags:"optional"`
}

type placeholderTestMatch struct {
	A int
	_ Placeholder `match:"T"`
	B int
}

type placeholderTestRegex struct {
	A int
	_ Placeholder `regex:"^T"`
	B int
}

type placeholderTestRegexOptional struct {
	_ Placeholder `regex:"^[^0-9]" flags:"optional"`
	I int
}

type placeholderTestOptionalMultiple struct {
	_ Placeholder `match:"m" flags:"optional,multiple"`
	I int
}

func TestReflect_Placeholder(t *testing.T) {
	d1, err := Reflect[placeholderTest]([]string{"value"})
	assert.Zero(t, d1)
	assert.NoError(t, err)
	d2, err := Reflect[placeholderTest]([]string{})
	assert.Zero(t, d2)
	assert.ErrorIs(t, err, NotEnoughStrings)
	d3, err := Reflect[placeholderTest]([]string{"str", "extra"})
	assert.Zero(t, d3)
	assert.ErrorIs(t, err, TooManyStrings)

	o1, err := Reflect[placeholderTestOptional]([]string{"1", "2", "3", "and", "previous"})
	assert.Equal(t, placeholderTestOptional{LHS: []int{1, 2, 3}}, o1)
	assert.NoError(t, err)
	o2, err := Reflect[placeholderTestOptional]([]string{"1", "and", "1000"})
	assert.Equal(t, placeholderTestOptional{LHS: []int{1}, RHS: lo.ToPtr(1000)}, o2)
	assert.NoError(t, err)
	o3, err := Reflect[placeholderTestOptional]([]string{"1", "and", "previous", "extra"})
	assert.Zero(t, o3)
	assert.ErrorIs(t, err, TooManyStrings)

	m1, err := Reflect[placeholderTestMatch]([]string{"1", "T", "2"})
	assert.Equal(t, placeholderTestMatch{A: 1, B: 2}, m1)
	assert.NoError(t, err)
	m2, err := Reflect[placeholderTestMatch]([]string{"1", "t", "2"})
	assert.Zero(t, m2)
	assert.ErrorIs(t, err, PlaceholderDoesNotMatch)

	r1, err := Reflect[placeholderTestRegex]([]string{"1", "T", "2"})
	assert.Equal(t, placeholderTestRegex{A: 1, B: 2}, r1)
	assert.NoError(t, err)
	r2, err := Reflect[placeholderTestRegex]([]string{"1", "Tt", "2"})
	assert.Equal(t, placeholderTestRegex{A: 1, B: 2}, r2)
	assert.NoError(t, err)
	r3, err := Reflect[placeholderTestRegex]([]string{"1", "", "2"})
	assert.Zero(t, r3)
	assert.ErrorIs(t, err, PlaceholderDoesNotMatch)

	ro1, err := Reflect[placeholderTestRegexOptional]([]string{"string", "123"})
	assert.Equal(t, placeholderTestRegexOptional{I: 123}, ro1)
	assert.NoError(t, err)
	ro2, err := Reflect[placeholderTestRegexOptional]([]string{"456"})
	assert.Equal(t, placeholderTestRegexOptional{I: 456}, ro2)
	assert.NoError(t, err)

	om1, err := Reflect[placeholderTestOptionalMultiple]([]string{"m", "m", "m", "101"})
	assert.Equal(t, placeholderTestOptionalMultiple{I: 101}, om1)
	assert.NoError(t, err)
	om2, err := Reflect[placeholderTestOptionalMultiple]([]string{"m", "1001"})
	assert.Equal(t, placeholderTestOptionalMultiple{I: 1001}, om2)
	assert.NoError(t, err)
	om3, err := Reflect[placeholderTestOptionalMultiple]([]string{"1001"})
	assert.Equal(t, placeholderTestOptionalMultiple{I: 1001}, om3)
	assert.NoError(t, err)

	assert.Panics(t, func() {
		_, _ = Reflect[struct{ _ int }]([]string{"1"})
	})
	assert.Panics(t, func() {
		_, _ = Reflect[struct{ X Placeholder }]([]string{"1"})
	})
	assert.Panics(t, func() {
		_, _ = Reflect[struct {
			_ Placeholder `flags:"somethingElse"`
		}]([]string{""})
	})
	assert.Panics(t, func() {
		_, _ = Reflect[struct {
			_ Placeholder `regex:"["`
		}]([]string{""})
	})
	assert.Panics(t, func() {
		_, _ = Reflect[struct {
			_ Placeholder `regex:".+" match:"oops"`
		}]([]string{""})
	})
}
