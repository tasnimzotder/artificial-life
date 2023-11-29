package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	app2 "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tasnimzotder/artificial-life/gameOfLife"
	"github.com/tasnimzotder/artificial-life/smoothLife"
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

	window.Resize(fyne.NewSize(windowWidth, windowHeight))

	gameSettings.Rows = windowHeight / minTileSize
	gameSettings.Cols = windowWidth / minTileSize
	gameSettings.IsPaused = true
	gameSettings.AliveColor = color.RGBA{R: 0, G: 255, B: 0, A: 0xff}
	gameSettings.DeathColor = color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	gameSettings.FPS = 2
	gameSettings.GameTypes = []string{"GoL", "SmoothLife", "Lenia"}
	gameSettings.GameType = gameSettings.GameTypes[0]

	grid := container.NewGridWithColumns(gameSettings.Cols)
	tileGrid := make([][]*canvas.Rectangle, gameSettings.Rows)

	gameSettings.TileGrid = &tileGrid

	fps := widget.NewLabel("0")
	fps.SetText("60")

	//loadGoLButton := widget.NewButton("Load GoL", func() {
	//	// ...
	//})

	startStopButton := widget.NewButton("Start/Pause", func() {
		// ...
	})

	resetButton := widget.NewButton("Reset", func() {
		//	...
	})

	fpsSlider := widget.NewSlider(1, 100)
	fpsSlider.SetValue(2)

	gameTypeSelector := widget.NewSelect(gameSettings.GameTypes, func(s string) {
		gameSettings.GameType = s
	})

	gameTypeSelector.SetSelected("GoL")

	beginButton := widget.NewButton("Begin", func() {
		GameRunner(gameSettings, startStopButton, resetButton, fps, fpsSlider, gameTypeSelector)
	})

	fpsSliderContainer := container.NewVBox(widget.NewLabel("FPS"), fpsSlider)

	// generate grid
	GenerateGrid(minTileSize, gameSettings, grid)

	buttons := container.NewHBox(gameTypeSelector, beginButton, startStopButton, resetButton, fps)
	content := container.NewVBox(grid, fpsSliderContainer, buttons)

	window.SetContent(content)
	window.ShowAndRun()
}

func GameRunner(gs *utils.GameSettings, startStopButton, resetButton *widget.Button, fpsWidget *widget.Label, fpsSlider *widget.Slider, gameSelector *widget.Select) {
	gs.IsReset = true
	HandleReset(gs)
	time.Sleep(500 * time.Millisecond)

	//gs.IsReset = false
	gs.IsPaused = true

	if gs.GameType == gs.GameTypes[0] { // "GoL"
		gameOfLife.GenerateInitialRandomGrid(gs)
	} else if gs.GameType == gs.GameTypes[1] { // "SmoothLife"
		smoothLife.GenerateInitialRandomGridSmoothLife(gs)
	}

	time.Sleep(500 * time.Millisecond)

	go func() {
		for {
			startStopButton.OnTapped = func() {
				fmt.Println("startStopButton.OnTapped")
				gs.IsPaused = !gs.IsPaused

				if !gs.IsPaused && gs.IsReset {
					gs.IsReset = false
				}
			}

			resetButton.OnTapped = func() {
				fmt.Println("resetButton.OnTapped")
				gs.IsReset = true
				gs.FPS = 2
				fpsSlider.SetValue(2)

				HandleReset(gs)
			}

			fpsSlider.OnChanged = func(value float64) {
				fmt.Println("fpsSlider.OnChanged")
				gs.FPS = int(value)
			}

			gameSelector.OnChanged = func(s string) {
				fmt.Printf("gameSelector.OnChanged: %s\n", s)
				gs.GameType = s
			}

			if gs.IsReset {
				break
			}
		}
	}()

	prevTime := time.Now()

	go func() {
		for {
			if !gs.IsPaused {
				if gs.GameType == gs.GameTypes[0] { // "GoL"
					gameOfLife.MovementHandler(gs)
				} else if gs.GameType == gs.GameTypes[1] { // "SmoothLife"
					smoothLife.MovementHandlerSmoothLife(gs)
				}
			}

			elapsed := time.Since(prevTime).Seconds()
			fps := int(1 / elapsed)
			prevTime = time.Now()

			fpsWidget.SetText("FPS: " + strconv.Itoa(fps))

			fpsMillis := 1000 / (gs.FPS)
			time.Sleep(time.Duration(fpsMillis) * time.Millisecond)

			if gs.IsReset {
				break
			}
		}
	}()
}

func HandleReset(gameSettings *utils.GameSettings) {
	//gameOfLife.ClearGoLGrid(*gameSettings.TileGrid, gameSettings)
	ClearGrid(gameSettings)
	gameSettings.IsReset = false
	gameSettings.IsPaused = true
}

func GenerateGrid(tileSize int, gs *utils.GameSettings, grid *fyne.Container) {
	tileGrid := *gs.TileGrid

	for i := 0; i < gs.Rows; i++ {
		tileGrid[i] = make([]*canvas.Rectangle, gs.Cols)

		for j := 0; j < gs.Cols; j++ {
			rect := canvas.NewRectangle(gs.DeathColor)
			rect.SetMinSize(fyne.NewSize(float32(tileSize), float32(tileSize)))

			tileGrid[i][j] = rect
			grid.Add(rect)
		}
	}
}

func ClearGrid(gs *utils.GameSettings) {
	tileGrid := *gs.TileGrid

	for i := 0; i < gs.Rows; i++ {
		for j := 0; j < gs.Cols; j++ {
			tile := tileGrid[i][j]

			tile.FillColor = gs.DeathColor
			tile.Refresh()
		}
	}
}
