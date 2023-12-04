package utils

import (
	"github.com/tasnimzotder/artificial-life/settings"
)

func GenerateRandomGrid(grid *[][]uint8, height, width int, gameType string) {
	var newGrid [][]uint8

	if gameType != "GoL" {
		// todo: implement
	}

	for i := 0; i < height; i++ {
		newGrid = append(newGrid, []uint8{})

		for j := 0; j < width; j++ {
			randomValue := uint8(RandomInt(0, 1))
			G := uint8(0)

			if randomValue == 1 {
				G = uint8(255)
			}

			newGrid[i] = append(newGrid[i], G)
		}
	}

	*grid = newGrid
}

func Neighbors(grid *[][]uint8, x, y int, gs *settings.GameSettings) (neighbors []uint8, aliveCount int) {
	neighbors = []uint8{}
	aliveCount = 0
	R := int(gs.Radius)

	for k := -R; k <= R; k++ {
		for l := -R; l <= R; l++ {
			if !(k == 0 && l == 0) {
				cellValue := uint8(0)

				if gs.WrapAround {
					cellValue = (*grid)[(x+k+len(*grid))%len(*grid)][(y+l+len((*grid)[x]))%len((*grid)[x])]
					neighbors = append(neighbors, cellValue)
				} else {
					if x+k >= 0 && x+k < len(*grid) && y+l >= 0 && y+l < len((*grid)[x]) {
						cellValue = (*grid)[x+k][y+l]
						neighbors = append(neighbors, cellValue)
					}
				}

				if cellValue == 255 {
					aliveCount += 1
				}
			}
		}
	}

	return neighbors, aliveCount
}

func GeneratePresetGrid(grid *[][]uint8, preset string, gameType string) {
	var newGrid [][]uint8
	presetData := settings.GetPreset(gameType, preset)

	presetRows, presetCols := len(presetData), len(presetData[0])

	cRow := len(*grid)/2 - presetRows/2
	cCol := len((*grid)[0])/2 - presetCols/2

	for i := 0; i < len(*grid); i++ {
		newGrid = append(newGrid, []uint8{})

		for j := 0; j < len((*grid)[i]); j++ {
			if i >= cRow && i < cRow+presetRows && j >= cCol && j < cCol+presetCols {
				if presetData[i-cRow][j-cCol] == 1 {
					newGrid[i] = append(newGrid[i], uint8(255))
				} else {
					newGrid[i] = append(newGrid[i], uint8(0))
				}
			} else {
				newGrid[i] = append(newGrid[i], uint8(0))
			}
		}
	}

	*grid = newGrid
}

func LtLLiveDeathLimit(gameType, rule string) (Radius int, S, B settings.LimitMinMax) {
	Radius, S, B = 0, settings.LimitMinMax{}, settings.LimitMinMax{}

	if gameType == "GoL" {
		S = settings.LimitMinMax{
			Min: 2,
			Max: 3,
		}

		B = settings.LimitMinMax{
			Min: 3,
			Max: 3,
		}

		Radius = 1
	} else if gameType == "LtL" {
		rule := settings.GetLtLRule(rule)

		S = rule.S
		B = rule.B
		Radius = rule.Radius
	}

	return Radius, S, B
}
