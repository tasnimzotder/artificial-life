package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func HandleReset(gs *GameSettings) {
	//gameOfLife.ClearGoLGrid(*gs.TileGrid, gs)
	ClearGrid(gs)
	gs.IsReset = false
	gs.IsPaused = true
}

func GenerateGrid(tileSize int, gs *GameSettings, grid *fyne.Container) {
	tileGrid := *gs.TileGrid

	for i := 0; i < gs.Rows; i++ {
		tileGrid[i] = make([]*canvas.Rectangle, gs.Cols)

		for j := 0; j < gs.Cols; j++ {
			rect := canvas.NewRectangle(gs.DeathColor)
			rect.SetMinSize(fyne.NewSize(float32(tileSize), float32(tileSize)))
			rect.CornerRadius = TileCornerRadius

			tileGrid[i][j] = rect
			grid.Add(rect)
		}
	}
}

func ClearGrid(gs *GameSettings) {
	tileGrid := *gs.TileGrid

	for i := 0; i < gs.Rows; i++ {
		for j := 0; j < gs.Cols; j++ {
			tile := tileGrid[i][j]

			tile.FillColor = gs.DeathColor
			tile.Refresh()
		}
	}
}
