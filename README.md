# griddle
A grid based pattern drawing helper using http://github.com/llgcode/draw2d/draw2dimg

![Top Example Image](https://raw.githubusercontent.com/elkelk/griddle/master/examples/hello-top.png)

## Usage
Import giddle, image, and draw2dimg
```go
import (
	"github.com/elkelk/griddle"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
)
```

Instantiate a griddle canvas by calling `New(Width, Height, CanvasWidth, CanvasHeight)`
```go
  griddle.New(800, 400, 800, 400)
```

Create a GraphicContext to draw with
```go
	dest := image.NewRGBA(image.Rect(0, 0, int(mainCanvas.Width), int(mainCanvas.Height)))
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetDPI(300)
```

Fill the canvas with canvas.Fill
```
	mainCanvas.Fill(gc, color.RGBA{0, 0, 0, 255})
```

Repeat a pattern over a grid with `DrawGrid((*draw2dimg.GraphicContext, Element, box, variance)`
```go
	element := func(gc *draw2dimg.GraphicContext, coord griddle.Coord, box float64) {
		fillColor := color.RGBA{100, 0, 0, 255}
		gc.SetFillColor(fillColor)
		gc.MoveTo(coord.X, coord.Y+box)
		gc.ArcTo(coord.X, coord.Y, box/4, box/4, 0, -math.Pi*2)
		gc.Close()
		gc.FillStroke()
	}

	mainCanvas.DrawGrid(gc, element, 42.0, 0)
```

Simple complete Example:
```go
package main

import (
	"github.com/elkelk/griddle"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"math"
)

func main() {
	mainCanvas := griddle.New(800, 400, 1000, 600)
	dest := image.NewRGBA(image.Rect(0, 0, int(mainCanvas.Width), int(mainCanvas.Height)))
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetDPI(300)

	mainCanvas.Fill(gc, color.RGBA{0, 0, 0, 255})
	element := func(gc *draw2dimg.GraphicContext, coord griddle.Coord, box float64) {
		fillColor := color.RGBA{100, 0, 0, 255}
		gc.SetFillColor(fillColor)
		gc.MoveTo(coord.X+box, coord.Y+box)
		gc.ArcTo(coord.X, coord.Y, box/4, box/4, 0, -math.Pi*2)
		gc.Close()
		gc.FillStroke()
	}

	mainCanvas.DrawGrid(gc, element, 40.0, 0)

	draw2dimg.SaveToPngFile("hello.png", dest)
}
```

![Simple Example Image](https://raw.githubusercontent.com/elkelk/griddle/master/examples/hello.png)

More Complex Example:
```go
package main

import (
	"github.com/elkelk/griddle"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"math"
)

func main() {
	mainCanvas := griddle.New(4200, 2520, 5600, 3360)
	dest := image.NewRGBA(image.Rect(0, 0, int(mainCanvas.Width), int(mainCanvas.Height)))
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetDPI(300)

	mainCanvas.Fill(gc, color.RGBA{0, 0, 0, 255})
	mainCanvas.DrawGrid(gc, elementFunction(mainCanvas), 42.0, 42.0)

	draw2dimg.SaveToPngFile("hello.png", dest)
}

func elementFunction(c griddle.Canvas) griddle.Element {
	return func(gc *draw2dimg.GraphicContext, coord griddle.Coord, box float64) {
		yFade := uint8(coord.Y / c.CanvasHeight)
		high := 255 - yFade
		strokeColor := color.RGBA{0, 7, 120, 1}
		fillColor := color.RGBA{0, 0, high, 255}
		gc.SetFillColor(fillColor)
		gc.SetStrokeColor(strokeColor)
		gc.SetLineWidth(200.0)
		gc.MoveTo(coord.X, coord.Y+box)
		if coord.Column%2 == 0 {
			gc.LineTo(coord.X, coord.Y)
			gc.LineTo(coord.X+c.Height, coord.Y)
			gc.LineTo(coord.X+c.Width, coord.Y+c.Height)
		} else {
			box := 42.0
			boxMultiplier := 0.5 + coord.Y/c.Height
			gc.ArcTo(coord.X, coord.Y, box*boxMultiplier, box*boxMultiplier, 0, -math.Pi*2)
		}
		gc.Close()
		gc.FillStroke()
	}
}
```

![Complex Example Image](https://raw.githubusercontent.com/elkelk/griddle/master/examples/hello-complex.png)
