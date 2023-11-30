package utils

import (
	"image/color"
)

func GetSurroundingAliveTiles(row int, col int, gs *GameSettings) []color.Color {
	var aliveTiles []color.Color

	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			// if out of bounds, wrap around
			ix, jx := i, j

			if gs.WrapAround {
				if i < 0 {
					ix = gs.Rows - 1
				} else if i >= gs.Rows {
					ix = 0
				}

				if j < 0 {
					jx = gs.Cols - 1
				} else if j >= gs.Cols {
					jx = 0
				}
			} else {
				if i < 0 || i >= gs.Rows || j < 0 || j >= gs.Cols {
					continue
				}
			}

			if ix == row && jx == col {
				continue
			}

			//if ix < 0 || ix >= gs.Rows || jx < 0 || jx >= gs.Cols {
			//	log.Fatalf(">> R: %d, C: %d, i: %d, j: %d, ix: %d, jx: %d\n", gs.Rows, gs.Cols, i, j, ix, jx)
			//}

			tile := (*gs.TileGrid)[ix][jx]

			if tile.FillColor != gs.DeathColor {
				aliveTiles = append(aliveTiles, tile.FillColor)
			} else {
				continue
			}
		}
	}

	return aliveTiles
}
