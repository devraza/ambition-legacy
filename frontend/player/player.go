package player

import ()

// The player struct
type Player struct {
	Health      int
	MaxHealth   int
	Defence     int
	Level       int
	Exp         float32
	NextExp     float32
	Ambition    float32
	MaxAmbition float32
}

// Create the maps for the level/(max) ambition gates
var level_gates = make(map[int]float32)
var level_ambition = make(map[int]float32)

// TODO(midnadimple): Move player initialization to server upon login
func GetPlayer() Player {
	return Player{
		Health:      100,
		MaxHealth:   100,
		Defence:     0,
		Level:       1,
		Exp:         0.0,
		NextExp:     0.0,
		Ambition:    0.0,               // NOTE(midnadimple): In the future this will be affected by player activity
		MaxAmbition: level_ambition[1], // NOTE(midnadimple): In the future this will be affected by player activity
	}
}

// Formula for XP gain - extremely simple
func gain(basexp float32, modifier float32) float32 {
	gain := basexp * modifier
	return gain
}

// Update the player
func (p *Player) Update() {
	// TODO(midnadimple): update health upon damage

	// Auto-generate the level gates
	level_gates[1] = 500
	for i := 0; i <= 1000; i++ {
		if i >= 2 {
			switch {
			case i <= 10:
				level_gates[i] = level_gates[i-1] * 1.2
			case (i <= 100) && (i > 10):
				level_gates[i] = level_gates[i-1] * 1.1
			case (i <= 1000) && (i > 100):
				level_gates[i] = level_gates[i-1] * 1.005
			}
		}
	}

	// Auto-generate maximum ambition gates
	level_ambition[1] = 10
	for i := 0; i <= 1000; i++ {
		if i >= 2 {
			switch {
			case i <= 10:
				level_ambition[i] = level_ambition[i-1] * 1.1
			case (i <= 100) && (i > 10):
				level_ambition[i] = level_ambition[i-1] * 1.05
			case (i <= 1000) && (i > 100):
				level_ambition[i] = level_ambition[i-1] * 1.00005
			}
		}
	}

	// Set the XP needed for the player to reach the next level
	p.NextExp = level_gates[p.Level]

	// Set the XP to 0 and increase both the level and max ambition on level up
	if p.Exp >= level_gates[p.Level] {
		p.Exp = 0.0
		p.Level += 1
		p.MaxAmbition = level_ambition[p.Level]
	}
}
