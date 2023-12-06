package main

import (
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tasnimzotder/artificial-life/settings"
	"github.com/tasnimzotder/artificial-life/ui2"
	"log"
)

func main() {
	gameSettings := settings.NewGameSettings()
	uiWidgets := ui2.NewUiWidget(gameSettings)

	eui := &ebitenui.UI{
		Container: uiWidgets.RootContainer,
	}

	g := &ui2.Game{
		Settings: gameSettings,
		World: settings.NewWorld(
			gameSettings.WorldWidth,
			gameSettings.WorldHeight,
			(gameSettings.WorldWidth*gameSettings.WorldHeight)/2,
		),
		UI: eui,
	}

	ebiten.SetWindowSize(ui2.ScreenWidth, ui2.ScreenHeight)
	ebiten.SetWindowTitle("Artificial Life")
	ebiten.SetVsyncEnabled(true)
	//ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("Ebiten run game error: %v", err)
	}
}
