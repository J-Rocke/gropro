package model

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	a := Koordinate{1, 2}
	b := Koordinate{3, 4}
	assert.Equal(t, Koordinate{1 + 3, 2 + 4}, a.Add(b))
	assert.Equal(t, Koordinate{1 + 3, 2 + 4}, b.Add(a))
}

func Test_Inverse(t *testing.T) {
	a := Koordinate{1, 2}
	b := Koordinate{-3, -4}
	assert.Equal(t, Koordinate{-1, -2}, a.Inverse())
	assert.Equal(t, Koordinate{3, 4}, b.Inverse())
}

func Test_Length(t *testing.T) {
	assert.Equal(t, 5., Koordinate{3, 4}.Length())
	assert.Equal(t, 5., Koordinate{-3, 4}.Length())
	assert.Equal(t, 4., Koordinate{0, 4}.Length())
	assert.Equal(t, 3., Koordinate{3, 0}.Length())
	assert.Equal(t, math.Sqrt(2.), Koordinate{1, 1}.Length())
	// NaN lenght
	assert.Equal(t, 0., Koordinate{math.NaN(), 1}.Length())
}

func Test_Abstand(t *testing.T) {
	assert.Equal(t, 1., Koordinate{0, 4}.Abstand(Koordinate{0, 5}))
	assert.Equal(t, 3., Koordinate{-1, 5}.Abstand(Koordinate{2, 5}))
	assert.Equal(t, math.Sqrt(3*3+5*5), Koordinate{-1, -3}.Abstand(Koordinate{2, 2}))
	assert.Equal(t, 0., Koordinate{-1.123, 593.3}.Abstand(Koordinate{-1.123, 593.3}))
}

func Test_Normalize(t *testing.T) {
	assert.InEpsilon(t, 1., Koordinate{132, 3}.Normalize().Length(), 0.00001)
	assert.Equal(t, Koordinate{}, Koordinate{math.NaN(), 2}.Normalize())
	assert.Equal(t, Koordinate{}, Koordinate{0, 0}.Normalize())
}

func Test_Mulitply(t *testing.T) {
	assert.Equal(t, Koordinate{3 * 0.5, 7 * 0.5}, Koordinate{3, 7}.Multiply(0.5))
	assert.Equal(t, Koordinate{0, 7 * 0.5}, Koordinate{math.NaN(), 7}.Multiply(0.5))
	assert.Equal(t, Koordinate{3 * 0.5, 0}, Koordinate{3, math.NaN()}.Multiply(0.5))
	assert.Equal(t, Koordinate{0, 0}, Koordinate{3, 7}.Multiply(math.NaN()))
}

func Test_Richtung(t *testing.T) {
	assert.Equal(t, Koordinate{-1, 0}, Koordinate{132, 42}.Richtung(Koordinate{42, 42}))
	assert.InEpsilon(t, 1.0, Koordinate{132, 42}.Richtung(Koordinate{-21, 420}).Length(), 0.00001)
}
