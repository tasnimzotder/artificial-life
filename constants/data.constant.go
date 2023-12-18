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
	BGColor = color.RGBA{R: 10, G: 20, B: 30, A: 255}
	// AliveCellColor alive cell
	AliveCellColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
)
