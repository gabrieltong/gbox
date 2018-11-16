package poker

import (
	"github.com/gabrieltong/gbox/models"
)

type Easy struct {
	*EasyConfig
	*models.Game
}

type EasyConfig struct {
	*models.GameConfig
}

func NewEasy(easyConfig *EasyConfig) *Easy {
	game := models.NewGame(easyConfig.GameConfig)

	easy := &Easy{
		EasyConfig: easyConfig,
		Game:       game,
	}

	return easy
}

func NewEasyConfig() *EasyConfig {
	easyConfig := &EasyConfig{
		GameConfig: &models.GameConfig{
			TableConfig: &models.TableConfig{
				BeginIndex: 0,
				Max:        4,
				Min:        4,
			},
		},
	}
	return easyConfig
}
