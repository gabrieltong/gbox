package models

import (
	"fmt"
)

// TODO support Review
// TODO support randomPosition

func NewTable(tableConfig *TableConfig) *Table {
	table := &Table{}
	table.TableConfig = tableConfig
	return table
}

func NewTableViewer(player *Player) *TableViewer {
	tableViewer := &TableViewer{player}
	return tableViewer
}

func NewTablePlayer() *TablePlayer {
	tablePlayer := &TablePlayer{}
	return tablePlayer
}

type Table struct {
	*TableConfig
	TablePlayers map[uint8]*TablePlayer
	Viewers      map[PlayerId]*TableViewer
}

type TableConfig struct {
	BeginIndex     uint8
	Max            uint8
	Min            uint8
	viewer         bool
	RandomPosition bool
}

func NewTableConfig() *TableConfig {
	return &TableConfig{}
}

type TableViewer struct {
	*Player
}

type TablePlayer struct {
	*Player
	OutList []Card
	Ready   bool
	Score   int32
}

func (t *Table) JoinAsViewer(Player *Player) bool {
	if t.Viewers[Player.id] == nil {
		t.Viewers[Player.id] = NewTableViewer(Player)
	}
	return true
}

func (t *Table) AllReady() bool {
	for _, tablePlayer := range t.TablePlayers {
		if tablePlayer.Ready != true {
			return false
		}
	}
	return true
}

func (t *Table) Ready(player *Player, ready bool) bool {
	tablePlayer := t.GetTablePlayer(player)
	if tablePlayer == nil {
		return false
	} else {
		tablePlayer.Ready = ready
		return true
	}
}

func (t *Table) Enough() bool {
	return uint8(len(t.TablePlayers)) >= t.Min
}

func (t *Table) GetTablePlayer(player *Player) *TablePlayer {
	for _, tablePlayer := range t.TablePlayers {
		if tablePlayer.Player == player {
			return tablePlayer
		}
	}
	return nil
}

func (t *Table) NextPlayer(current uint8) (*Player, uint8) {
	fmt.Println("in...")
	if current > t.Max-1 {
		current = t.Max - 1
	}
	position := t.NextPosition(current)
	var player *Player = nil
	for position != current && player == nil {
		if t.TablePlayers[position] != nil {
			player = t.TablePlayers[position].Player
		} else {
			position = t.NextPosition(position)
		}
	}
	return player, position
}

func (t *Table) PrePlayer(current uint8) (*Player, uint8) {
	fmt.Println("in...")
	if current > t.Max-1 {
		current = t.Max - 1
	}
	position := t.PrePosition(current)
	var player *Player = nil
	for position != current && player == nil {
		if t.TablePlayers[position] != nil {
			player = t.TablePlayers[position].Player
		}
		position = t.PrePosition(position)
	}
	return player, position
}

func (t *Table) NextPosition(position uint8) uint8 {
	if position == t.Max {
		return 0
	} else {
		return position + 1
	}
}

func (t *Table) PrePosition(position uint8) uint8 {
	if position == 0 {
		return t.Max - 1
	} else {
		return position - 1
	}
}

func (t *Table) Join(player *Player, position uint8) (uint8, bool) {
	if t.TablePlayers[position] == nil {
		return 0, false
	} else {
		t.TablePlayers[position] = &TablePlayer{Player: player}
		return position, true
	}
}

func (t *Table) Leave(player *Player) bool {
	for position, tablePlayer := range t.TablePlayers {
		if tablePlayer.Player == player {
			t.TablePlayers[position] = nil
			return true
		}
	}
	return false
}
