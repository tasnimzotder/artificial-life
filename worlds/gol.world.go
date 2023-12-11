package worlds

import (
	"math/rand"
)

type GoLWorld struct {
	area     []byte
	nextArea []byte // temporary area for next generation
	width    int
	height   int
	length   int
}

func NewGoLWorld(width, height int) *GoLWorld {
	area := make([]byte, width*height)
	nextArea := make([]byte, width*height)

	w := &GoLWorld{
		area:     area,
		nextArea: nextArea,
		width:    width,
		height:   height,
		length:   width * height,
	}

	return w
}

func (gl *GoLWorld) GetSize() (width, height int) {
	return gl.width, gl.height
}

func (gl *GoLWorld) GetLength() int {
	return gl.length
}

func (gl *GoLWorld) InitRandom() {
	gl.ClearWorld()

	for idx := 0; idx < gl.length/4; idx++ {
		x := rand.Intn(gl.width)
		y := rand.Intn(gl.height)

		(gl.area)[y*gl.width+x] = 1
	}
}

func (gl *GoLWorld) InitPreset(preset [][]uint8) {
	gl.ClearWorld()

	cRow := gl.height/2 - len(preset)/2
	cCol := gl.width/2 - len(preset[0])/2

	for r := 0; r < len(preset); r++ {
		for c := 0; c < len(preset[r]); c++ {
			//(gl.area)[(cRow+r)*gl.width+(cCol+c)] = preset[r][c]
			//	rotate by 90 degrees
			(gl.area)[(cCol+c)*gl.width+(cRow+r)] = preset[r][c]
		}
	}
}

func (gl *GoLWorld) ClearWorld() {
	for idx := 0; idx < gl.length; idx++ {
		(gl.area)[idx] = 0
		(gl.nextArea)[idx] = 0
	}
}

func (gl *GoLWorld) IsAlive(x, y int) bool {
	x = (x + gl.width) % gl.width
	y = (y + gl.height) % gl.height

	return (gl.area)[y*gl.width+x]&0x01 == 0x01
}

func (gl *GoLWorld) AliveNeighbours(x, y int) uint8 {
	aliveCount := uint8(0)

	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}

			if gl.IsAlive(x+c, y+r) {
				aliveCount += 1
			}
		}
	}

	return aliveCount
}

func (gl *GoLWorld) SetCell(x, y int, alive bool, neighbours uint8) {
	idx := y*gl.width + x

	if alive {
		// sets the first bit to 1 and shifts the neighbours to the left by 1
		(gl.nextArea)[idx] = 1 | byte(neighbours<<1)
	} else {
		// sets the first bit to 0 and shifts the neighbours to the left by 1
		(gl.nextArea)[idx] = neighbours << 1
	}
}

func (gl *GoLWorld) NextCellGeneration(x, y int) bool {
	currentState := gl.IsAlive(x, y)
	neighbours := gl.AliveNeighbours(x, y)

	if currentState {
		return neighbours == 2 || neighbours == 3
	} else {
		return neighbours == 3
	}
}

func (gl *GoLWorld) NextGeneration() {
	for y := 0; y < gl.height; y++ {
		for x := 0; x < gl.width; x++ {
			nextState := gl.NextCellGeneration(x, y)
			neighbours := gl.AliveNeighbours(x, y)

			gl.SetCell(x, y, nextState, neighbours)
		}
	}

	// swap the areas
	gl.area, gl.nextArea = gl.nextArea, gl.area
}
