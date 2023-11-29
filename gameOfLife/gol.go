package gameOfLife

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/tasnimzotder/artificial-life/utils"
	"image/color"
	"strconv"
	"time"
)

func HandleGOL(gameSettings *utils.GameSettings, fpsWidget *widget.Label) {
	gameSettings.IsReset = false

	GenerateInitialRandomGrid(gameSettings)

	time.Sleep(500 * time.Millisecond)

	desiredFPS := 2
	prevTime := time.Now()

	for range time.Tick(1 * time.Microsecond) {
		currentTime := time.Now()

		if gameSettings.IsReset {
			break
		}

		if currentTime.Sub(prevTime).Seconds() > 1/float64(desiredFPS) {
			MovementHandler(gameSettings)

			elapsed := time.Since(prevTime).Seconds()
			fps := int(1 / elapsed)

			fpsWidget.SetText("FPS: " + strconv.Itoa(fps))
			prevTime = currentTime
		}
	}
}

func clearGrid(tileGrid [][]*canvas.Rectangle, gameSettings *utils.GameSettings) {
	for i := 0; i < gameSettings.Rows; i++ {
		for j := 0; j < gameSettings.Cols; j++ {
			tile := tileGrid[i][j]
			tile.FillColor = color.RGBA{R: 150, G: 122, B: 200, A: 0xff}

			tile.Refresh()
		}
	}
}

func HandleReset(tileGrid [][]*canvas.Rectangle, gameSettings *utils.GameSettings) {
	clearGrid(tileGrid, gameSettings)
}
