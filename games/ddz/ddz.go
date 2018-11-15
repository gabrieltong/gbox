package ddz

import (
	"fmt"
	"gabrieltong/gbox/games/poker"
	"gabrieltong/gbox/models"

	"github.com/looplab/fsm"
)

func Sum(x int, y int) int {
	return x + y
}

type DdzConfig struct {
	*poker.CardsConfig
}

type Ddz struct {
	config *DdzConfig
	*models.Game
	FSM *fsm.FSM
}

func (d *Ddz) enterState(e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", e.Src, e.Dst)
}

func NewDdzConfig() *DdzConfig {
	config := &DdzConfig{}
	return config
}

func NewDdz(config *DdzConfig) *Ddz {
	ddz := &Ddz{}
	ddz.Game = models.NewGame()
	ddz.config = config
	ddz.FSM = fsm.NewFSM(
		"dispatch",
		fsm.Events{
			{Name: "dispatched", Src: []string{"dispatch"}, Dst: "vie"},
			{Name: "vied_to_getbottom", Src: []string{"vie"}, Dst: "getbottom"},
			{Name: "vied_to_ti", Src: []string{"vie"}, Dst: "ti"},
			{Name: "ti_over", Src: []string{"ti"}, Dst: "getbottom"},
			{Name: "bottom_getted_to_play", Src: []string{"getbottom"}, Dst: "play"},
			{Name: "bottom_getted_to_base", Src: []string{"getbottom"}, Dst: "base"},
			{Name: "base_over", Src: []string{"base"}, Dst: "play"},
			{Name: "game_over", Src: []string{"play"}, Dst: "over"},
			{Name: "vie_failed", Src: []string{"vie"}, Dst: "abortion"},
			{Name: "restart", Src: []string{"abortion"}, Dst: "dispatch"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { ddz.enterState(e) },
		},
	)
	return ddz
}
