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

	gameSettings := &utils.GameSettings{}

	//window.Resize(fyne.NewSize(utils.WindowWidth, utils.WindowHeight))
	window.SetFixedSize(true)
	window.RequestFocus()

	gameSettings.Rows = utils.WindowHeight / utils.TileSize
	gameSettings.Cols = utils.WindowWidth / utils.TileSize
	gameSettings.IsPaused = true
	gameSettings.AliveColor = color.RGBA{R: 0, G: 255, B: 0, A: 0xff}
	gameSettings.DeathColor = color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	gameSettings.FPS = 2
	gameSettings.WrapAround = true
	gameSettings.GameTypes = []string{"GoL", "SmoothLife", "Lenia"}
	gameSettings.GameType = gameSettings.GameTypes[0]
	gameSettings.Presets = map[string][]string{
		"GoL":        {"Random", "Glider", "GliderGun", "Pulsar", "Pentadecathlon"},
		"SmoothLife": {"Random"},
	}
	gameSettings.Preset = "Glider"

	grid := container.NewGridWithColumns(gameSettings.Cols)
	tileGrid := make([][]*canvas.Rectangle, gameSettings.Rows)
	gameSettings.TileGrid = &tileGrid
	
	utils.GenerateGrid(utils.TileSize, gameSettings, grid)
	services.GameRunner(gameSettings)

	content := container.NewVBox(grid, windowManager.GameWindow(gameSettings))

	window.SetContent(content)
	window.ShowAndRun()
}
