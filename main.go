package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "What should I name this game btw")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	icon, _, err := ebitenutil.NewImageFromFile("title-icon.png")
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("This game is just like a whiteboard")
	ebiten.SetWindowIcon([]image.Image{icon}) // Set the icon image

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
