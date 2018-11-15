package models

type GameInterface interface {
	SetConfig(config interface{})
	GetConfig(config interface{}) interface{}
}

type GameConfig struct {
	Size int
}

type Game struct {
	beginIndex uint8
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
	game := &Game{}
	game.Table = NewTable(&TableConfig{})
	return game
}
