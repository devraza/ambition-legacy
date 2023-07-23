package main

import (
	"math/rand"
)

type Player struct {
	health int
	level int
	exp float32
	ambition float32
}

// NOTE(midnadimple): These gates are temporary. We'll decide on real values later
var level_gates = map[int]float32{
	1: 100.0,
	2: 150.0,
	3: 300.0,
}

// TODO(midnadimple): Move player initialization to server upon login
func initPlayer() Player {
	return Player{
		health: 100,
		level: 1,
		exp: 0.0,
		ambition: (rand.Float32() * 10), // NOTE(midnadimple): In the future this will be affected by player activity
	}
}

func (p *Player) update() {
	// TODO(midnadimple): update health upon damage
	
	if p.exp >= level_gates[p.level] {
		p.exp = 0.0
		p.level += 1
	}

	// NOTE(midnadimple): This formula for exp gain is pretty simple, maybe devraza can think of
    // a more practical one
	p.exp += 10.0 * (1.0/60.0) * p.ambition
	
}