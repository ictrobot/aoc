package day24

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/numbers"
	"github.com/ictrobot/aoc/internal/util/parse"
	"math"
)

//go:embed example
var Example string

type Day24 struct {
	hailstones []hailstone
	example    bool
}

type hailstone struct {
	pos, vel struct {
		X, Y, Z float64
	}
}

const epsilon = 1e-14

func (d *Day24) Parse(input string) {
	nums := parse.ExtractFloat64s(input)

	d.hailstones = make([]hailstone, 0, len(nums)/6)
	d.example = false

	for i := 0; i < len(nums)-5; i += 6 {
		d.hailstones = append(d.hailstones, hailstone{
			struct{ X, Y, Z float64 }{nums[i+0], nums[i+1], nums[i+2]},
			struct{ X, Y, Z float64 }{nums[i+3], nums[i+4], nums[i+5]},
		})
	}
}

func (d *Day24) ParseExample() {
	d.Parse(Example)
	d.example = true
}

func (d *Day24) Part1() any {
	testAreaMin, testAreaMax := 200_000_000_000_000., 400_000_000_000_000.
	if d.example {
		testAreaMin, testAreaMax = 7., 27.
	}

	var count int
	for i, h1 := range d.hailstones {
		for _, h2 := range d.hailstones[i+1:] {
			m1 := h1.vel.Y / h1.vel.X
			m2 := h2.vel.Y / h2.vel.X
			if m1 == m2 || math.IsNaN(m1) || math.IsNaN(m2) || math.IsInf(m1, 0) || math.IsInf(m2, 0) {
				continue
			}

			c1 := h1.pos.Y - m1*h1.pos.X
			c2 := h2.pos.Y - m2*h2.pos.X

			x := (c2 - c1) / (m1 - m2)
			y := m1*x + c1

			t1 := (x - h1.pos.X) / h1.vel.X
			t2 := (x - h2.pos.X) / h2.vel.X
			if x >= testAreaMin && x <= testAreaMax && y >= testAreaMin && y <= testAreaMax && t1 >= 0 && t2 >= 0 {
				count++
			}
		}
	}
	return count
}

func (d *Day24) Part2() any {
	// iterate over all possible (vx, vy) in increasing absolute sum order
	for sum := 0; sum <= 1_000_000; sum++ {
		for vx := 0; vx <= sum; vx++ {
			vy := sum - vx

			rx, ry, rz, ok := d.findRockPosition(vx, vy)
			if ok {
				return rx + ry + rz
			}

			rx, ry, rz, ok = d.findRockPosition(vx, -vy)
			if ok {
				return rx + ry + rz
			}

			rx, ry, rz, ok = d.findRockPosition(-vx, vy)
			if ok {
				return rx + ry + rz
			}

			rx, ry, rz, ok = d.findRockPosition(-vx, -vy)
			if ok {
				return rx + ry + rz
			}
		}
	}

	panic("no solution found")
}

