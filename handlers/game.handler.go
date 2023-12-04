package handlers

import (
	"github.com/tasnimzotder/artificial-life/constants"
	"github.com/tasnimzotder/artificial-life/settings"
	"github.com/tasnimzotder/artificial-life/ui"
	"github.com/tasnimzotder/artificial-life/utils"
	"time"
)

func GameLoopHandler(g *ui.Game, gs *settings.GameSettings, gd *settings.GameData) {
	uiUpdateDelay := time.Second.Milliseconds() / constants.TickRate
	//logicLoopDelay := time.Second.Milliseconds() / int64(gs.GPS)
	gs.LogicLoopDelay = utils.FPSToMilliseconds(int64(gs.GPS))

	prevUIUpdateTime := time.Now()
	prevLogicUpdateTime := time.Now()

	h := constants.GridHeight / 2
	w := constants.GridWidth / 2

	// initial settings
	gs.States.Paused = true
	gs.GameType = constants.GameTypeGof

	// todo: fix redundancy
	utils.GenerateRandomGrid(gd.Grid, h, w, gs.GameType)
	g.Update(gd.Grid)
	gs.States.Loaded = true

	for range time.Tick(time.Millisecond * 10) {
		currTime := time.Now()

		// logic loop
		if time.Since(prevLogicUpdateTime) > time.Duration(gs.LogicLoopDelay)*time.Millisecond {
			prevLogicUpdateTime = currTime

			if gs.States.Loaded == false {
				if gs.GameType == constants.GameTypeGof {
					if gs.Preset == constants.RANDOM {
						utils.GenerateRandomGrid(gd.Grid, h, w, gs.GameType)
					} else {
						utils.GeneratePresetGrid(gd.Grid, gs.Preset, gs.GameType)
					}
				}

				g.Update(gd.Grid)

				gs.States.Paused = true
				gs.States.Loaded = true

			}

			if !gs.States.Paused {
				//g.Update(gd.Grid)

				utils.NextGenerationGrid(gd.Grid, gs)

			}
		}

		// ui loop
		if time.Since(prevUIUpdateTime) > time.Duration(uiUpdateDelay)*time.Millisecond {
			if !gs.States.Paused {
				prevUIUpdateTime = currTime

				//utils.NextGenerationGrid(gd.Grid, gs)
				var newGrid [][]uint8
				for _, row := range *gd.Grid {
					var newRow []uint8
					for _, cell := range row {
						newRow = append(newRow, cell)
					}
					newGrid = append(newGrid, newRow)
				}

				g.Update(&newGrid)
			}
		}

	}
}
