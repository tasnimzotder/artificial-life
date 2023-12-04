package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/tasnimzotder/artificial-life/constants"
	"image/color"
)

type Game struct {
	App     fyne.App
	Window  fyne.Window
	Canvas  fyne.Canvas
	Widgets *fyne.Container
}

func (g *Game) NewGame() {
	g.App = app.New()
	g.Window = g.App.NewWindow("Hello")
	g.Canvas = g.App.NewWindow("Canvas and Widget Example").Canvas()

	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			return color.RGBA{B: 0, A: 255}
		},
	)

	raster.SetMinSize(fyne.NewSize(constants.CanvasWidth, constants.CanvasHeight))
	g.Canvas.SetContent(raster)
}

func (g *Game) Run() {

	g.Widgets.Layout = layout.NewAdaptiveGridLayout(2)

	content := container.New(
		layout.NewVBoxLayout(),
		g.Canvas.Content(),
		g.Widgets,
	)

	g.Window.SetContent(content)
	//g.Window.Resize(fyne.NewSize(900, 800))
	g.Window.SetFixedSize(true)
	g.Window.SetFixedSize(true)
	g.Window.CenterOnScreen()
	g.Window.SetCloseIntercept(func() {
		g.Window.Close()
	})

	g.Window.ShowAndRun()
}

func (g *Game) Update(grid *[][]uint8) {
	g.DrawGrid(grid)
	//g.Canvas.Refresh(g.Canvas.Content())
}

func (g *Game) DrawGrid(grid *[][]uint8) {
	//log.Printf("Canvas Size: %v", g.Canvas.Size())

	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			canvasHeight := g.Canvas.Size().Height
			canvasWidth := g.Canvas.Size().Width

			heightRatio := float32(h) / canvasHeight
			widthRatio := float32(w) / canvasWidth

			cellHeight := canvasHeight / float32(len(*grid)) * heightRatio
			cellWidth := canvasWidth / float32(len((*grid)[0])) * widthRatio

			gridX := float32(x) / cellWidth
			gridY := float32(y) / cellHeight

			gridWidth := constants.GridWidth
			gridHeight := constants.GridHeight

			if int(gridX) >= gridWidth || int(gridY) >= gridHeight {
				return color.RGBA{0, 0, 255, 255}
			}

			//val := (*grid)[gridY][gridX]
			//
			//if gridX < gridWidth && gridY < gridHeight {
			//	return color.RGBA{0, val, 0, 255}
			//}
			//
			//return color.RGBA{255, 0, 0, 255}

			//	create a border around the cells
			//	so that the cells don't touch each other

			lineX := int(float32(x) / cellWidth)
			lineY := int(float32(y) / cellHeight)

			if float32(lineX) == gridX || float32(lineY) == gridY {
				return color.RGBA{R: 50, G: 50, B: 50, A: 255}
			}

			val := (*grid)[int(gridY)][int(gridX)]

			if int(gridX) < gridWidth && int(gridY) < gridHeight {
				return color.RGBA{0, val, 0, 255}
			}

			return color.RGBA{255, 0, 0, 255}

		})

	raster.SetMinSize(fyne.NewSize(float32(constants.CanvasWidth), float32(constants.CanvasHeight)))
	raster.ScaleMode = canvas.ImageScaleFastest

	g.Canvas.SetContent(raster)

	//g.Canvas.Refresh(g.Canvas.Content())

	content := container.New(
		layout.NewVBoxLayout(),
		g.Canvas.Content(),
		g.Widgets,
	)

	g.Window.SetContent(content)

}

//func (g *Game) SetContent(content fyne.CanvasObject) {
//	g.Window.SetContent(content)
//}
