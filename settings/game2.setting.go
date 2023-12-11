package settings

type GameSettings2 struct {
	GameType  string
	GameTypes []string

	Presets []PresetType
	Preset  string

	WrapAround  bool
	IsPaused    bool
	IsRunning   bool
	IsReset     bool
	DesiredTPS  int
	ZoomLevel   int
	WorldWidth  int
	WorldHeight int

	//	ltl
	LtLRule string
	Radius  uint8
}

func NewGameSettings() *GameSettings2 {
	gs := &GameSettings2{}
	setDefaultSettings(gs)

	gs.IsReset = true
	gs.Preset = "Random"
	gs.GameType = gs.GameTypes[0]

	gs.ZoomLevel = 15
	gs.Radius = 1
	gs.LtLRule = "Life"

	return gs
}

func (gs *GameSettings2) ResetSettings() {
	setDefaultSettings(gs)
	gs.IsReset = false
}

func setDefaultSettings(gs *GameSettings2) {
	gs.WrapAround = true
	gs.IsPaused = true
	gs.IsReset = false
	gs.DesiredTPS = 10
	//gs.ZoomLevel = 15
	gs.WorldWidth = 60 * 13
	gs.WorldHeight = 70 * 13

	gs.GameTypes = []string{
		"Game of Life",
		"Larger than Life",
	}
	gs.Presets = GetAllGoLPresets()
}
