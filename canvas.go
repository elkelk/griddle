package griddle

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"
	"math"
)

type Canvas struct {
	Width, Height, CanvasWidth, CanvasHeight float64
}

type Element func(gc *draw2dimg.GraphicContext, c Coord, box float64)

func (c Canvas) Fill(gc *draw2dimg.GraphicContext, fill color.RGBA) {
	gc.SetFillColor(fill)
	gc.MoveTo(0, 0)
	gc.LineTo(c.CanvasWidth, 0)
	gc.LineTo(c.CanvasWidth, c.CanvasHeight)
	gc.LineTo(0, c.CanvasHeight)
	gc.Close()
	gc.Fill()
}

func (c Canvas) DrawGrid(gc *draw2dimg.GraphicContext, element Element, box, variance float64) {
	gridCoords := c.GenerateGrid(0, box)

	for i := 0; i < len(gridCoords); i++ {
		gridCoord := gridCoords[i]
		varyCoord := gridCoord
		if variance > 0 {
			varyCoord = gridCoord.withVariance(variance)
		}
		element(gc, varyCoord, box)
	}
}

func (c Canvas) GenerateGrid(offset float64, box float64) []Coord {
	count := (c.CanvasWidth / box) * (c.CanvasHeight / box)
	var coords []Coord
	for i := 0.0; i < count; i += 1 {
		pixelBox := c.CanvasWidth / box
		column := int(i) % int(pixelBox)
		row := int(math.Floor(i / pixelBox))
		itemX := (float64(column) * box) + offset
		itemY := (float64(row) * box) + offset
		coords = append(coords, Coord{itemX, itemY, column, row})
	}
	return coords
}

func (c Canvas) HLine(yOffset float64, box float64) []Coord {
	count := c.CanvasHeight * 4 / box
	var coords []Coord
	for i := 0.0; i < count; i += 1 {
		iBox := i * box
		itemX := 0.0
		itemY := iBox + yOffset
		coords = append(coords, Coord{itemX, itemY, 0, int(i)})
	}
	return coords
}

func (c Canvas) VLine(xOffset float64, box float64) []Coord {
	count := c.CanvasWidth / box
	var coords []Coord
	for i := 0.0; i < count; i += 1 {
		iBox := i * box
		itemX := iBox + xOffset
		itemY := 0.0
		coords = append(coords, Coord{itemX, itemY, int(i), 0})
	}
	return coords
}
