package logic

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	Player    *Player
	Obstacles []*Obstacle
	Score     int
	HighScore int
	GameOver  bool
}

func NewGame() *Game {
	highScore := loadHighScore()
	return &Game{
		Player:    NewPlayer(),
		Obstacles: []*Obstacle{NewObstacle()},
		Score:     0,
		HighScore: highScore,
		GameOver:  false,
	}
}

func (g *Game) Update() error {
	if g.GameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			*g = *NewGame()
		}
		return nil
	}

	g.Player.Update()

	for _, o := range g.Obstacles {
		o.Update()

		if checkCollision(g.Player, o) {
			g.GameOver = true
			if g.Score > g.HighScore {
				g.HighScore = g.Score
				saveHighScore(g.HighScore)
			}
		}

		if o.Y > 800 {
			o.Y = -100
			o.X = float64(50 + (randInt(0, 7) * 50))
			g.Score++
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)

	for _, o := range g.Obstacles {
		o.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Score: %d", g.Score), basicfont.Face7x13, 10, 20, colorWhite)
	text.Draw(screen, fmt.Sprintf("High Score: %d", g.HighScore), basicfont.Face7x13, 10, 40, colorWhite)

	if g.GameOver {
		text.Draw(screen, "Game Over!", basicfont.Face7x13, 180, 400, colorRed)
		text.Draw(screen, "Press SPACE to Restart", basicfont.Face7x13, 120, 430, colorRed)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 480, 800
}

func checkCollision(p *Player, o *Obstacle) bool {
	px1, py1, px2, py2 := p.Rect()
	ox1, oy1, ox2, oy2 := o.Rect()
	return px1 < ox2 && px2 > ox1 && py1 < oy2 && py2 > oy1
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func saveHighScore(score int) {
	_ = ioutil.WriteFile("highscore.txt", []byte(strconv.Itoa(score)), 0644)
}

func loadHighScore() int {
	data, err := ioutil.ReadFile("highscore.txt")
	if err != nil {
		return 0
	}
	score, _ := strconv.Atoi(string(data))
	return score
}

var colorWhite = colorRGBA(255, 255, 255, 255)
var colorRed = colorRGBA(255, 0, 0, 255)

func colorRGBA(r, g, b, a uint8) color.Color {
	return color.RGBA{R: r, G: g, B: b, A: a}
}
