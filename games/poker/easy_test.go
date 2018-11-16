package poker_test

import (
	"testing"

	"github.com/gabrieltong/gbox/games/poker"
)

func TestNewEasyConfig(T *testing.T) {
	easyConfig := poker.NewEasyConfig()

	if easyConfig.BeginIndex != 0 {
		T.Errorf("easyConfig.BeginIndex should set to 0")
	}

	if easyConfig.Max != 4 {
		T.Errorf("easyConfig.Max should set to 4")
	}

	if easyConfig.Min != 4 {
		T.Errorf("easyConfig.Max should set to 4")
	}
}

func TestNewEasy(T *testing.T) {
	easyConfig := poker.NewEasyConfig()
	easy := poker.NewEasy(easyConfig)

	if easy.EasyConfig.Max != 4 {
		T.Errorf("easy.easyConfig.Max should set to 4")
	}
}
