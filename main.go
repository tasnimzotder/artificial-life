package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	app2 "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tasnimzotder/artificial-life/gameOfLife"
	"github.com/tasnimzotder/artificial-life/utils"
	"image/color"
	"strconv"
	"time"
)

func main() {
	app := app2.New()
	window := app.NewWindow("Artificial Life")

	gameSettings := &utils.GameSettings{}

	const (
		minTileSize  = 10
		windowWidth  = 16 * 60
		windowHeight = 9 * 60
	)

	// set window size
	window.Resize(fyne.NewSize(windowWidth, windowHeight))

	//windowSize := window.Canvas().Size()

	//fmt.Printf("window size: %v\n", windowSize)

	gameSettings.Rows = windowHeight / minTileSize
	gameSettings.Cols = windowWidth / minTileSize
	gameSettings.IsPaused = true
	gameSettings.AliveColor = color.RGBA{R: 0, G: 255, B: 0, A: 0xff}
	gameSettings.DeathColor = color.RGBA{R: 255, G: 0, B: 0, A: 0xff}

	grid := container.NewGridWithColumns(gameSettings.Cols)
	tileGrid := make([][]*canvas.Rectangle, gameSettings.Rows)

	for i := 0; i < gameSettings.Rows; i++ {
		tileGrid[i] = make([]*canvas.Rectangle, gameSettings.Cols)

		for j := 0; j < gameSettings.Cols; j++ {
			rect := canvas.NewRectangle(color.RGBA{R: 150, G: 122, B: 200, A: 0xff})
			rect.SetMinSize(fyne.NewSize(minTileSize, minTileSize))

			tileGrid[i][j] = rect
			grid.Add(rect)
		}
	}

	gameSettings.TileGrid = &tileGrid

	//for i := 0; i < rows; i++ {
	//	for j := 0; j < cols; j++ {
	//		rect := canvas.NewRectangle(color.RGBA{R: 150, G: 122, B: 200, A: 0xff})
	//		rect.SetMinSize(fyne.NewSize(10, 10))
	//
	//		grid.Add(rect)
	//	}
	//}

	fps := widget.NewLabel("0")
	fps.SetText("60")

	loadGoLButton := widget.NewButton("Load GoL", func() {
		//gameOfLife.HandleGOL(gameSettings, fps)
	})

	startStopButton := widget.NewButton("Start/Pause", func() {
		//gameSettings.IsPaused = !gameSettings.IsPaused
	})

	resetButton := widget.NewButton("Reset", func() {
		//gameSettings.IsReset = true
		//
		//time.Sleep(2 * time.Second)
		//
		//gameOfLife.HandleReset(tileGrid, gameSettings)
	})

	beginButton := widget.NewButton("Begin", func() {
		GameRunner(gameSettings, startStopButton, resetButton, fps)
	})

	buttons := container.NewHBox(beginButton, loadGoLButton, startStopButton, resetButton, fps)
	content := container.NewVBox(grid, buttons)

	window.SetContent(content)
	window.ShowAndRun()

}

func GameRunner(gameSettings *utils.GameSettings, startStopButton, resetButton *widget.Button, fpsWidget *widget.Label) {
	gameSettings.IsReset = false
	gameSettings.IsPaused = true

	gameOfLife.GenerateInitialRandomGrid(gameSettings)
	time.Sleep(500 * time.Millisecond)

	go func() {
		for {
			startStopButton.OnTapped = func() {
				fmt.Println("startStopButton.OnTapped")
				gameSettings.IsPaused = !gameSettings.IsPaused
			}

			resetButton.OnTapped = func() {
				fmt.Println("resetButton.OnTapped")
				gameSettings.IsReset = true
				HandleReset(gameSettings)
			}
		}
	}()

	go func() {
		for range time.Tick(1 * time.Second) {
			prevTime := time.Now()

			if !gameSettings.IsPaused {
				gameOfLife.MovementHandler(gameSettings)
			}

			elapsed := time.Since(prevTime).Seconds()
			fps := int(1 / elapsed)

			fpsWidget.SetText("FPS: " + strconv.Itoa(fps))

			fmt.Println("running")
		}
	}()
}

func HandleReset(gameSettings *utils.GameSettings) {
	gameOfLife.ClearGoLGrid(*gameSettings.TileGrid, gameSettings)

	gameSettings.IsReset = false
	gameSettings.IsPaused = true
}
