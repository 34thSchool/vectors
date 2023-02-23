package maths

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Vector:
type Vec struct {
	x, y, z float64
}

func Add(a, b Vec) Vec {
	return Vec{a.x + b.x, a.y + b.y, 0}
}

func Sub(a, b Vec) Vec {
	return Vec{a.x - b.x, a.y - b.y, 0}
}

func Divide(v Vec, a float64) Vec {
	return Vec{v.x / a, v.y / a, 0}
}

func Mod(a Vec) float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y)
}

// Line:
func DrawLine(screen *ebiten.Image, a, b Vec) {
	ebitenutil.DrawLine(screen, a.x, a.y, b.x, b.y, color.RGBA{255, 102, 204, 255})
}

// Rectangle:
type Rect struct {
	A, B, C, D Vec
}

func (r *Rect) Draw(screen *ebiten.Image) {
	DrawLine(screen, r.A, r.B)
	DrawLine(screen, r.B, r.C)
	DrawLine(screen, r.C, r.D)
	DrawLine(screen, r.D, r.A)
}
