package mahjong

import (
	"gabrieltong/gbox/models"
)

type Mahjong struct {
	*models.Game
}

type MahjongConfig struct {
}

func NewMahjong(config *MahjongConfig) *Mahjong {
	return &Mahjong{}
}
