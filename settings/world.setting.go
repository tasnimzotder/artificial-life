package settings

import (
	"github.com/tasnimzotder/artificial-life/constants"
	"math/rand"
)

func RandomColor0n255() uint8 {
	r := rand.Intn(2)

	if r == 0 {
		return 0
	}

	return 255
}

type World struct {
	Area   [][]byte
	Width  int
	Height int
}

func NewWorld(width, height, maxInitialLiveCells int) *World {
	area := make([][]byte, height)

	for i := range area {
		area[i] = make([]byte, width)
	}

	w := &World{
		Area:   area,
		Width:  width,
		Height: height,
	}

	//w.initRandom(maxInitialLiveCells)l

	return w
}

func (w *World) InitRandom(maxInitialLiveCells int) {
	for n := 0; n < maxInitialLiveCells; n++ {
		x := rand.Intn(w.Width)
		y := rand.Intn(w.Height)

		w.SetCell(x, y, RandomColor0n255())
	}
}

func (w *World) InitPreset(gameType string, presetString string) {
	preset := GetPreset(gameType, presetString)

	presetRows, presetCols := len(preset.Matrix), len(preset.Matrix[0])

	cRow := w.Height/2 - presetRows/2
	cCol := w.Width/2 - presetCols/2

	for i := 0; i < w.Height; i++ {
		w.Area[i] = make([]byte, w.Width)

		for j := 0; j < w.Width; j++ {
			if i >= cRow && i < cRow+presetRows && j >= cCol && j < cCol+presetCols {
				cell := preset.Matrix[i-cRow][j-cCol]

				if cell == 1 {
					w.Area[i][j] = 255
				} else {
					w.Area[i][j] = 0
				}
			} else {
				w.Area[i][j] = 0
			}
		}
	}
}

func (w *World) ClearCells() {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.SetCell(x, y, 0)
		}
	}
}

func (w *World) GetCell(x, y int) byte {
	return w.Area[y][x]
}

func (w *World) SetCell(x, y int, value byte) {
	w.Area[y][x] = value
}

func (w *World) GetNeighbours(Radius uint8, x, y int) (neighbours []byte, aliveCount int) {
	neighbours = make([]byte, 0)
	aliveCount = 0
	R := int(Radius)

	for r := -R; r <= R; r++ {
		for c := -R; c <= R; c++ {
			if !(r == 0 && c == 0) {
				cellValue := byte(0)

				cellValue = w.GetCell((x+r+w.Width)%w.Width, (y+c+w.Height)%w.Height)
				neighbours = append(neighbours, cellValue)

				if cellValue == 255 {
					aliveCount += 1
				}
			}
		}
	}

	//for i := -1; i < 2; i++ {
	//	for j := -1; j < 2; j++ {
	//		if i == 0 && j == 0 {
	//			continue
	//		}
	//
	//		//if x+i < 0 || x+i >= w.Width || y+j < 0 || y+j >= w.Height {
	//		//	continue
	//		//}
	//
	//		// wrap around
	//		if x+i < 0 {
	//			x = w.Width - 1
	//		} else if x+i >= w.Width {
	//			x = 0
	//		}
	//
	//		if y+j < 0 {
	//			y = w.Height - 1
	//		} else if y+j >= w.Height {
	//			y = 0
	//		}
	//
	//		neighbours = append(neighbours, w.GetCell(x+i, y+j))
	//
	//		if w.GetCell(x+i, y+j) != 0 {
	//			aliveCount++
	//		}
	//
	//	}
	//}

	return neighbours, aliveCount
}

func (w *World) NextGeneration(gs *GameSettings2) {
	width := w.Width
	height := w.Height

	nextGeneration := make([][]byte, height)

	for y := 0; y < height; y++ {
		nextGeneration[y] = make([]byte, width)

		for x := 0; x < width; x++ {
			Radius, S, B := LtLLiveDeathLimit(gs.GameType, gs.LtLRule)
			_, aliveCount := w.GetNeighbours(uint8(Radius), x, y)

			//log.Printf("Radius: %d", Radius)
			//log.Printf("GameType: %s", gs.GameType)
			//log.Printf("LtLRule: %s", gs.LtLRule)

			//break

			//log.Printf("aliveCount: %d", aliveCount)

			if w.GetCell(x, y) == 255 {
				if aliveCount < S.Min || aliveCount > S.Max {
					nextGeneration[y][x] = 0
				} else {
					nextGeneration[y][x] = 255
				}
			} else {
				if aliveCount >= B.Min && aliveCount <= B.Max {
					nextGeneration[y][x] = 255
				} else {
					nextGeneration[y][x] = 0
				}
			}

			//if w.GetCell(x, y) != 0 {
			//	if aliveCount < 2 || aliveCount > 3 {
			//		nextGeneration[y][x] = 0
			//	} else {
			//		nextGeneration[y][x] = 255
			//	}
			//} else {
			//	if aliveCount == 3 {
			//		nextGeneration[y][x] = 255
			//	} else {
			//		nextGeneration[y][x] = 0
			//	}
			//}
		}
	}

	w.Area = nextGeneration
}

func (w *World) Draw(pixels []byte) {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			if w.GetCell(x, y) != 0 {
				pixels[(y*w.Width+x)*4] = 100
				pixels[(y*w.Width+x)*4+1] = 130
				pixels[(y*w.Width+x)*4+2] = 160
				pixels[(y*w.Width+x)*4+3] = 255
			} else {
				pixels[(y*w.Width+x)*4] = 25
				pixels[(y*w.Width+x)*4+1] = 25
				pixels[(y*w.Width+x)*4+2] = 25
				pixels[(y*w.Width+x)*4+3] = 255
			}
		}
	}
	//for i, v := range w.Area {
	//	if v != 0 {
	//		pixels[i*4] = 100
	//		pixels[i*4+1] = 150
	//		pixels[i*4+2] = 160
	//		pixels[i*4+3] = 255
	//	} else {
	//		pixels[i*4] = 50
	//		pixels[i*4+1] = 40
	//		pixels[i*4+2] = 30
	//		pixels[i*4+3] = 255
	//	}
	//}
}

func LtLLiveDeathLimit(gameType, rule string) (Radius int, S, B LimitMinMax) {
	Radius, S, B = 0, LimitMinMax{}, LimitMinMax{}

	if gameType == constants.GAME_TYPE_GOL {
		S = LimitMinMax{
			Min: 2,
			Max: 3,
		}

		B = LimitMinMax{
			Min: 3,
			Max: 3,
		}

		Radius = 1
	} else if gameType == constants.GAME_TYPE_LTL {
		rule := GetLtLRule(rule)

		S = rule.S
		B = rule.B
		Radius = rule.Radius
	}

	return Radius, S, B
}
