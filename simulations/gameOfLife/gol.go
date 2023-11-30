package gameOfLife

import (
	"github.com/tasnimzotder/artificial-life/presets"
	"github.com/tasnimzotder/artificial-life/utils"
	"math/rand"
	"sync"
)

func GenerateInitialRandomGrid(gs *utils.GameSettings) {
	rowCenter := gs.Rows / 2
	colCenter := gs.Cols / 2

	// todo: make this a user input
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
				aliveTiles := utils.GetSurroundingAliveTiles(i, j, gameSettings)
				numAliveTiles := len(aliveTiles)

				if tile.FillColor == gameSettings.AliveColor {
					if numAliveTiles < 2 || numAliveTiles > 3 {
						newTileGrid[i][j] = false

					} else if numAliveTiles == 2 || numAliveTiles == 3 {
						newTileGrid[i][j] = true
					}
				} else {
					if numAliveTiles == 3 {
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

func GenerateInitialPresetGrid(gs *utils.GameSettings) {
	preset := presets.GetPreset(gs.GameType, gs.Preset)

	presetRows, presetCols := len(preset), len(preset[0])

	rowCenter := gs.Rows/2 - presetRows/2
	colCenter := gs.Cols/2 - presetCols/2

	for i := 0; i < presetRows; i++ {
		for j := 0; j < presetCols; j++ {
			tile := (*gs.TileGrid)[rowCenter+i][colCenter+j]

			if preset[i][j] == 1 {
				go func() {
					tile.FillColor = gs.AliveColor
					tile.Refresh()
				}()
			}
		}
	}
}
