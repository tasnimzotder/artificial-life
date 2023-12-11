package settings

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tasnimzotder/artificial-life/constants"
	"github.com/tasnimzotder/artificial-life/data"
	"github.com/tasnimzotder/artificial-life/worlds"
	"image/color"
)

type World struct {
	GoLWorld *worlds.GoLWorld
	GS       *GameSettings
}

func NewWorld(width, height int, gs *GameSettings) *World {
	w := &World{
		GoLWorld: worlds.NewGoLWorld(width, height),
		GS:       gs,
	}

	return w
}

func (w *World) InitRandom(gameType string) {
	if gameType == constants.GAME_TYPE_GOL {
		w.GoLWorld.InitRandom()
	}
}

func (w *World) InitPreset(gameType, presetString string) {

	preset := data.GetPreset(gameType, presetString)

	if gameType == constants.GAME_TYPE_GOL {
		w.GoLWorld.InitPreset(preset.Matrix)
	}
}

func (w *World) NextGeneration() {
	if w.GS.GameType == constants.GAME_TYPE_GOL {
		w.GoLWorld.NextGeneration()
	}
}

func (w *World) Draw(pixels *ebiten.Image, visibleWidth, visibleHeight int) {
	if w.GS.GameType == constants.GAME_TYPE_GOL {
		cy := w.GS.WorldHeight / 2
		cx := w.GS.WorldWidth / 2

		for y := 0; y < visibleHeight; y++ {
			for x := 0; x < visibleWidth; x++ {
				if w.GoLWorld.IsAlive((cy-visibleHeight/2)+y, (cx-visibleWidth/2)+x) {
					pixels.Set(x, y, color.RGBA{200, 200, 200, 255})
				} else {
					pixels.Set(x, y, color.RGBA{50, 50, 50, 255})
				}
			}
		}
	}
}
