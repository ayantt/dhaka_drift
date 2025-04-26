package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	X, Y  float64
	Img   *ebiten.Image
	Speed float64
}

func NewPlayer() *Player {
	img, _, err := ebitenutil.NewImageFromFile("asset/rickshaw.png")
	if err != nil {
		panic(err)
	}
	return &Player{
		X:     200,
		Y:     700,
		Img:   img,
		Speed: 5,
	}
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.X -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.X += p.Speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Img, op)
}

func (p *Player) Rect() (float64, float64, float64, float64) {
	w, h := float64(p.Img.Bounds().Dx()), float64(p.Img.Bounds().Dy())
	return p.X, p.Y, p.X + w, p.Y + h
}
