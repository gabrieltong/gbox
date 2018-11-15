package games

import (
	"gabrieltong/gbox/games/ddz"
	"gabrieltong/gbox/models"
)

func GetInstance(Type string) models.GameInterface {
	var game interface{}
	// var config interface{}
	switch Type {
	case "ddz":
		game = &ddz.Ddz{}
		// config = &ddz.DdzConfig{}
		break
	}
	// game.SetConfig(config)
	return game
}
