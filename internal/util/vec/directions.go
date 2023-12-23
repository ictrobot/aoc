package vec

const (
	// NegX is the index of the vector with X=-1 in I2Directions,
	// I2DirectionsWithDiagonals, I3Directions & I3DirectionsWithDiagonals
	NegX = 0
	// PosX is the index of the vector with X=1 in I2Directions,
	// I2DirectionsWithDiagonals, I3Directions & I3DirectionsWithDiagonals
	PosX = 1
	// NegY is the index of the vector with Y=-1 in I2Directions,
	// I2DirectionsWithDiagonals, I3Directions & I3DirectionsWithDiagonals
	NegY = 2
	// PosY is the index of the vector with Y=1 in I2Directions,
	// I2DirectionsWithDiagonals, I3Directions & I3DirectionsWithDiagonals
	PosY = 3
	// NegZ is the index of the vector with Z=-1 in I3Directions
	// & I3DirectionsWithDiagonals
	NegZ = 4
	// PosZ is the index of the vector with Z=1 in I3Directions
	// & I3DirectionsWithDiagonals
	PosZ = 5
)

var I2Directions = [...]I2[int]{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var I2DirectionsWithDiagonals = [...]I2[int]{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	// diagonals
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

// I2Opposites contains the index of the opposite direction for each index in
// I2Directions & I2DirectionsWithDiagonals
var I2Opposites = [...]int{
	1,
	0,
	3,
	2,
	// diagonals
	7,
	6,
	5,
	4,
}

var I3Directions = [...]I3[int]{
	{-1, 0, 0},
	{1, 0, 0},
	{0, -1, 0},
	{0, 1, 0},
	{0, 0, -1},
	{0, 0, 1},
}

var I3DirectionsWithDiagonals = [...]I3[int]{
	{-1, 0, 0},
	{1, 0, 0},
	{0, -1, 0},
	{0, 1, 0},
	{0, 0, -1},
	{0, 0, 1},
	// diagonals
	{-1, -1, -1},
	{-1, -1, 0},
	{-1, -1, 1},
	{-1, 0, -1},
	{-1, 0, 1},
	{-1, 1, -1},
	{-1, 1, 0},
	{-1, 1, 1},
	{0, -1, -1},
	{0, -1, 1},
	{0, 1, -1},
	{0, 1, 1},
	{1, -1, -1},
	{1, -1, 0},
	{1, -1, 1},
	{1, 0, -1},
	{1, 0, 1},
	{1, 1, -1},
	{1, 1, 0},
	{1, 1, 1},
}

// I3Opposites contains the index of the opposite direction for each index in
// I3Directions & I3DirectionsWithDiagonals
var I3Opposites = [...]int{
	1,
	0,
	3,
	2,
	5,
	4,
	// diagonals
	25,
	24,
	23,
	22,
	21,
	20,
	19,
	18,
	17,
	16,
	15,
	14,
	13,
	12,
	11,
	10,
	9,
	8,
	7,
	6,
}
