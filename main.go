package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/draw"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "What should I name this game btw")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	iconFile, err := os.Open("assets/img/title-icon.png")
	if err != nil {
		log.Fatal(err)
	}
	defer iconFile.Close()

	iconImg, _, err := image.Decode(iconFile)
	if err != nil {
		log.Fatal(err)
	}

	// resize title icon config
	const iconSize = 128
	resizedImg := image.NewRGBA(image.Rect(0, 0, iconSize, iconSize))
	draw.NearestNeighbor.Scale(resizedImg, resizedImg.Bounds(), iconImg, iconImg.Bounds(), draw.Over, nil)

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("This game is just like a whiteboard")
	ebiten.SetWindowIcon([]image.Image{resizedImg}) // Set the icon image

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
