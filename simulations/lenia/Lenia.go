package lenia

import (
	"github.com/tasnimzotder/artificial-life/utils"
	"math/rand"
)

func GenerateInitialRandomGridLenia(gs *utils.GameSettings) {
	rowCenter := gs.Rows / 2
	colCenter := gs.Cols / 2
	buffer := 10

	for i := rowCenter - buffer; i < rowCenter+buffer; i++ {
		for j := colCenter - buffer; j < colCenter+buffer; j++ {
			tile := (*gs.TileGrid)[i][j]

			if rand.Int()%2 == 0 {
				continue
			}

			go func() {
				tile.FillColor = gs.AliveColor
				tile.Refresh()
			}()
		}
	}
}
