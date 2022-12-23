// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/gen2brain/dlgs"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"os"
)

type Game struct {
	WW int
	WH int

	WX int
	WY int
}

func NewGame() *Game {
	x, y := ebiten.ScreenSizeInFullscreen()
	return &Game{
		WH: 64,
		WW: 256,
		WX: x - 256,
		WY: y - 100,
	}
}

func (g *Game) Update() error {
	g.Exit()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 144, G: 139, B: 198, A: 0xff})

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.WW, g.WH
}

func (g *Game) Exit() {
	if ebiten.IsWindowBeingClosed() {
		ebiten.SetFullscreen(true)
		/*err := beeep.Notify("Title", "Message body", "assets/information.png")
		if err != nil {
			panic(err)
		}*/
		question, err := dlgs.Question("Exit", "Are you ready to exit?", false)
		if question != false || err != nil {
			os.Exit(0)
		}
	}
}
