package poker

import (
	"github.com/gabrieltong/gbox/models"
)

const Heart uint8 = 11
const Spade uint8 = 12
const Diamond uint8 = 13
const Club uint8 = 14
const Joker uint8 = 15

var Seed = [13]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
var JokerSeed = [2]uint8{1, 2}

type CardsConfig struct {
	unit           uint8
	enable_heart   bool
	enable_spade   bool
	enable_diamond bool
	enable_club    bool
	enable_joker   bool
}

func NewCardsConfig() *CardsConfig {
	return &CardsConfig{
		unit:           1,
		enable_heart:   true,
		enable_spade:   true,
		enable_diamond: true,
		enable_club:    true,
		enable_joker:   true,
	}
}

func InitCards(config *CardsConfig) []*models.Card {
	cards := make([]*models.Card, 0, 10)
	for index := uint8(0); index < config.unit; index++ {
		if config.enable_heart {
			for _, No := range Seed {
				cards = append(cards, &models.Card{Type: Heart, No: No})
			}
		}

		if config.enable_spade {
			for _, No := range Seed {
				cards = append(cards, &models.Card{Type: Spade, No: No})
			}
		}

		if config.enable_diamond {
			for _, No := range Seed {
				cards = append(cards, &models.Card{Type: Diamond, No: No})
			}
		}

		if config.enable_club {
			for _, No := range Seed {
				cards = append(cards, &models.Card{Type: Club, No: No})
			}
		}

		if config.enable_joker {
			for _, No := range JokerSeed {
				cards = append(cards, &models.Card{Type: Joker, No: No})
			}
		}
	}
	return cards
}

// func init() {
// 	InitCards(nil)
// }
