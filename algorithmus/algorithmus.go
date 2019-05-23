package algorithmus

import (
	"math"

	"github.com/J-Rocke/gropro/model"
)

// Dämpfung
const c float64 = 0.5

// Anzahl Iterationen
const n = 1000

// Loese Problemstellung
func Loese(ac *model.AusgangsdatenContainer) *model.LoesungsContainer {
	algo := algorithmus{
		loesung: model.LoesungsContainer{
			Titel: ac.Titel,
		},
		nachbarschaften: ac.Nachbarschaften,
	}
	algo.initStaaten(ac.Staaten)
	for i := 0; i < n; i++ {
		algo.iteration()
	}

	return &algo.loesung
}

type algorithmus struct {
	loesung         model.LoesungsContainer
	nachbarschaften model.Nachbarschaften
	kraefte         map[string]model.Koordinate
}

// Initialisiere Staaten und berechne Radius
func (algo *algorithmus) initStaaten(staaten []model.Staat) {
	for _, Staat := range staaten {
		Staat.Kennwert = math.Sqrt(Staat.Kennwert) / math.Pi
		algo.loesung.Staaten = append(algo.loesung.Staaten, Staat)
	}
}

func (algo *algorithmus) iteration() {
	// initialisiere Kräfte für dies Iteration
	algo.kraefte = map[string]model.Koordinate{}
	// Iteriere über alle Staaten und Abstände
	for _, a := range algo.loesung.Staaten {
		for _, b := range algo.loesung.Staaten {
			if a.ID == b.ID {
				continue
			}
			// Berechne Kraft
			algo.kraftAufVon(a, b)
		}
	}
	// Wende Kräfte an
	algo.wendeAn()
	algo.loesung.Iteration++
}

// Berechne Kraft auf Staat a von Staat b
func (algo *algorithmus) kraftAufVon(a, b model.Staat) {
	abstand := abstand(a, b)
	kraft := model.Koordinate{}
	if algo.nachbarschaften.SindNachbarn(a, b) {
		// Falls Nachbarn soll der Abstand eleminiert werden. Egal ob positiv oder negativ
		kraft = a.Position.Richtung(b.Position).
			Multiply(abstand)
	} else if abstand < 0 {
		// Falls keine Nachbarn sollen sie sich nicht Überschneiden
		kraft = a.Position.Richtung(b.Position).
			Multiply(abstand).
			Multiply(1.2) // Lieber etwas weiter auseinander
	} else {
		// keine Überschneidung und keine Nachbarn => keine Kraft
	}
	algo.kraefte[a.ID] = algo.kraefte[a.ID].Add(
		kraft.Multiply(0.5), // Hälftig auf jeden Staat
	)
}

// Wende berechnete Kräfte an
func (algo *algorithmus) wendeAn() {
	for i, staat := range algo.loesung.Staaten {
		staat.Position = staat.Position.Add( // Wende Kraft an
			algo.kraefte[staat.ID].Multiply(c), // mit Dämpfung
		)
		algo.loesung.Staaten[i] = staat
	}
}

// Abstand zwischen zwei Kreisen
func abstand(a, b model.Staat) float64 {
	return a.Position.Abstand(b.Position) - a.Kennwert - b.Kennwert
}
