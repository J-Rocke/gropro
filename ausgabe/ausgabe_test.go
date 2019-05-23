package ausgabe

import (
	"io/ioutil"
	"testing"

	"github.com/J-Rocke/gropro/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Ausgabe(t *testing.T) {
	soll := `reset
set xrange [2:12]
set yrange [-1:9]
set size ratio 1.0
set title "Test der Staaten, Iteration: 42"
unset xtics
unset ytics
$data << EOD
10 6 2 D 0
6 4 4 NL 1
EOD
plot \
'$data' using 1:2:3:5 with circles lc var notitle, \
'$data' using 1:2:4:5 with labels font "arial,9" tc var notitle
`

	l := &model.LoesungsContainer{
		Titel:     "Test der Staaten",
		Iteration: 42,
		Staaten: []model.Staat{
			model.Staat{"D", 2., model.Koordinate{10, 6}},
			model.Staat{"NL", 4., model.Koordinate{6, 4}},
		},
	}

	file, _ := ioutil.TempFile("", "test-*.gnu")
	err := GNUPlotToFile(file.Name(), l)
	require.NoError(t, err)
	bytes, _ := ioutil.ReadFile(file.Name())
	assert.Equal(t, soll, string(bytes))
}
