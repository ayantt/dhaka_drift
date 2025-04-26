package logic

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Obstacle struct {
	X, Y  float64
	Img   *ebiten.Image
	Speed float64
}

func NewObstacle() *Obstacle {
	img, _, err := ebitenutil.NewImageFromFile("asset/obstacle.png")
	if err != nil {
		panic(err)
	}
	x := randInt(50, 480)
	fmt.Println(x)
	return &Obstacle{
		X:     float64(x),
		Y:     0,
		Img:   img,
		Speed: 10,
	}
}

func (o *Obstacle) Update() {
	o.Y += o.Speed
}

func (o *Obstacle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.X, o.Y)
	screen.DrawImage(o.Img, op)
}

func (o *Obstacle) Rect() (float64, float64, float64, float64) {
	w, h := float64(o.Img.Bounds().Dx()), float64(o.Img.Bounds().Dy())
	return o.X, o.Y, o.X + w, o.Y + h
}
