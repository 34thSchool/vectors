package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Vector:
type vec struct {
	x, y float64
}

func add(a, b vec) vec {
	return vec{a.x + b.x, a.y + b.y}
}

func sub(a, b vec) vec {
	return vec{a.x - b.x, a.y - b.y}
}

func divide(v vec, a float64) vec {
	return vec{v.x / a, v.y / a}
}

func mod(a vec) float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y)
}

// Line:
func drawLine(screen *ebiten.Image, a, b vec) {
	ebitenutil.DrawLine(screen, a.x, a.y, b.x, b.y, color.RGBA{255, 102, 204, 255})
}

// Rectangle:
type rect struct {
	A, B, C, D vec
}

func (r *rect) Draw(screen *ebiten.Image) {
	drawLine(screen, r.A, r.B)
	drawLine(screen, r.B, r.C)
	drawLine(screen, r.C, r.D)
	drawLine(screen, r.D, r.A)
}
