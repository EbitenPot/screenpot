// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.PressButton(ebiten.CursorPosition())
	}
	if ebiten.IsWindowBeingClosed() {
		g.Exit()
	}

	return nil
}
