package utils

import (
	"github.com/tasnimzotder/artificial-life/settings"
)

func NextGenerationGrid(grid *[][]uint8, gs *settings.GameSettings) {
	var newGrid [][]uint8

	for i := 0; i < len(*grid); i++ {
		newGrid = append(newGrid, []uint8{})

		for j := 0; j < len((*grid)[i]); j++ {
			Radius, S, B := LtLLiveDeathLimit(gs.GameType, gs.LtLRule)
			gs.Radius = uint8(Radius)
			_, aliveCount := Neighbors(grid, i, j, gs)

			if (*grid)[i][j] == 255 {
				if aliveCount < S.Min || aliveCount > S.Max {
					//newGrid[i] = append(newGrid[i], uint8(0))
					newGrid[i] = append(newGrid[i], uint8(50))
				} else {
					newGrid[i] = append(newGrid[i], uint8(255))
				}
			} else {
				if aliveCount >= B.Min && aliveCount <= B.Max {
					newGrid[i] = append(newGrid[i], uint8(255))
				} else {
					if (*grid)[i][j] > 0 {
						newGrid[i] = append(newGrid[i], uint8(50))
					} else {
						newGrid[i] = append(newGrid[i], uint8(0))
					}

				}
			}
		}
	}

	*grid = newGrid
}
