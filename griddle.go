package griddle

func New(width, height, canvasWidth, canvasHeight float64) Canvas {
	return Canvas{
		Width:        width,
		Height:       height,
		CanvasWidth:  canvasWidth,
		CanvasHeight: canvasHeight,
	}
}
