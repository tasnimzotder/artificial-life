package gameOfLife

import (
	"fyne.io/fyne/v2/canvas"
	"github.com/tasnimzotder/artificial-life/utils"
	"image/color"
)

func ClearGoLGrid(tileGrid [][]*canvas.Rectangle, gameSettings *utils.GameSettings) {
	for i := 0; i < gameSettings.Rows; i++ {
		for j := 0; j < gameSettings.Cols; j++ {
			tile := tileGrid[i][j]
			tile.FillColor = color.RGBA{R: 150, G: 122, B: 200, A: 0xff}

			tile.Refresh()
		}
	}
}
