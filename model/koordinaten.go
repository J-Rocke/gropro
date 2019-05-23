package model

import (
	"math"
)

type Koordinate struct {
	X float64
	Y float64
}

// Add addiert zwei Koordinaten
func (k Koordinate) Add(o Koordinate) Koordinate {
	k.X += o.X
	k.Y += o.Y
	return k
}

// Inverse gibt das additive Inverse der Koordinate
func (k Koordinate) Inverse() Koordinate {
	return k.Multiply(-1)
}

// Length berechnet die Länge der Koordinate
func (k Koordinate) Length() float64 {
	l := math.Sqrt(k.X*k.X + k.Y*k.Y)
	if math.IsNaN(l) {
		return 0
	}
	return l
}

// Abstand berechnet den Abstand zwischen zwei Koordinaten
func (k Koordinate) Abstand(o Koordinate) float64 {
	return k.Add(o.Inverse()).Length()
}

// Normalize normiert die Koordinate auf die Länge 1
func (k Koordinate) Normalize() Koordinate {
	l := k.Length()
	if l == 0 {
		return Koordinate{}
	}
	return k.Multiply(1. / l)
}

// Multiply streckt die Koordinate um Faktor c
func (k Koordinate) Multiply(c float64) Koordinate {
	x := k.X * c
	y := k.Y * c
	if math.IsNaN(x) {
		x = 0
	}
	if math.IsNaN(y) {
		y = 0
	}
	return Koordinate{x, y}
}

//Richtung berechnet den normalisierten Richtungsvektor von der Koordinate in Richtung Koordinate o
func (k Koordinate) Richtung(o Koordinate) Koordinate {
	return o.Add(k.Inverse()).Normalize()
}
