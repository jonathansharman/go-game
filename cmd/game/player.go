package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"
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
	Object   *resolv.Circle
	Velocity resolv.Vector
	Image    *ebiten.Image
}

func NewPlayer() Player {
	playerImage, _, err := image.Decode(bytes.NewReader(nicheImage))
	if err != nil {
		log.Fatal(err)
	}

	return Player{
		Object: resolv.NewCircle(screenWidth/2, screenHeight/2, playerWidth/2),
		Image:  ebiten.NewImageFromImage(playerImage),
	}
}

func (p Player) Collider() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: int(p.Object.Position().X - playerWidth/2),
			Y: int(p.Object.Position().Y - playerHeight),
		},
		Max: image.Point{
			X: int(p.Object.Position().X + playerWidth/2),
			Y: int(p.Object.Position().Y + 1),
		},
	}
}

func (p *Player) Update(blocks []image.Rectangle) {
	onGround := p.Object.Position().Y >= ground

	collider := p.Collider()
	for _, block := range blocks {
		if p.Velocity.Y >= 0 && collider.Overlaps(block) {
			onGround = true
			p.Object.SetY(float64(block.Min.Y))
			p.Velocity.Y = 0
		}
	}

	if onGround && p.Velocity.Y >= 0 {
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			p.Velocity.X -= playerAcceleration
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			p.Velocity.X += playerAcceleration
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			p.Velocity.Y = -playerJumpSpeed
		}
	}

	p.Object.MoveVec(p.Velocity)
	if onGround {
		p.Velocity.X *= friction
	} else {
		p.Velocity.Y += gravity
	}
	if p.Object.Position().Y > ground {
		p.Object.SetY(ground)
		p.Velocity.Y = 0
	}
}

func (p Player) Draw(screen *ebiten.Image) {
	geoM := ebiten.GeoM{}
	geoM.Scale(
		playerWidth/float64(p.Image.Bounds().Dx()),
		playerHeight/float64(p.Image.Bounds().Dy()),
	)
	geoM.Translate(-playerWidth/2, -playerHeight/2)
	geoM.Rotate(p.Velocity.X * p.Velocity.Y / 100)
	geoM.Translate(p.Object.Position().X, p.Object.Position().Y-playerHeight/2)
	screen.DrawImage(p.Image, &ebiten.DrawImageOptions{GeoM: geoM})
}
