package games

import (
	"github.com/gabrieltong/gbox/games/ddz"
	"github.com/gabrieltong/gbox/games/poker"
	"github.com/gabrieltong/gbox/models"
)

func GetGameInstance(Type string) models.GameInterface {
	var game models.GameInterface = nil
	// var config interface{}
	switch Type {
	case "ddz":
		game = &ddz.Ddz{}
		// config = &ddz.DdzConfig{}
		break
	case "easy":
		config := &poker.EashConfig{}
		game = &poker.NewEasy(config)
	}
	// game.SetConfig(config)
	return game
}
