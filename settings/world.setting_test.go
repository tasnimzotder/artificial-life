package settings

import (
	"github.com/tasnimzotder/artificial-life/data"
	"testing"

	"github.com/tasnimzotder/artificial-life/constants"
)

func TestNewWorld(t *testing.T) {
	width, height := 100, 100
	gs := NewGameSettings()
	world := NewWorld(width, height, gs)

	if world == nil {
		t.Fatal("NewWorld() should not return nil")
	}

	if world.GoLWorld == nil {
		t.Error("NewWorld() should initialize GoLWorld")
	}

	if world.GS != gs {
		t.Error("NewWorld() should assign GameSettings to GS")
	}
}

func TestWorldInitRandom(t *testing.T) {
	width, height := 100, 100
	gs := NewGameSettings()
	world := NewWorld(width, height, gs)

	world.InitRandom(constants.GameTypeGol)

	// Since the world is initialized randomly, it's not feasible to test the exact state.
	// Instead, we can test if the world is not empty.
	isEmpty := true
	for _, cell := range world.GoLWorld.GetArea() {
		if cell != 0 {
			isEmpty = false
			break
		}
	}

	if isEmpty {
		t.Error("InitRandom() should initialize a non-empty world")
	}
}

func TestWorldInitPreset(t *testing.T) {
	width, height := 100, 100
	gs := NewGameSettings()
	world := NewWorld(width, height, gs)

	presetName := "Glider" // Assuming "Glider" is a valid preset in the data package
	world.InitPreset(constants.GameTypeGol, presetName)

	// Test if the world contains the preset pattern.
	// This assumes that the preset is smaller than the world size and starts at the top-left corner.
	preset := data.GetPreset(constants.GameTypeGol, presetName)
	if preset.Name == "" {
		t.Fatalf("Preset %s does not exist", presetName)
	}

	// todo: implement this
}

func TestWorldNextGeneration(t *testing.T) {
	width, height := 100, 100
	gs := NewGameSettings()
	world := NewWorld(width, height, gs)

	world.InitRandom(constants.GameTypeGol)
	initialState := make([]byte, len(world.GoLWorld.GetArea()))
	copy(initialState, world.GoLWorld.GetArea())

	world.NextGeneration()

	// Test if the world has changed after calling NextGeneration.
	hasChanged := false
	for i, cell := range world.GoLWorld.GetArea() {
		if cell != initialState[i] {
			hasChanged = true
			break
		}
	}

	if !hasChanged {
		t.Error("NextGeneration() should change the world state")
	}
}

func TestWorldDraw(t *testing.T) {
	// todo: implement this
}
