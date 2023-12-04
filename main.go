package main

import (
	"github.com/tasnimzotder/artificial-life/handlers"
	"github.com/tasnimzotder/artificial-life/settings"
	"github.com/tasnimzotder/artificial-life/ui"
	"time"
)

func main() {
	g := &ui.Game{}
	g.NewGame()

	gs := &settings.GameSettings{}
	gs.NewGameSettings()

	gd := &settings.GameData{}
	gd.NewGameData()

	uic := &ui.UIController{}
	uic.NewUIController(gs)

	time.Sleep(time.Second)

	// main loop
	go func() {
		handlers.GameLoopHandler(g, gs, gd)
	}()

	g.Widgets = uic.Content
	g.Run()

}
