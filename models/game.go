package models

type GameInterface interface {
	setConfig(config interface{})
	getConfig(config interface{})
	handleAction(action interface{})
}

type GameConfig struct {
	*TableConfig
	Size int
}

func NewGameConfig() *GameConfig {
	config := &GameConfig{}
	config.TableConfig = NewTableConfig()
	return config
}

type Game struct {
	*Table
	config  interface{}
	Size    uint8
	OutList map[Player]Card
}

func (g Game) SetConfig(config interface{}) {
	g.config = config
}

func (g Game) GetConfig(config interface{}) interface{} {
	return g.config
}

func NewGame() *Game {
	gameConfig := &GameConfig{Size: 4}
	game := &Game{}
	game.config = gameConfig
	game.Table = NewTable(gameConfig.TableConfig)
	return game
}