// findRockPosition returns the starting position of the rock (rx, ry, rz)
// with velocity (vx, vy, vz) which collides with all the hailstones.
// collisions are initially found in the XY plane, and then the corresponding
// vz which would collide at the same time is calculated, meaning only vx & vy
// must be brute forced
func (d *Day24) findRockPosition(vx, vy int) (rx, ry, rz int64, ok bool) {
	// system is overconstrained, finding solution for 3 hailstones should be
	// the solution for all provided such a solution exists
	h1, h2, h3 := d.hailstones[0], d.hailstones[1], d.hailstones[2]

	// subtract thrown rock's XY velocity from each hailstone velocity so that
	// the rock has zero relative XY velocity and so stays at a constant XY
	// position, which we can find by checking where the paths collide
	h1.vel.X -= float64(vx)
	h2.vel.X -= float64(vx)
	h3.vel.X -= float64(vx)
	h1.vel.Y -= float64(vy)
	h2.vel.Y -= float64(vy)
	h3.vel.Y -= float64(vy)

	m1 := h1.vel.Y / h1.vel.X
	m2 := h2.vel.Y / h2.vel.X
	m3 := h3.vel.Y / h3.vel.X
	if withinEpsilon(m1, m2) || withinEpsilon(m1, m3) || withinEpsilon(m2, m3) ||
		math.IsNaN(m1) || math.IsNaN(m2) || math.IsNaN(m3) ||
		math.IsInf(m1, 0) || math.IsInf(m2, 0) || math.IsInf(m3, 0) {
		// at least two of the paths never intersect, no solution
		return 0, 0, 0, false
	}

	c1 := h1.pos.Y - m1*h1.pos.X
	c2 := h2.pos.Y - m2*h2.pos.X
	c3 := h3.pos.Y - m3*h3.pos.X

	// calculate collision location in XY plane for each pair
	x1 := (c2 - c1) / (m1 - m2)
	y1 := m1*x1 + c1
	x2 := (c3 - c1) / (m1 - m3)
	y2 := m1*x2 + c1
	x3 := (c3 - c2) / (m2 - m3)
	y3 := m2*x3 + c2
	if !withinEpsilon(x1, x2) || !withinEpsilon(x1, x3) ||
		!withinEpsilon(y1, y2) || !withinEpsilon(y1, y3) {
		// paths collide at different XY coordinates, no solution
		return 0, 0, 0, false
	}

	// calculate time each hailstone reach XY collision point
	t1 := (x1 - h1.pos.X) / h1.vel.X
	t2 := (x1 - h2.pos.X) / h2.vel.X
	t3 := (x1 - h3.pos.X) / h3.vel.X
	if t1 < 0 || t2 < 0 || t3 < 0 {
		// paths collide in the past, no solution
		return 0, 0, 0, false
	}

	// now in 3D, calculate what vz must be for each pair in order for them to
	// reach the same collision Z at the same time as reaching XY collision
	//
	// - p_i = hailstone i z position
	// - v_i = hailstone i z velocity
	// - t_i = hailstone i collision time
	//
	// p_i + t_i*(v_i - vz) = rz
	// p_i + t_i*(v_i - vz) = p_j + t_j*(v_j - vz)
	// vz = (p_i - p_j + t_i*v_i - t_j*v_j) / (t_i - t_j)
	vz1 := (h1.pos.Z - h2.pos.Z + t1*h1.vel.Z - t2*h2.vel.Z) / (t1 - t2)
	vz2 := (h1.pos.Z - h3.pos.Z + t1*h1.vel.Z - t3*h3.vel.Z) / (t1 - t3)
	vz3 := (h3.pos.Z - h2.pos.Z + t3*h3.vel.Z - t2*h2.vel.Z) / (t3 - t2)
	if !withinEpsilon(vz1, vz2) || !withinEpsilon(vz1, vz3) {
		// would require different vz values, no solution
		return 0, 0, 0, false
	}

	// calculate collision z coordinate for each hailstone
	z1 := h1.pos.Z + t1*(h1.vel.Z-vz1)
	z2 := h2.pos.Z + t2*(h2.vel.Z-vz1)
	z3 := h3.pos.Z + t3*(h3.vel.Z-vz1)
	if !withinEpsilon(z1, z2) || !withinEpsilon(z2, z3) {
		// pretty sure this can't happen given above simultaneous equations
		// equate collision Z, but just in case
		return 0, 0, 0, false
	}

	// check we haven't exceeded the max int value which can be stored in a
	// float without loss of precision
	if x1 < -numbers.Float64MaxInt || x1 > numbers.Float64MaxInt ||
		y1 < -numbers.Float64MaxInt || y1 > numbers.Float64MaxInt ||
		z1 < -numbers.Float64MaxInt || z1 > numbers.Float64MaxInt {
		panic("solution too large, loss of precision")
	}

	// fmt.Printf("(%f, %f, %f) (%d, %d, %f)\n", x1, y1, z1, vx, vy, vz1)
	return int64(x1), int64(y1), int64(z1), true
}

func withinEpsilon(a, b float64) bool {
	return a*(1-epsilon) <= b && a*(1+epsilon) >= b
}
