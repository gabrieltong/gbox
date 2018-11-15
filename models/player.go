package models

import "math/rand"

type PlayerId uint32

type Player struct {
	id PlayerId
}

type PlayerConfig struct {
}

// func NewPlayer(id PlayerId) *Player {

// }

func SamplePlayer() *Player {
	id := PlayerId(rand.Uint32())
	return &Player{id}
}
