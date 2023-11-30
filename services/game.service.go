package services

import (
	"github.com/tasnimzotder/artificial-life/simulations/gameOfLife"
	"github.com/tasnimzotder/artificial-life/simulations/smoothLife"
	"github.com/tasnimzotder/artificial-life/utils"
	"time"
)

func GameRunner(
	gs *utils.GameSettings,
) {
	gs.IsReset = true
	utils.HandleReset(gs)
	time.Sleep(500 * time.Millisecond)

	//gs.IsReset = false
	gs.IsPaused = true

	if gs.GameType == gs.GameTypes[0] { // "GoL"
		if gs.Preset == "Random" {
			gameOfLife.GenerateInitialRandomGrid(gs)
		} else {
			gameOfLife.GenerateInitialPresetGrid(gs)
		}
	} else if gs.GameType == gs.GameTypes[1] { // "SmoothLife"
		smoothLife.GenerateInitialRandomGridSmoothLife(gs)
	}

	time.Sleep(500 * time.Millisecond)

	prevTime := time.Now()

	go func() {
		for {
			prevStartTime := time.Now()

			if !gs.IsPaused {
				// update the T0 value
				gs.Parameters.T += 1
			
				if gs.GameType == gs.GameTypes[0] { // "GoL"
					gameOfLife.MovementHandler(gs)
				} else if gs.GameType == gs.GameTypes[1] { // "SmoothLife"
					smoothLife.MovementHandlerSmoothLife(gs)
				}
			}

			currTime := time.Now()
			elapsed := currTime.Sub(prevTime).Seconds()
			fps := 1.0 / elapsed
			prevTime = currTime

			if gs.IsPaused {
				fps = 0
			}

			//fpsWidget.SetText("FPS: " + strconv.Itoa(fps))
			gs.CurrentFPS = fps

			fpsMillis := 1000 / (gs.FPS)
			loopTime := time.Now().Sub(prevStartTime).Milliseconds()

			time.Sleep(time.Duration(fpsMillis)*time.Millisecond - time.Duration(loopTime)*time.Millisecond)

			if gs.IsReset {
				break
			}
		}
	}()
}
