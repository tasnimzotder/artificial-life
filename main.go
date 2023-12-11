package main

import (
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tasnimzotder/artificial-life/settings"
	"github.com/tasnimzotder/artificial-life/ui"
	"log"
)

func main() {
	gameSettings := settings.NewGameSettings()
	uiWidgets := ui.NewUiWidget(gameSettings)

	eui := &ebitenui.UI{
		Container: uiWidgets.RootContainer,
	}

	g := &ui.Game{
		Settings: gameSettings,
		World: settings.NewWorld(
			gameSettings.WorldWidth,
			gameSettings.WorldHeight,
			gameSettings,
		),
		UI: eui,
	}

	ebiten.SetWindowSize(ui.ScreenWidth, ui.ScreenHeight)
	ebiten.SetWindowTitle("Artificial Life")
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(100)
	//ebiten.SetFullscreen(true)

	g.UpdateActualSpeed()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("Ebiten run game error: %v", err)
	}
}
