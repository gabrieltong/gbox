package models

type Action string

type PlayerWaiting map[Action]WaitingDetail

type WaitingDetail interface {
	Valid(data interface{}) bool
}

type Waiting struct {
	game    *Game
	counter int
	items   map[PlayerId]map[Action]WaitingDetail
}

func (w *Waiting) wait(playerId PlayerId, action Action, extra WaitingDetail) bool {
	if w.items[playerId] == nil {
		w.items[playerId] = map[Action]WaitingDetail{}
	}
	if w.items[playerId][action] == nil {
		w.items[playerId][action] = extra
	} else {
		return false
	}
	return true
}

func (w *Waiting) isWaiting(playerId PlayerId, action Action, data interface{}) bool {
	if w.items[playerId] == nil {
		return false
	}

	if w.items[playerId][action] == nil {
		return false
	}

	return w.items[playerId][action].Valid(data)
}
