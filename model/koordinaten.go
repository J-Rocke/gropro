package model

import (
	"math"
)

type Koordinate struct {
	X float64
	Y float64
}

func (k Koordinate) Add(o Koordinate) Koordinate {
	k.X += o.X
	k.Y += o.Y
	return k
}

func (k Koordinate) Inverse() Koordinate {
	return k.Multiply(-1)
}

func (k Koordinate) Length() float64 {
	l := math.Sqrt(k.X*k.X + k.Y*k.Y)
	if math.IsNaN(l) {
		return 0
	}
	return l
}

func (k Koordinate) Abstand(o Koordinate) float64 {
	return k.Add(o.Inverse()).Length()
}

func (k Koordinate) Normalize() Koordinate {
	l := k.Length()
	if l == 0 {
		return Koordinate{}
	}
	return k.Multiply(1. / l)
}

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

// Normalisierter Richtungsvektor von der Koordinate in Richtung o
func (k Koordinate) Richtung(o Koordinate) Koordinate {
	return o.Add(k.Inverse()).Normalize()
}
