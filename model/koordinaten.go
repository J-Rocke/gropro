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
	k.X = -k.X
	k.Y = -k.Y
	return k
}

func (k Koordinate) Length() float64 {
	return math.Sqrt(k.X*k.X + k.Y*k.Y)
}

func (k Koordinate) Abstand(o Koordinate) float64 {
	return k.Add(o.Inverse()).Length()
}

func (k Koordinate) Normalize() Koordinate {
	l := k.Length()
	return Koordinate{X: k.X / l, Y: k.Y / l}
}

func (k Koordinate) Multiply(c float64) Koordinate {
	return Koordinate{X: k.X * c, Y: k.Y * c}
}

// Normalisierter Richtungsvektor von der Koordinate in Richtung o
func (k Koordinate) Richtung(o Koordinate) Koordinate {
	return o.Add(k.Inverse()).Normalize()
}
