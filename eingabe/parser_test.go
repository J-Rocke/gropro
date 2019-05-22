package eingabe

import (
	"testing"

	"github.com/J-Rocke/gropro/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Parse_Bierkonsum(t *testing.T) {
	c, err := ParseFromFile("../tests/bierkonsum.txt")
	assert.NoError(t, err)
	require.NotNil(t, c)
	assert.Equal(t, model.AusgangsdatenContainer{
		Titel: "Bierkonsum",
		Staaten: []model.Staat{
			{
				ID:       "D",
				Kennwert: 8692,
				Position: model.Koordinate{
					X: 10.0,
					Y: 51.3,
				},
			},
			{
				ID:       "NL",
				Kennwert: 1156,
				Position: model.Koordinate{
					X: 5.3,
					Y: 52.2,
				},
			},
			{
				ID:       "B",
				Kennwert: 781,
				Position: model.Koordinate{
					X: 4.8,
					Y: 50.7,
				},
			},
			{
				ID:       "L",
				Kennwert: 80,
				Position: model.Koordinate{
					X: 6.1,
					Y: 49.8,
				},
			},
			{
				ID:       "F",
				Kennwert: 2077,
				Position: model.Koordinate{
					X: 2.8,
					Y: 47.4,
				},
			},
			{
				ID:       "CH",
				Kennwert: 440,
				Position: model.Koordinate{
					X: 8.2,
					Y: 46.9,
				},
			},
			{
				ID:       "A",
				Kennwert: 945,
				Position: model.Koordinate{
					X: 14.2,
					Y: 47.6,
				},
			},
			{
				ID:       "CZ",
				Kennwert: 1573,
				Position: model.Koordinate{
					X: 15.3,
					Y: 49.8,
				},
			},
			{
				ID:       "PL",
				Kennwert: 3724,
				Position: model.Koordinate{
					X: 18.9,
					Y: 52.2,
				},
			},
			{
				ID:       "DK",
				Kennwert: 360,
				Position: model.Koordinate{
					X: 9.6,
					Y: 56.0,
				},
			},
		},
		Nachbarschaften: model.Nachbarschaften{
			"D":  map[string]bool{"A": true, "B": true, "CH": true, "CZ": true, "DK": true, "F": true, "L": true, "NL": true, "PL": true},
			"NL": map[string]bool{"B": true},
			"B":  map[string]bool{"L": true, "F": true},
			"F":  map[string]bool{"CH": true},
			"CH": map[string]bool{"A": true},
			"A":  map[string]bool{"CZ": true},
			"CZ": map[string]bool{"PL": true},
		},
	}, *c)
}
