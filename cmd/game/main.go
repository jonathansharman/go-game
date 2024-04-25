package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	ground       = 0.8 * screenHeight
)

type Game struct {
	Player Player
	Blocks []image.Rectangle
}

func (g *Game) Update() error {
	g.Player.Update(g.Blocks)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{G: 128, B: 255, A: 255})
	vector.DrawFilledRect(screen, 0, ground, screenWidth, screenHeight-ground, color.RGBA{G: 128, A: 255}, false)
	for _, block := range g.Blocks {
		vector.DrawFilledRect(screen, float32(block.Min.X), float32(block.Min.Y), float32(block.Max.X-block.Min.X), float32(block.Max.Y-block.Min.Y), color.Black, false)
	}

	g.Player.Draw(screen)

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, Learn Over Lunch!")
	if err := ebiten.RunGame(&Game{
		Player: NewPlayer(),
		Blocks: []image.Rectangle{
			{
				Min: image.Point{
					X: screenWidth / 2,
					Y: 0.6 * screenHeight,
				},
				Max: image.Point{
					X: screenWidth,
					Y: 0.6*screenHeight + 10,
				},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}
}
