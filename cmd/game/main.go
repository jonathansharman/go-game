package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
)

const (
	screenWidth  = 640
	screenHeight = 480
	ground       = 0.8 * screenHeight
)

type Game struct {
	Player Player
	Blocks []Block
}

func NewGame() *Game {
	player := NewPlayer()
	blocks := []Block{
		NewBlock(
			screenWidth/2,
			0.6*screenHeight,
			screenWidth/4,
			10,
			color.Black,
		),
		NewBlock(
			screenWidth/4,
			0.3*screenHeight,
			screenWidth/4,
			10,
			color.RGBA{R: 255, A: 255},
		),
	}

	space := resolv.NewSpace(screenWidth, screenHeight, playerWidth, playerHeight)
	space.Add(player.Object)
	for _, block := range blocks {
		space.Add(block.Object)
	}

	return &Game{
		Player: player,
		Blocks: blocks,
	}
}

func (g *Game) Update() error {
	g.Player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{G: 128, B: 255, A: 255})
	vector.DrawFilledRect(screen, 0, ground, screenWidth, screenHeight-ground, color.RGBA{G: 128, A: 255}, false)
	for _, block := range g.Blocks {
		block.Draw(screen)
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
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
