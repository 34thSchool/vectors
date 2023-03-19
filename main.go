package v

import (
	"math"
)

type Vec struct {
	X, Y, Z float64
}

func Add(a, b Vec) Vec {
	return Vec{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func Sub(a, b Vec) Vec {
	return Vec{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func Div(v Vec, a float64) Vec {
	return Vec{v.X / a, v.Y / a, v.Z / a}
}

func Mul(v Vec, a float64) Vec {
	return Vec{v.X * a, v.Y * a, v.Z * a}
}

func Mod(a Vec) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func Normalize(v Vec) Vec {
	return Vec{
		v.X / Mod(v),
		v.Y / Mod(v),
		v.Z / Mod(v),
	}
}

func Cross(a, b Vec) Vec {
	return Vec{
		a.Y*b.Z - b.Y*a.Z,
		a.Z*b.X - b.Z*a.X,
		a.X*b.Y - b.X*a.Y,
	}
}

func Dot(a, b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

type Rotator struct {
	X, Y, Z float64
}

func (v *Vec) RotateZ(rad float64) {
	x := v.X*math.Cos(rad) - v.Y*math.Sin(rad)
	y := v.X*math.Sin(rad) + v.Y*math.Cos(rad)
	v.X, v.Y = x, y
}

func (v *Vec) RotateY(rad float64) {
	x := v.X*math.Cos(rad) + v.Z*math.Sin(rad)
	z := v.X*-math.Sin(rad) + v.Z*math.Cos(rad)
	v.X, v.Z = x, z
}

func (v *Vec) RotateX(rad float64) {
	y := v.Y*math.Cos(rad) - v.Z*math.Sin(rad)
	z := v.Y*math.Sin(rad) + v.Z*math.Cos(rad)
	v.Y, v.Z = y, z
}

func (v *Vec) Rotate(rad Rotator) {
	v.RotateZ(rad.Z)
	v.RotateY(rad.Y)
	v.RotateX(rad.X)
}
