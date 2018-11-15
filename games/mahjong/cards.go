package mahjong

import (
	"gabrieltong/gbox/models"
)

const Tiao uint8 = 1
const Bing uint8 = 2
const Wan uint8 = 3
const Feng uint8 = 4
const Hua uint8 = 5

var Seed = [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
var FengSeed = [7]uint8{1, 2, 3, 4, 5, 6, 7}
var Feng1Seed = [4]uint8{1, 2, 3, 4}
var Feng5Seed = [3]uint8{5, 6, 7}
var HuaSeed = [8]uint8{1, 2, 3, 4, 5, 6, 7, 8}

type CardsConfig struct {
	enable_tiao bool
	enable_bing bool
	enable_wan  bool
	enable_feng bool
	enable_f1   bool
	enable_f5   bool
	enable_hua  bool
}

func InitCards(config *CardsConfig) []*models.Card {
	cards := make([]*models.Card, 0, 10)
	if config.enable_tiao {
		for _, No := range Seed {
			cards = append(cards, &models.Card{Type: Tiao, No: No})
		}
	}

	if config.enable_bing {
		for _, No := range Seed {
			cards = append(cards, &models.Card{Type: Bing, No: No})
		}
	}

	if config.enable_wan {
		for _, No := range Seed {
			cards = append(cards, &models.Card{Type: Wan, No: No})
		}
	}

	if config.enable_feng {
		for _, No := range FengSeed {
			cards = append(cards, &models.Card{Type: Feng, No: No})
		}
	} else if config.enable_f1 {
		for _, No := range Feng1Seed {
			cards = append(cards, &models.Card{Type: Feng, No: No})
		}
	} else if config.enable_f5 {
		for _, No := range Feng5Seed {
			cards = append(cards, &models.Card{Type: Feng, No: No})
		}
	}

	if config.enable_hua {
		for _, No := range HuaSeed {
			cards = append(cards, &models.Card{Type: Hua, No: No})
		}
	}

	return cards
}

// func init() {
// 	InitCards(nil)
// }
