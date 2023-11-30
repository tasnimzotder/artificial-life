package presets

func GetPreset(gameType string, preset string) [][]int {
	if gameType == "GoL" {
		if preset == "Glider" {
			return GliderGOL
		} else if preset == "GliderGun" {
			return GliderGunGOL
		} else if preset == "Pulsar" {
			return PulsarGOL
		} else if preset == "Pentadecathlon" {
			return PentadecathlonGOL
		}

		return [][]int{}
	}

	return [][]int{}
}
