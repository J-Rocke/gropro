package algorithmus

import (
	"math"

	"github.com/J-Rocke/gropro/model"
)

const c float64 = 0.5 // Dämpfung
var n = 1000          // Anzahl Iterationen

func Loese(ac *model.AusgangsdatenContainer) *model.LoesungsContainer {
	loesung := model.LoesungsContainer{
		Titel:   ac.Titel,
		Staaten: initStaaten(ac.Staaten),
	}
	for i := 0; i < n; i++ {
		loesung = iteration(loesung, ac.Nachbarschaften)
		loesung.Iteration = i + 1
	}

	return &loesung
}

func iteration(l model.LoesungsContainer, nachbarn model.Nachbarschaften) model.LoesungsContainer {
	kraefte := map[string]model.Koordinate{}

	for _, a := range l.Staaten {
		for _, b := range l.Staaten {
			if a.ID == b.ID {
				continue
			}
			kraefte[a.ID] = kraefte[a.ID].Add(kraftAufVon(a, b, nachbarn.SindNachbarn(a, b)))
		}
	}
	return model.LoesungsContainer{
		Titel:   l.Titel,
		Staaten: wendeAn(l.Staaten, kraefte),
	}
}

// Berechne Radius
func initStaaten(before []model.Staat) []model.Staat {
	after := []model.Staat{}
	for _, Staat := range before {
		Staat.Kennwert = math.Sqrt(Staat.Kennwert) / math.Pi
		after = append(after, Staat)
	}

	return after
}

func wendeAn(staaten []model.Staat, kraefte map[string]model.Koordinate) []model.Staat {
	for i, staat := range staaten {
		new := staat.Position.Add(kraefte[staat.ID].Multiply(c))
		staat.Position = new
		staaten[i] = staat
	}
	return staaten
}

// Kraft auf Staat a von Staat b
func kraftAufVon(a, b model.Staat, nachbarn bool) model.Koordinate {
	abstand := abstand(a, b)
	if nachbarn { // Falls Nachbarn soll der Abstand eleminiert werden. Egal ob positiv oder negativ
		return a.Position.Richtung(b.Position).
			Multiply(abstand).
			Multiply(0.5) // Hälftig auf jeden Staat
	}
	if abstand < 0 { // Falls keine Nachbarn sollen sie sich nicht Überschneiden
		return a.Position.Richtung(b.Position).
			Multiply(abstand).
			Multiply(1.2). // Push more than needed
			Multiply(0.5)  // Hälftig auf jeden Staat
	}
	return model.Koordinate{} // keine Überschneidung und keine Nachbarn => keine Kraft
}

// Abstand zwischen zwei Kreisen
func abstand(a, b model.Staat) float64 {
	return a.Position.Abstand(b.Position) - a.Kennwert - b.Kennwert
}
