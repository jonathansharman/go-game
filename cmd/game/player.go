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
	playerJumpSpeed    = 15

	gravity  = 0.8
	friction = 0.85
)

type Player struct {
	Object   *resolv.ConvexPolygon
	Velocity resolv.Vector
	Image    *ebiten.Image
}

func NewPlayer() Player {
	playerImage, _, err := image.Decode(bytes.NewReader(nicheImage))
	if err != nil {
		log.Fatal(err)
	}

	return Player{
		Object: resolv.NewRectangleFromTopLeft(
			screenWidth/2-playerWidth/2,
			screenHeight/2-playerHeight/2,
			playerWidth,
			playerHeight,
		),
		Image: ebiten.NewImageFromImage(playerImage),
	}
}

func (p *Player) Update() {
	onGround := false
	if p.Velocity.Y >= 0 {
		p.Object.IntersectionTest(resolv.IntersectionTestSettings{
			TestAgainst: p.Object.SelectTouchingCells(1).FilterShapes(),
			OnIntersect: func(set resolv.IntersectionSet) bool {
				onGround = true
				p.Object.SetY(set.TopmostPoint().Y - playerHeight/2 + 1)
				p.Velocity.Y = 0
				return false
			},
		})
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
}

func (p Player) Draw(screen *ebiten.Image) {
	geoM := ebiten.GeoM{}
	geoM.Scale(
		playerWidth/float64(p.Image.Bounds().Dx()),
		playerHeight/float64(p.Image.Bounds().Dy()),
	)
	geoM.Translate(-playerWidth/2, -playerHeight/2)
	geoM.Rotate(p.Velocity.X * p.Velocity.Y / 100)
	geoM.Translate(p.Object.Position().X, p.Object.Position().Y-1)
	screen.DrawImage(p.Image, &ebiten.DrawImageOptions{GeoM: geoM})
}
