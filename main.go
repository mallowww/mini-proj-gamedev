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
	once sync.Once

	owlImg      *ebiten.Image
	owlPosition ebiten.GeoM

	bgImg      *ebiten.Image
	bgPosition ebiten.GeoM
	bgSpeed    float64
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

		// thought: plan to remove the comments below after try something that works better then "must calculate the screen lines first to do it"
		// init owl postion, center of the screen
		_, screenHeight := ebiten.WindowSize()
		// screenWidth, screenHeight := ebiten.WindowSize()
		// owlWidth, owlHeight := owlImg.Size()
		// g.owlPosition.Translate(float64(screenWidth-owlWidth)/2, float64(screenHeight-owlHeight)/2)

		// init background position and speed
		bgWidth, _ := bgImg.Size()
		g.bgPosition.Translate(0, float64(screenHeight-bgWidth))
		g.bgSpeed = 3
	})

	// movements
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.owlPosition.Translate(0, -10)
	} else {
		g.owlPosition.Translate(0, 5)
	}
	g.owlPosition.Translate(3, 0)

	// background moving position
	g.bgPosition.Translate(-g.bgSpeed, 0)
	bgWidth, _ := g.bgImg.Size()
	if g.bgPosition.Element(0, 0) <= -float64(bgWidth) {
		g.bgPosition.Translate(2*float64(bgWidth), 0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := ebiten.WindowSize()

	// Draw the background image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.bgPosition.Element(0, 2), 0)
	screen.DrawImage(g.bgImg, op)

	// I will use the other pictures later but for now, let's settle it with the same image
	// Draw the second background image immediately after the first instance ends
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.bgPosition.Element(0, 2)+float64(g.bgImg.Bounds().Dx()), 0)
	screen.DrawImage(g.bgImg, op)

	// Draw the third background image
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.bgPosition.Element(0, 2)+2*float64(g.bgImg.Bounds().Dx()), 0)
	screen.DrawImage(g.bgImg, op)

	// Draw the owl image
	owlWidth, owlHeight := g.owlImg.Size()
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(screenWidth-owlWidth)/2, float64(screenHeight-owlHeight)/2)
	// Init owl postion
	op.GeoM.Translate(g.owlPosition.Element(0, 2), g.owlPosition.Element(1, 2)) 
	op.GeoM.Scale(.3, .3)
	screen.DrawImage(g.owlImg, op)

	// Update the background position
	g.bgPosition.Translate(-g.bgSpeed, 0)
	bgWidth := g.bgImg.Bounds().Dx()
	if g.bgPosition.Element(0, 2) <= -2*float64(bgWidth) {
		g.bgPosition.Translate(2*float64(bgWidth), 0)
	}

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
