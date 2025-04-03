package main

import (
	"image/color"
	"log"
	"math/rand"
	"slices"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	Space  *resolv.Space
	Player Player
	Blocks []Block
	Timer  *time.Ticker
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
	}

	space := resolv.NewSpace(screenWidth, screenHeight, playerWidth, playerHeight)
	space.Add(player.Object)
	for _, block := range blocks {
		space.Add(block.Object)
	}

	return &Game{
		Space:  space,
		Player: player,
		Blocks: blocks,
		Timer:  time.NewTicker(3 * time.Second),
	}
}

func (g *Game) Update() error {
	g.Player.Update()

	for i := range g.Blocks {
		block := &g.Blocks[i]
		block.Update()
		if block.Dead() {
			g.Space.Remove(block.Object)
		}
	}

	g.Blocks = slices.DeleteFunc(g.Blocks, Block.Dead)

	select {
	case <-g.Timer.C:
		block := NewBlock(
			screenWidth,
			rand.Float64()*screenHeight,
			screenWidth/8,
			10,
			color.RGBA{
				R: uint8(rand.Intn(256)),
				G: uint8(rand.Intn(256)),
				B: uint8(rand.Intn(256)),
				A: 255,
			},
		)
		g.Space.Add(block.Object)
		g.Blocks = append(g.Blocks, block)
	default:
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{G: 128, B: 255, A: 255})

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
