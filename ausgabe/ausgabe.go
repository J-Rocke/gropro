package ausgabe

import (
	"fmt"
	"math"
	"os"

	"github.com/J-Rocke/gropro/model"
)

func GNUPlotToFile(filename string, loesung *model.LoesungsContainer) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	min, max := findExtrema(loesung.Staaten)
	file.WriteString("reset\n")
	fmt.Fprintf(file, "set xrange [%v:%v]\n", min.X, max.X)
	fmt.Fprintf(file, "set yrange [%v:%v]\n", min.Y, max.Y)
	file.WriteString("set size ratio 1.0\n")
	file.WriteString("unset xtics\n")
	file.WriteString("unset ytics\n")
	file.WriteString("$data << EOD\n")
	for i, staat := range loesung.Staaten {
		fmt.Fprintf(file, "%v %v %v %s %v\n", staat.Position.X, staat.Position.Y, staat.Kennwert, staat.ID, i)
	}
	file.WriteString("EOD\n")
	file.WriteString("plot \\\n")
	file.WriteString("'$data' using 1:2:3:5 with circles lc var notitle, \\\n")
	file.WriteString("'$data' using 1:2:4:5 with labels font \"arial,9\" tc var notitle \n")
	file.WriteString("pause -1")

	return nil
}

func findExtrema(staaten []model.Staat) (min model.Koordinate, max model.Koordinate) {
	min = model.Koordinate{math.MaxFloat64, math.MaxFloat64}
	max = model.Koordinate{-math.MaxFloat64, -math.MaxFloat64}

	for _, staat := range staaten {
		if staat.Position.X-float64(staat.Kennwert) < min.X {
			min.X = staat.Position.X - float64(staat.Kennwert)
		}
		if staat.Position.Y-float64(staat.Kennwert) < min.Y {
			min.Y = staat.Position.Y - float64(staat.Kennwert)
		}
		if staat.Position.X+float64(staat.Kennwert) > max.X {
			max.X = staat.Position.X + float64(staat.Kennwert)
		}
		if staat.Position.Y+float64(staat.Kennwert) > max.Y {
			max.Y = staat.Position.Y + float64(staat.Kennwert)
		}
	}
	xRange := max.X - min.X
	yRange := max.Y - min.Y
	// Make sure xRange == yRange
	if xRange > yRange {
		max.Y = max.Y + (xRange-yRange)/2
		min.Y = min.Y - (xRange-yRange)/2

	} else {
		max.X = max.X + (yRange-xRange)/2
		min.X = min.X - (yRange-xRange)/2
	}
	return
}
