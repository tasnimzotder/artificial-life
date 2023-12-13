package settings

import "github.com/tasnimzotder/artificial-life/data"

type GameSettings struct {
	T         uint
	GameType  string
	GameTypes []string

	Presets []data.PresetType
	Preset  string

	WrapAround  bool
	IsPaused    bool
	IsRunning   bool
	IsReset     bool
	Replay      bool
	DesiredTPS  int
	ZoomLevel   int
	WorldWidth  int
	WorldHeight int

	//	ltl
	LtLRule string
	Radius  uint8
}

func NewGameSettings() *GameSettings {
	gs := &GameSettings{}
	setDefaultSettings(gs)

	gs.IsReset = true
	gs.Preset = "Random"
	gs.GameType = gs.GameTypes[0]

	gs.ZoomLevel = 1
	gs.DesiredTPS = 2
	gs.Radius = 1
	gs.LtLRule = "Life"

	return gs
}

func (gs *GameSettings) ResetSettings() {
	setDefaultSettings(gs)
	gs.IsReset = false
	gs.Replay = false
}

func setDefaultSettings(gs *GameSettings) {
	gs.T = 0
	gs.WrapAround = true
	gs.IsPaused = true
	gs.IsReset = false

	gs.WorldWidth = 256 * 4
	gs.WorldHeight = 256 * 4

	gs.GameTypes = []string{
		"Game of Life",
		//"Larger than Life",
		//"Smooth Life",
	}
	gs.Presets = data.GetAllGoLPresets()
}
