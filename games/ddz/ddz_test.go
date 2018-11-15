package ddz_test

import (
	"testing"

	"github.com/gabrieltong/gbox/games/ddz"
	"github.com/gabrieltong/gbox/models"
)

func TestSum(t *testing.T) {
	total := ddz.Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestNewDdz(t *testing.T) {
	ddzConfig := &ddz.DdzConfig{}
	ddz := ddz.NewDdz(ddzConfig)

	p1 := models.SamplePlayer()
	p2 := models.SamplePlayer()
	p3 := models.SamplePlayer()

	ddz.Join(p1, 0)
	ddz.Join(p2, 1)
	ddz.Join(p3, 3)

	ddz.Ready(p1, true)
	ddz.Ready(p2, true)
	ddz.Ready(p3, true)

	if ddz.AllReady() != true {
		t.Errorf("Allready should be true")
	}

	// ddz.Table.BeginIndex
}
