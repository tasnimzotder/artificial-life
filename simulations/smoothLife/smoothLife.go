package smoothLife

import (
	"github.com/tasnimzotder/artificial-life/utils"
	"image/color"
	"math/rand"
	"sync"
)

//func SmoothLife() {
//
//}

func GenerateInitialRandomGridSmoothLife(gs *utils.GameSettings) {
	rowCenter := gs.Rows / 2
	colCenter := gs.Cols / 2
	buffer := 10

	for i := rowCenter - buffer; i < rowCenter+buffer; i++ {
		for j := colCenter - buffer; j < colCenter+buffer; j++ {
			tile := (*gs.TileGrid)[i][j]

			if rand.Int()%2 == 0 {
				continue
			}

			red := 0
			green := uint8(rand.Intn(255))
			blue := 0

			go func() {
				tile.FillColor = color.RGBA{R: uint8(red), G: green, B: uint8(blue), A: 0xff}
				tile.Refresh()
			}()
		}
	}

}

func MovementHandlerSmoothLife(gs *utils.GameSettings) {
	var wg sync.WaitGroup

	// for smooth life
	newTileGrid := make([][]color.Color, gs.Rows)

	for i := 0; i < gs.Rows; i++ {
		newTileGrid[i] = make([]color.Color, gs.Cols)

		for j := 0; j < gs.Cols; j++ {
			wg.Add(1)

			go func(i int, j int) {
				defer wg.Done()

				tile := (*gs.TileGrid)[i][j]
				newValue := 0.0

				aliveTiles := utils.GetSurroundingAliveTiles(i, j, gs)
				numAliveTiles := len(aliveTiles)
				avgValue := 0.0

				for _, tile := range aliveTiles {
					avgValue += float64(tile.(color.RGBA).G)
				}

				avgValue /= float64(numAliveTiles)
				avgValue /= 255.0

				currentTileValue := float64(tile.FillColor.(color.RGBA).G) / 255.0

				if currentTileValue >= 0.3 && currentTileValue <= 0.7 {
					newValue = currentTileValue + 0.05
				} else if currentTileValue < 0.3 || currentTileValue > 0.7 {
					newValue = 0
				} else if currentTileValue < 0.1 && avgValue > 0.5 {
					newValue = 0.1
				} else if currentTileValue > 0.9 && avgValue < 0.5 {
					newValue = 0.9
				} else {
					newValue = currentTileValue
				}

				red := 0
				green := uint8(newValue * 255)
				blue := 0

				newTileGrid[i][j] = color.RGBA{R: uint8(red), G: green, B: uint8(blue), A: 0xff}
			}(i, j)
		}
	}

	wg.Wait()

	for i := 0; i < gs.Rows; i++ {
		for j := 0; j < gs.Cols; j++ {
			wg.Add(1)

			go func(i int, j int) {
				defer wg.Done()

				tile := (*gs.TileGrid)[i][j]

				if newTileGrid[i][j] != gs.DeathColor {
					tile.FillColor = newTileGrid[i][j]
				} else {
					tile.FillColor = gs.DeathColor
				}

				tile.Refresh()
			}(i, j)
		}
	}

	wg.Wait()
}
