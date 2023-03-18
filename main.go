package m

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Vec struct {
	X, Y, Z float64
}

func Add(a, b Vec) Vec {
	return Vec{a.X + b.X, a.Y + b.Y, a.Z + b.Y}
}

func Sub(a, b Vec) Vec {
	return Vec{a.X - b.X, a.Y - b.Y, a.Z - b.Y}
}

func Div(v Vec, a float64) Vec {
	return Vec{v.X / a, v.Y / a, v.Z / a}
}

func Mul(v Vec, a float64) Vec {
	return Vec{v.X * a, v.Y * a, v.Z / a}
}

func Mod(a Vec) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
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

func DrawLine(screen *ebiten.Image, a, b Vec) {
	ebitenutil.DrawLine(screen, a.X, a.Y, b.X, b.Y, color.RGBA{255, 102, 204, 255})
}

type Rect struct {
	A, B, C, D Vec
}

func (r *Rect) Draw(screen *ebiten.Image) {
	DrawLine(screen, r.A, r.B)
	DrawLine(screen, r.B, r.C)
	DrawLine(screen, r.C, r.D)
	DrawLine(screen, r.D, r.A)
}

type Cube struct {
	p [8]Vec
}

func (c *Cube) Rotate(r Rotator) {
	for i := range c.p {
		c.p[i].Rotate(r)
	}
}
