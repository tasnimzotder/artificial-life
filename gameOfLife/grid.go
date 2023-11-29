package gameOfLife

import (
	"github.com/tasnimzotder/artificial-life/utils"
	"math/rand"
	"sync"
)

func GenerateInitialRandomGrid(gs *utils.GameSettings) {
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

func GetSurroundingAliveTiles(row int, col int, gs *utils.GameSettings) int {
	var aliveTiles int

	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i == row && j == col {
				continue
			}

			if i < 0 || i >= gs.Rows || j < 0 || j >= gs.Cols {
				continue
			}

			tile := (*gs.TileGrid)[i][j]

			if tile.FillColor == gs.AliveColor {
				aliveTiles++
			} else {
				continue
			}
		}
	}

	return aliveTiles
}

func MovementHandler(gameSettings *utils.GameSettings) {
	newTileGrid := make([][]bool, gameSettings.Rows)
	var wg sync.WaitGroup

	//for range time.Tick(1000 * time.Millisecond) {
	for i := 0; i < gameSettings.Rows; i++ {
		newTileGrid[i] = make([]bool, gameSettings.Cols)

		for j := 0; j < gameSettings.Cols; j++ {
			wg.Add(1)

			go func(i, j int) {
				wg.Done()

				tile := (*gameSettings.TileGrid)[i][j]
				aliveTiles := GetSurroundingAliveTiles(i, j, gameSettings)

				if tile.FillColor == gameSettings.AliveColor {
					if aliveTiles < 2 || aliveTiles > 3 {
						newTileGrid[i][j] = false

					} else if aliveTiles == 2 || aliveTiles == 3 {
						newTileGrid[i][j] = true
					}
				} else {
					if aliveTiles == 3 {
						newTileGrid[i][j] = true
					} else {
						newTileGrid[i][j] = false
					}
				}
			}(i, j)
		}
	}

	wg.Wait()

	for i := 0; i < gameSettings.Rows; i++ {
		for j := 0; j < gameSettings.Cols; j++ {

			wg.Add(1)

			go func(i, j int) {
				wg.Done()

				tile := (*gameSettings.TileGrid)[i][j]

				if newTileGrid[i][j] {
					tile.FillColor = gameSettings.AliveColor
				} else {
					tile.FillColor = gameSettings.DeathColor
				}

				tile.Refresh()
			}(i, j)
		}
	}

	wg.Wait()
}
