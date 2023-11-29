package utils

import (
	"fyne.io/fyne/v2/canvas"
	"image/color"
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
	GameType   string // "GoL", "Lenia", "SmoothLife"
	GameTypes  []string
}
