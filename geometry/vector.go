package geometry

import (
	"errors"
	"math"
)

type Vector struct {
	x float64
	y float64
	z float64
}

func (vo *Vector) Clone() *Vector {
	v := new(Vector)

	v.x = vo.x
	v.y = vo.y
	v.z = vo.z

	return v
}

func (v *Vector) Add(v2 *Vector) *Vector {
	v3 := new(Vector)

	v3.x = v.x + v2.x
	v3.y = v.y + v2.y
	v3.z = v.z + v2.z

	return v3
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Vector) DotProduct(v2 *Vector) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v *Vector) CrossProduct(v2 *Vector) *Vector {
	v3 := new(Vector)

	v3.x = v.y*v2.z - v.z*v2.y
	v3.y = v.z*v2.x - v.x*v2.z
	v3.z = v.x*v2.y - v.y*v2.x

	return v3
}

func (v *Vector) Normalize() (*Vector, error) {
	magnitude := v.Magnitude()

	if magnitude > 0 {
		return computeNormalisation(v, magnitude), nil
	} else {
		return v, errors.New("Vector: could not normalize Vector")
	}
}

func computeNormalisation(v *Vector, magnitude float64) *Vector {
	v2 := new(Vector)

	invert := 1 / magnitude

	v2.x = v.x * invert
	v2.y = v.y * invert
	v2.z = v.z * invert

	return v2
}
