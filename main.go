package main

import (
	"fmt"
	"image"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	_ "github.com/golang/freetype/truetype"
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

type MenuOption int

const (
	StartOption MenuOption = iota
	SettingsOption
	QuitOption
)

type Game struct {
	state   GameState
	menuSel MenuOption
	// menuFont *ebiten.Font
	menuFont font.Face
}

func (g *Game) Update() error {
	switch g.state {
	case MenuState:
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			g.menuSel++
			if g.menuSel > QuitOption {
				g.menuSel = StartOption
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			g.menuSel--
			if g.menuSel < StartOption {
				g.menuSel = QuitOption
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			switch g.menuSel {
			case StartOption:
				g.state = PlayState
			case SettingsOption:
				// TODO; but not right now btw 55555
			case QuitOption:
				return fmt.Errorf("user quit")
			}
		}
	case PlayState:
		// TODO
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case MenuState:
		msg := "Menu\n\n"
		if g.menuSel == StartOption {
			msg += "[START]  Settings  Quit"
		} else if g.menuSel == SettingsOption {
			msg += "Start  [SETTINGS]  Quit"
		} else {
			msg += "Start  Settings  [QUIT]"
		}

		ebitenutil.DebugPrint(screen, msg)
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

	// font initialize
	
	// fontFile, err := os.Open("path/to/font.ttf")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer fontFile.Close()

	// fontBytes, err := io.ReadAll(fontFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// font, err := truetype.Parse(fontBytes)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// menuFont, err := ebitenutil.NewFont(fontBytes, 12)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fontFile, err := os.Open("path/to/font.ttf")
	if err != nil {
		log.Fatal(err)
	}
	defer fontFile.Close()

	fontBytes, err := io.ReadAll(fontFile)
	if err != nil {
		log.Fatal(err)
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	menuFont, err := freetype.NewContext().NewFace(font, &freetype.Options{
		Size:    12,
		DPI:     72,
		Hinting: font.Hinting,
	})
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
		state:    MenuState,
		menuSel:  StartOption,
		menuFont: menuFont,
	}

	if err := ebiten.RunGame(game); err != nil {
		if err.Error() == "user quit" {
			fmt.Println("Goodbye!")
		}
	}

}
