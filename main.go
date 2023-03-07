package main

import (
	"image"
	"log"
	"sync"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	bgImg *ebiten.Image
	img   *ebiten.Image
	once  sync.Once
}

func (g *Game) Update() error {
	g.once.Do(func() {
		backgroundImage := "./assets/img/example-background.png"
		bgImg, _, err := ebitenutil.NewImageFromFile(backgroundImage)
		if err != nil {
			log.Fatal(err)
		}
		g.bgImg = bgImg

		characterImage := "./assets/img/cutey-owl.png"
		img, _, err := ebitenutil.NewImageFromFile(characterImage)
		if err != nil {
			log.Fatal(err)
		}
		g.img = img
	})
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bgImg, nil)
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(.3, .3)
	op.GeoM.Translate(300, 300)
	screen.DrawImage(g.img, &op)
}

func (g *Game) Layout(width, height int) (int, int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowTitle("กลายเป็น flappybird game type ไปซะเอง (55555555)")

	titleIcon := "./assets/img/title-icon.png"
	iconImage, _, err := ebitenutil.NewImageFromFile(titleIcon)
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowIcon([]image.Image{iconImage})

	ebiten.SetWindowSize(1280, 720)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
