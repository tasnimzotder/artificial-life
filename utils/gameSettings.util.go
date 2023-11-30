package utils

import (
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

const (
	TileSize         = 5
	WindowWidth      = 16 * 60 / (1.5 * 5 / TileSize)
	WindowHeight     = 9 * 60 / (1.5 * 5 / TileSize)
	TileCornerRadius = TileSize / 10.0
)

type GameSettings struct {
	Rows       int
	Cols       int
	IsPaused   bool
	IsReset    bool
	TileGrid   *[][]*canvas.Rectangle
	AliveColor color.RGBA
	DeathColor color.RGBA
	FPS        int
	CurrentFPS int
	GameType   string // "GoL", "Lenia", "SmoothLife"
	GameTypes  []string
	Preset     string // "Random", "Glider", "GliderGun", "Pulsar", "Pentadecathlon"
	Presets    map[string][]string
	WrapAround bool
}
