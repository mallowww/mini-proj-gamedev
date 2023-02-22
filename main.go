package main

import (
    "image"
    "log"
    "os"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "golang.org/x/image/draw"
)

const (
    screenWidth  = 800
    screenHeight = 600
)

type GameState int

const (
    MenuState GameState = iota
    PlayState
)

type Game struct {
    state GameState
}

func (g *Game) Update() error {
    // update the game state based on user input or other events
    switch g.state {
    case MenuState:
        // update the menu state
    case PlayState:
        // update the play state
    }

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    switch g.state {
    case MenuState:
        ebitenutil.DebugPrint(screen, "Menu")
    case PlayState:
        ebitenutil.DebugPrint(screen, "Play")
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return screenWidth, screenHeight
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

    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Game Title")
    ebiten.SetWindowIcon([]image.Image{resizedImg}) // Set the icon image

    // initialize the game state
    game := &Game{
        state: MenuState,
    }

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
