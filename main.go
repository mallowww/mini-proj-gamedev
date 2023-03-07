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
	once  sync.Once

	owlImg      *ebiten.Image
	owlPosition ebiten.GeoM
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
		owlImg, _, err := ebitenutil.NewImageFromFile(characterImage)
		if err != nil {
			log.Fatal(err)
		}
		g.owlImg = owlImg

		// init owl postion, center of the screen
		screenWidth, screenHeight := ebiten.WindowSize()
		owlWidth, owlHeight := owlImg.Size()
		g.owlPosition.Translate(float64(screenWidth-owlWidth)/2, float64(screenHeight-owlHeight)/2)
	})

	// movements
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.owlPosition.Translate(0, -10)
	} else {
		g.owlPosition.Translate(0, 5)
	}
	g.owlPosition.Translate(3, 0)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bgImg, nil)
	op := ebiten.DrawImageOptions{}
	op.GeoM = g.owlPosition
	op.GeoM.Scale(.3, .3)
	op.GeoM.Translate(300, 300)
	screen.DrawImage(g.owlImg, &op)
}

func (g *Game) Layout(width, height int) (int, int) {
	return 1280, 720
}

func main() {
	windowTitle := "กลายเป็น flappybird game type ไปซะเอง (55555555)"
	ebiten.SetWindowTitle(windowTitle)

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
