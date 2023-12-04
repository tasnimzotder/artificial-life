package settings

import (
	"image/color"
)

type GameStates struct {
	Running bool
	Paused  bool
	Loaded  bool
}

type GameColors struct {
	AliveColor color.RGBA
	DeadColor  color.RGBA
	SeenColor  color.RGBA
}

type GameSettings struct {
	GameType       string
	GameTypes      []string
	Preset         string
	WrapAround     bool
	Colors         GameColors
	States         GameStates
	GPS            uint8 // Generations Per Second
	LogicLoopDelay int64
	ZoomLevel      int

	// Common
	Radius uint8

	//	 GoL
	GoLPresets map[string][]string

	// LtL
	LtLRule string
}

func (gs *GameSettings) NewGameSettings() {
	gs.GameTypes = []string{"GoL", "LtL"}
	gs.GameType = gs.GameTypes[0]

	gs.GoLPresets = map[string][]string{
		"GoL": {"Random", "Glider", "GliderGun", "Pulsar", "Pentadecathlon"},
		"LtL": {"Random"},
	}
	gs.Preset = gs.GoLPresets[gs.GameType][0]

	gs.WrapAround = true

	gs.Colors = GameColors{
		AliveColor: color.RGBA{R: 0, G: 255, B: 0, A: 0xff},
		DeadColor:  color.RGBA{R: 0, G: 0, B: 0, A: 0xff},
		SeenColor:  color.RGBA{R: 0, G: 50, B: 0, A: 0xff},
	}

	gs.States = GameStates{
		Running: false,
		Paused:  false,
		Loaded:  false,
	}

	gs.Radius = 1
	gs.GPS = 2
	//gs.LogicLoopDelay = utils.FPSToMilliseconds(int64(gs.GPS))

	gs.LtLRule = GetLtLRuleNames()[0]

	gs.ZoomLevel = 10
}
