package worlds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests

func TestNewGoLWorld(t *testing.T) {
	width, height := 10, 10
	world := NewGoLWorld(width, height)

	assert.NotNil(t, world, "NewGoLWorld() should not return nil")
	assert.Equal(t, width*height, len(world.GetArea()), "NewGoLWorld() should create an area with length of width*height")
}

func TestGoLWorld_InitRandom(t *testing.T) {
	width, height := 10, 10
	world := NewGoLWorld(width, height)

	world.InitRandom()

	// Check if the world has both alive and dead cells
	foundAlive, foundDead := false, false
	for _, cell := range world.GetArea() {
		if cell&0x01 == 0x01 {
			foundAlive = true
		} else {
			foundDead = true
		}

		if foundAlive && foundDead {
			break
		}
	}

	assert.True(t, foundAlive, "InitRandom() should create a world with both alive and dead cells")
	assert.True(t, foundDead, "InitRandom() should create a world with both alive and dead cells")
}

func TestGoLWorld_InitPreset(t *testing.T) {
	width, height := 10, 10
	world := NewGoLWorld(width, height)

	preset := [][]uint8{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}
	world.InitPreset(preset)

	// Check if the preset is correctly placed in the center of the world
	cx, cy := width/2, height/2
	for y, row := range preset {
		for x, cell := range row {
			worldX, worldY := cx-len(preset[0])/2+x, cy-len(preset)/2+y
			assert.Equal(t, cell == 1, world.IsAlive(worldX, worldY), "InitPreset() did not correctly initialize the preset at position (%d, %d)", worldX, worldY)
		}
	}
}

func TestGoLWorld_InitPreset_Neighbours(t *testing.T) {
	presets := [][][]uint8{
		{
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		},
		{

			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 1, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
		}, // glider
	}

	genZero := [][]byte{
		{

			0x02, 0x02, 0x02, 0x00,
			0x02, 0x01, 0x02, 0x00,
			0x02, 0x02, 0x02, 0x00,
			0x00, 0x00, 0x00, 0x00,
		}, {
			0x02, 0x02, 0x02, 0x00,
			0x04, 0x03, 0x04, 0x00,
			0x04, 0x03, 0x04, 0x00,
			0x02, 0x02, 0x02, 0x00,
		},
		{
			0x04, 0x03, 0x04, 0x00,
			0x06, 0x05, 0x06, 0x00,
			0x04, 0x03, 0x04, 0x00,
			0x04, 0x04, 0x04, 0x00,
		},
		{
			0x02, 0x02, 0x04, 0x02, 0x02,
			0x02, 0x03, 0x08, 0x05, 0x04,
			0x02, 0x06, 0x09, 0x07, 0x04,
			0x00, 0x04, 0x05, 0x06, 0x02,
			0x00, 0x02, 0x02, 0x02, 0x00,
		}, // glider
	}

	for i, preset := range presets {
		width, height := len(preset[0]), len(preset)
		world := NewGoLWorld(width, height)

		world.InitPreset(preset)

		// Check if the preset is correctly placed in the center of the world
		for y, row := range preset {
			for x, cell := range row {
				worldX, worldY := x, y
				assert.Equal(t, cell == 1, world.IsAlive(worldX, worldY), "InitPreset() did not correctly initialize the preset at position (%d, %d)", worldX, worldY)
			}
		}

		// check if the neighbors are correct
		assert.Equal(t, genZero[i], world.GetArea(), "InitPreset() did not correctly initialize the preset and its neighbors")

	}
}

func TestGoLWorld_UpdateCell(t *testing.T) {
	presets := [][][]uint8{
		{
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		},
		{

			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		},
	}

	genZero := [][]byte{
		{

			0x02, 0x02, 0x02, 0x00,
			0x02, 0x01, 0x02, 0x00,
			0x02, 0x02, 0x02, 0x00,
			0x00, 0x00, 0x00, 0x00,
		}, {
			0x02, 0x02, 0x02, 0x00,
			0x04, 0x03, 0x04, 0x00,
			0x04, 0x03, 0x04, 0x00,
			0x02, 0x02, 0x02, 0x00,
		},
		{
			0x04, 0x03, 0x04, 0x00,
			0x06, 0x05, 0x06, 0x00,
			0x04, 0x03, 0x04, 0x00,
			0x04, 0x04, 0x04, 0x00,
		},
	}

	for i, preset := range presets {
		world := NewGoLWorld(len(preset[0]), len(preset))

		for y := 0; y < len(preset); y++ {
			for x := 0; x < len(preset[y]); x++ {
				if preset[y][x] == 1 {
					world.UpdateCell(x, y, true)
				}
			}
		}

		// check if the neighbors are correct
		assert.Equal(t, genZero[i], world.GetArea(), "UpdateCell() did not correctly initialize the preset and its neighbors")
	}
}

func TestGoLWorld_ClearWorld(t *testing.T) {
	width, height := 10, 10
	world := NewGoLWorld(width, height)

	world.InitRandom()
	world.ClearWorld()

	for _, cell := range world.GetArea() {
		//if cell != 0 {
		//	t.Error("ClearWorld() should set all cells to dead")
		//}
		assert.Equal(t, byte(0), cell, "ClearWorld() should set all cells to dead")
	}
}

func TestGoLWorld_NextGeneration(t *testing.T) {
	presets := [][][]uint8{
		{
			{0, 0, 0, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 1, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
		}, // glider
	}

	genZero := [][]byte{
		{
			0x02, 0x02, 0x04, 0x02, 0x02,
			0x02, 0x03, 0x08, 0x05, 0x04,
			0x02, 0x06, 0x09, 0x07, 0x04,
			0x00, 0x04, 0x05, 0x06, 0x02,
			0x00, 0x02, 0x02, 0x02, 0x00,
		}, // glider
	}

	genOne := [][]byte{
		{
			0x00, 0x00, 0x02, 0x02, 0x02,
			0x02, 0x02, 0x06, 0x03, 0x04,
			0x02, 0x03, 0x0a, 0x07, 0x06,
			0x02, 0x04, 0x07, 0x05, 0x04,
			0x00, 0x02, 0x04, 0x04, 0x02,
		}, // glider
	}

	genTwo := [][]byte{
		{
			0x00, 0x02, 0x02, 0x02, 0x00,
			0x02, 0x02, 0x03, 0x06, 0x04,
			0x02, 0x04, 0x08, 0x09, 0x05,
			0x02, 0x02, 0x05, 0x07, 0x06,
			0x00, 0x02, 0x04, 0x04, 0x02,
		}, // glider
	}

	for i, preset := range presets {
		world := NewGoLWorld(len(preset[0]), len(preset))
		world.InitPreset(preset)

		assert.Equal(t, genZero[i], world.GetArea(), "InitPreset() did not correctly initialize the preset and its neighbors")

		world.NextGeneration()
		assert.Equal(t, genOne[i], world.GetArea(), "NextGeneration() 1 did not correctly initialize the preset and its neighbors")

		world.NextGeneration()
		assert.Equal(t, genTwo[i], world.GetArea(), "NextGeneration() 2 did not correctly initialize the preset and its neighbors")
	}

}

// Benchmarks

func BenchmarkGoLWorld_NextGeneration_512(b *testing.B) {
	world := NewGoLWorld(512, 512)

	for i := 0; i < b.N; i++ {
		world.NextGeneration()
	}
}

func BenchmarkGoLWorld_NextGeneration_1024(b *testing.B) {
	world := NewGoLWorld(1024, 1024)

	for i := 0; i < b.N; i++ {
		world.NextGeneration()
	}
}
