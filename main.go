package maths

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
	return Vec{a.X + b.X, a.Y + b.Y, 0}
}

func Sub(a, b Vec) Vec {
	return Vec{a.X - b.X, a.Y - b.Y, 0}
}

func Divide(v Vec, a float64) Vec {
	return Vec{v.X / a, v.Y / a, 0}
}

func Mod(a Vec) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

type Rotator struct {
	X, Y, Z float64
}

func (v *Vec) Rotate(rad Rotator) {
	// Rotation around Z
	v.X = v.X*math.Cos(rad.Z) - v.Y*math.Sin(rad.Z)
	v.Y = v.X*math.Sin(rad.Z) + v.Y*math.Cos(rad.Z)

	// Rotation around Y
	v.X = v.X*math.Cos(rad.Y) - v.Z*math.Sin(rad.Y)
	v.Z = v.X*math.Sin(rad.Y) + v.Z*math.Cos(rad.Y)

	// Rotation around X
	v.Y = v.Y*math.Cos(rad.X) + v.Z*math.Sin(rad.X)
	v.Z = -v.Y*math.Sin(rad.X) + v.Z*math.Cos(rad.X)
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
