package algorithmus

import (
	"math"
	"testing"

	"github.com/J-Rocke/gropro/model"
	"github.com/stretchr/testify/assert"
)

func Test_InitStaaten(t *testing.T) {
	before := []model.Staat{
		{"DE", 420, model.Koordinate{31, 32}},
		{"NL", 1230, model.Koordinate{11, 33}},
	}

	after := initStaaten(before)

	assert.Equal(t, []model.Staat{
		{"DE", int64(math.Sqrt(float64(420)) / math.Pi), model.Koordinate{31, 32}},
		{"NL", int64(math.Sqrt(float64(1230)) / math.Pi), model.Koordinate{11, 33}},
	}, after)
}

// func divideByPi(i int64) int64 {
// 	return int64(float64(i) / math.Pi)
// }

// func Test_DivideByPi(t *testing.T) {
// 	assert.Equal(t, int64(float64(42)/math.Pi), divideByPi(42))
// }
