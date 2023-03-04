package main

import (
	"image/color"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	sceneX = 320
	sceneY = 240

	move = 2
)

type Game struct {
	enemies []*Enemy
	player  *Player
	score   int
}

type Enemy struct {
	x float64
	y float64
}

type Player struct {
	x float64
	y float64
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.x -= move
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.x += move
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.y -= move
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.y += move
	}

	if rand.Float64() < 0.05 {
		g.enemies = append(g.enemies, &Enemy{
			x: rand.Float64() * sceneX,
			y: -10,
		})
	}

	for _, e := range g.enemies {
		e.y += 2
	}

	for _, e := range g.enemies {
		if e.x > g.player.x-16 && e.x < g.player.x+16 &&
			e.y > g.player.y-16 && e.y < g.player.y+16 {
			return nil
		}
	}
	g.score++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, g.player.x, g.player.y, 32, 32, color.RGBA{255, 0, 0, 255})

	for _, e := range g.enemies {
		ebitenutil.DrawRect(screen, e.x, e.y, 16, 16, color.RGBA{0, 0, 255, 255})
	}
	ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(g.score))
}

func (g *Game) Layout(outX, outY int) (int, int) {
	return sceneX, sceneY
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := &Game{
		player: &Player{
			x: sceneX / 2,
			y: sceneY - 40,
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}