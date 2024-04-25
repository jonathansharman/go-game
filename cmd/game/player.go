package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//go:embed niche.png
var nicheImage []byte

const (
	playerWidth  = 60
	playerHeight = 60

	playerAcceleration = 1
	playerJumpSpeed    = 10

	gravity  = 0.4
	friction = 0.85
)

type Player struct {
	X     float64
	Y     float64
	VX    float64
	VY    float64
	Image *ebiten.Image
}

func NewPlayer() Player {
	playerImage, _, err := image.Decode(bytes.NewReader(nicheImage))
	if err != nil {
		log.Fatal(err)
	}

	return Player{
		X:     screenWidth / 2,
		Y:     screenHeight / 2,
		Image: ebiten.NewImageFromImage(playerImage),
	}
}

func (p Player) Collider() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: int(p.X - playerWidth/2),
			Y: int(p.Y - playerHeight),
		},
		Max: image.Point{
			X: int(p.X + playerWidth/2),
			Y: int(p.Y + 1),
		},
	}
}

func (p *Player) Update(blocks []image.Rectangle) {
	onGround := p.Y >= ground

	collider := p.Collider()
	for _, block := range blocks {
		if p.VY >= 0 && collider.Overlaps(block) {
			onGround = true
			p.Y = float64(block.Min.Y)
			p.VY = 0
		}
	}

	if onGround && p.VY >= 0 {
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			p.VX -= playerAcceleration
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			p.VX += playerAcceleration
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			p.VY = -playerJumpSpeed
		}
	}

	p.X += p.VX
	p.Y += p.VY
	if onGround {
		p.VX *= friction
	} else {
		p.VY += gravity
	}
	if p.Y > ground {
		p.Y = ground
		p.VY = 0
	}
}

func (p Player) Draw(screen *ebiten.Image) {
	geoM := ebiten.GeoM{}
	geoM.Scale(
		playerWidth/float64(p.Image.Bounds().Dx()),
		playerHeight/float64(p.Image.Bounds().Dy()),
	)
	geoM.Translate(-playerWidth/2, -playerHeight/2)
	geoM.Rotate(p.VX * p.VY / 100)
	geoM.Translate(p.X, p.Y-playerHeight/2)
	screen.DrawImage(p.Image, &ebiten.DrawImageOptions{GeoM: geoM})
}
