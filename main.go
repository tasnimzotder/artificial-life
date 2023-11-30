package main

import (
	app2 "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/tasnimzotder/artificial-life/services"
	"github.com/tasnimzotder/artificial-life/utils"
	"github.com/tasnimzotder/artificial-life/windowManager"
	"image/color"
)

func main() {
	app := app2.New()
	window := app.NewWindow("Artificial Life")

	//app.SendNotification(&fyne.Notification{
	//	Title:   "Artificial Life",
	//	Content: "Welcome to Artificial Life!",
	//})

	gs := &utils.GameSettings{}
	gameWindow := &windowManager.GameWindow{}

	//window.Resize(fyne.NewSize(utils.WindowWidth, utils.WindowHeight))
	window.SetFixedSize(true)
	window.RequestFocus()

	gs.Rows = utils.WindowHeight / utils.TileSize
	gs.Cols = utils.WindowWidth / utils.TileSize
	gs.IsPaused = true
	gs.AliveColor = color.RGBA{R: 0, G: 255, B: 0, A: 0xff}
	gs.DeathColor = color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	gs.FPS = 2
	gs.WrapAround = true
	gs.GameTypes = []string{"GoL", "SmoothLife", "Lenia"}
	gs.GameType = gs.GameTypes[0]
	gs.Presets = map[string][]string{
		"GoL":        {"Random", "Glider", "GliderGun", "Pulsar", "Pentadecathlon"},
		"SmoothLife": {"Random"},
	}
	gs.Preset = gs.Presets[gs.GameType][0]

	grid := container.NewGridWithColumns(gs.Cols)
	tileGrid := make([][]*canvas.Rectangle, gs.Rows)
	gs.TileGrid = &tileGrid

	utils.GenerateGrid(utils.TileSize, gs, grid)
	services.GameRunner(gs)

	windowContent := gameWindow.New(gs)
	content := container.NewVBox(grid, windowContent)

	window.SetContent(content)
	window.ShowAndRun()
}
