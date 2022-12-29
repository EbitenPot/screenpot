// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 144, G: 139, B: 198, A: 0})
	op := new(ebiten.DrawImageOptions)
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(0, 0)
	screen.DrawImage(ImgCamera, op)
	op.GeoM.Translate(36, 0)
	screen.DrawImage(ImgRecord, op)
	op.GeoM.Translate(36, 0)
	screen.DrawImage(ImgStop, op)
	op.GeoM.Translate(36, 0)
	screen.DrawImage(ImgFolder, op)
	op.GeoM.Translate(36, 0)
	screen.DrawImage(ImgSite, op)
	op.GeoM.Translate(36, 0)
	screen.DrawImage(ImgInfo, op)
	op.GeoM.Translate(36, 0)
	screen.DrawImage(ImgX, op)
}
