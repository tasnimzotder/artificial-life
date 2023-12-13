package settings

import (
	"testing"

	"github.com/tasnimzotder/artificial-life/data"
)

func TestNewGameSettings(t *testing.T) {
	gs := NewGameSettings()

	if gs == nil {
		t.Error("NewGameSettings() should not return nil")
	}

	if gs.IsReset != true {
		t.Errorf("NewGameSettings() should set IsReset to true, got %v", gs.IsReset)
	}

	if gs.Preset != "Random" {
		t.Errorf("NewGameSettings() should set Preset to 'Random', got %s", gs.Preset)
	}

	if gs.GameType != gs.GameTypes[0] {
		t.Errorf("NewGameSettings() should set GameType to the first element of GameTypes, got %s", gs.GameType)
	}
}

func TestResetSettings(t *testing.T) {
	gs := NewGameSettings()
	gs.ResetSettings()

	if gs.IsReset != false {
		t.Errorf("ResetSettings() should set IsReset to false, got %v", gs.IsReset)
	}

	if gs.Replay != false {
		t.Errorf("ResetSettings() should set Replay to false, got %v", gs.Replay)
	}

	if gs.T != 0 {
		t.Errorf("ResetSettings() should set T to 0, got %d", gs.T)
	}

	if gs.WrapAround != true {
		t.Errorf("ResetSettings() should set WrapAround to true, got %v", gs.WrapAround)
	}

	if gs.IsPaused != true {
		t.Errorf("ResetSettings() should set IsPaused to true, got %v", gs.IsPaused)
	}

	if len(gs.GameTypes) == 0 {
		t.Error("ResetSettings() should populate GameTypes with at least one game type")
	}

	if len(gs.Presets) != len(data.GetAllGoLPresets()) {
		t.Errorf("ResetSettings() should populate Presets with GetAllGoLPresets(), got %d", len(gs.Presets))
	}
}
