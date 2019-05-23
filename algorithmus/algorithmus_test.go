package algorithmus

import (
	"math"
	"testing"

	"github.com/J-Rocke/gropro/model"
	"github.com/stretchr/testify/assert"
)

func Test_InitStaaten(t *testing.T) {
	staaten := []model.Staat{
		{"DE", 420, model.Koordinate{31, 32}},
		{"NL", 1230, model.Koordinate{11, 33}},
	}

	algo := algorithmus{}
	algo.initStaaten(staaten)

	assert.Equal(t, []model.Staat{
		{"DE", math.Sqrt(420) / math.Pi, model.Koordinate{31, 32}},
		{"NL", math.Sqrt(1230) / math.Pi, model.Koordinate{11, 33}},
	}, algo.loesung.Staaten)
}

func Test_Abstand(t *testing.T) {
	assert.Equal(t, 5., abstand(
		model.Staat{Position: model.Koordinate{45, 23}},
		model.Staat{Position: model.Koordinate{41, 26}},
	))
}
