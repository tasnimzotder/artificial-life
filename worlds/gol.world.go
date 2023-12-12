package worlds

import (
	"math/rand"
)

type GoLWorld struct {
	area     []byte
	nextArea []byte // temporary area for next generation
	width    int
	height   int
	length   int // width * height
}

func NewGoLWorld(width, height int) *GoLWorld {
	w := &GoLWorld{
		area:     make([]byte, width*height),
		nextArea: make([]byte, width*height),
		width:    width,
		height:   height,
		length:   width * height,
	}

	return w
}

func (w *GoLWorld) InitRandom() {
	// clear the world first to avoid having to check for dead cells
	w.ClearWorld()

	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			randomState := rand.Intn(2)

			if randomState == 1 {
				w.UpdateCell(x, y, true) // Set the cell to alive
			}
		}
	}
}

func (w *GoLWorld) InitPreset(preset [][]uint8) {
	// clear the world first to avoid having to check for dead cells
	w.ClearWorld()

	cRow := w.height/2 - len(preset)/2
	cCol := w.width/2 - len(preset[0])/2

	for r := 0; r < len(preset); r++ {
		for c := 0; c < len(preset[r]); c++ {
			cellState := preset[r][c]&0x01 == 0x01

			if cellState {
				w.UpdateCell(cCol+c, cRow+r, cellState) // Set the cell to alive
			}
		}
	}
}

func (w *GoLWorld) ClearWorld() {
	for idx := range w.area {
		w.area[idx] = 0
		w.nextArea[idx] = 0
	}
}

func (w *GoLWorld) IsAlive(x, y int) bool {
	x = (x + w.width) % w.width
	y = (y + w.height) % w.height

	return w.area[y*w.width+x]&1 == 1 // Check if the cell is alive
}

func (w *GoLWorld) GetArea() []byte {
	//	return a copy of the area
	area := make([]byte, w.length)
	copy(area, w.area) // copy the area to avoid changing the original area

	return area
}

func (w *GoLWorld) GetLength() int {
	return w.length
}

func (w *GoLWorld) UpdateCell(x, y int, set bool) {
	idx := y*w.width + x

	if set {
		// Set the cell to alive
		w.area[idx] |= 0x01 // 0b00000001 Set the last bit to 1
	} else {
		// Set the cell to dead
		w.area[idx] &= 0xFE // 0b11111110 Set the last bit to 0
	}

	w.updateSurroundCells(x, y, set)
}
func (w *GoLWorld) updateSurroundCells(x, y int, set bool) {
	delta := byte(0x02)

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue // Skip the cell itself
			}

			nx, ny := (x+dx+w.width)%w.width, (y+dy+w.height)%w.height
			idx := ny*w.width + nx

			if set {
				w.area[idx] += delta // Increment neighbor count
			} else {
				w.area[idx] -= delta // Decrement neighbor count
			}
		}
	}
}

func (w *GoLWorld) NextGeneration() {
	copy(w.nextArea, w.area)
	tempArea := w.nextArea

	for idx := range w.area {
		if tempArea[idx] == 0 {
			continue // Skip dead cells to avoid unnecessary calculations
		}

		x := idx % w.width
		y := idx / w.width

		liveNeighbors := tempArea[idx] >> 1

		// game rules
		if tempArea[idx]&0x01 == 0x01 {
			if liveNeighbors < 2 || liveNeighbors > 3 {
				w.UpdateCell(x, y, false)
			}
		} else {
			if liveNeighbors == 3 {
				w.UpdateCell(x, y, true)
			}
		}
	}
}
