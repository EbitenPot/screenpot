// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/EbitenPot/screenpot/app"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

func main() {
	a := app.NewGame()
	ebiten.SetWindowSize(a.WW, a.WH)
	ebiten.SetWindowTitle("screenpot")
	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetWindowPosition(a.WX, a.WY)
	ebiten.SetTPS(30)
	ebiten.SetRunnableOnUnfocused(true)
	icons := append([]image.Image{}, app.ImgCamera)
	ebiten.SetWindowIcon(icons)
	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}
}
