package constants

import "image/color"

const (
	GameTypeGol        = "Game of Life"
	GameTypeLtl        = "Larger than Life"
	GameTypeSmoothLife = "Smooth Life"
)

// colors
var (
	// BGColor background
	BGColor = color.RGBA{R: 50, G: 60, B: 70, A: 255}
	// AliveCellColor alive cell
	AliveCellColor = color.RGBA{R: 239, G: 254, B: 238, A: 255}
)
