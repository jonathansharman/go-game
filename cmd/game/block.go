package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
)

type Block struct {
	Object *resolv.ConvexPolygon
	Color  color.Color
}

func NewBlock(x, y, w, h float64, color color.Color) Block {
	return Block{
		Object: resolv.NewRectangleFromTopLeft(x, y, w, h),
		Color:  color,
	}
}

func (b Block) Draw(screen *ebiten.Image) {
	bounds := b.Object.Bounds()
	vector.DrawFilledRect(
		screen,
		float32(bounds.Min.X),
		float32(bounds.Min.Y),
		float32(bounds.Max.X-bounds.Min.X),
		float32(bounds.Max.Y-bounds.Min.Y),
		b.Color,
		false,
	)
}
