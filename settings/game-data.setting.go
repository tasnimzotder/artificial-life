package settings

import (
	"github.com/tasnimzotder/artificial-life/constants"
)

type GameData struct {
	Grid *[][]uint8
}

func (gd *GameData) NewGameData() {
	gd.Grid = &[][]uint8{}

	//totalCells := constants.GRID_HEIGHT * constants.GRID_WIDTH

	for i := 0; i < constants.GridHeight; i++ {
		*gd.Grid = append(*gd.Grid, []uint8{})

		for j := 0; j < constants.GridWidth; j++ {
			//R := rand.Int() % 255
			R := 0

			(*gd.Grid)[i] = append((*gd.Grid)[i], uint8(R))
		}
	}
}

func (gd *GameData) UpdateGrid(grid [][]uint8) {
	gd.Grid = &grid
}
