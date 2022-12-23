// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/EbitenPot/screenpot/app"
	"github.com/ebiten/emoji"
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
	icon := emoji.Image("📷")
	icons := append([]image.Image{}, icon.SubImage(icon.Bounds()))
	ebiten.SetWindowIcon(icons)
	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}
}
