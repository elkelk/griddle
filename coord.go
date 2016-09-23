package griddle

import (
	"math/rand"
)

type Coord struct {
	X, Y        float64
	Column, Row int
}

func (c Coord) withVariance(variance float64) Coord {
	rSeed := rand.New(rand.NewSource(int64(c.Column + c.Row)))
	xPos := float64(rSeed.Intn(int(variance)))
	xNeg := float64(rSeed.Intn(int(variance)))
	yPos := float64(rSeed.Intn(int(variance)))
	yNeg := float64(rSeed.Intn(int(variance)))
	x := c.X + xPos - xNeg
	y := c.Y + yPos - yNeg
	return Coord{x, y, c.Column, c.Row}
}
