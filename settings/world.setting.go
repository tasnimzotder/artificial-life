package settings

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tasnimzotder/artificial-life/constants"
	"github.com/tasnimzotder/artificial-life/data"
	"github.com/tasnimzotder/artificial-life/worlds"
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
	if gameType == constants.GameTypeGol {
		w.GoLWorld.InitRandom()
	}
}

func (w *World) InitPreset(gameType, presetString string) {

	preset := data.GetPreset(gameType, presetString)

	if gameType == constants.GameTypeGol {
		w.GoLWorld.InitPreset(preset.Matrix)
	}
}

func (w *World) NextGeneration() {
	if w.GS.GameType == constants.GameTypeGol {
		w.GoLWorld.NextGeneration()
	}
}

func (w *World) Draw(pixels *ebiten.Image, visibleWidth, visibleHeight int) {
	if w.GS.GameType == constants.GameTypeGol {
		cy := w.GS.WorldHeight / 2
		cx := w.GS.WorldWidth / 2

		newArea := w.GoLWorld.GetArea()

		for y := 0; y < visibleHeight; y++ {
			for x := 0; x < visibleWidth; x++ {
				//if w.GoLWorld.IsAlive((cy-visibleHeight/2)+y, (cx-visibleWidth/2)+x) {
				//	pixels.Set(x, y, color.RGBA{200, 200, 200, 255})
				//} else {
				//	pixels.Set(x, y, color.RGBA{50, 50, 50, 255})
				//}

				ny := cy - visibleHeight/2 + y
				nx := cx - visibleWidth/2 + x

				idx := ny*w.GS.WorldWidth + nx

				if newArea[idx]&0x01 == 0x01 {
					pixels.Set(x, y, constants.AliveCellColor)
				} else {
					pixels.Set(x, y, constants.BGColor)
				}
			}
		}
	}
}
