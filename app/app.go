// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"fmt"
	"github.com/EbitenPot/screenpot/res"
	"github.com/gen2brain/dlgs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/pkg/browser"
	"github.com/tidwall/gjson"
	"image/color"
	_ "image/png"
	"os"
	"runtime"
)

var (
	ImgCamera, _, _ = ebitenutil.NewImageFromFileSystem(res.Img, "camera.png")
	ImgRecord, _, _ = ebitenutil.NewImageFromFileSystem(res.Img, "record.png")
	ImgStop, _, _   = ebitenutil.NewImageFromFileSystem(res.Img, "stop.png")
	//ImgPause  = emoji.Image("⏸8️")
	ImgFolder, _, _ = ebitenutil.NewImageFromFileSystem(res.Img, "folder.png")
	ImgX, _, _      = ebitenutil.NewImageFromFileSystem(res.Img, "x.png")
	ImgInfo, _, _   = ebitenutil.NewImageFromFileSystem(res.Img, "info.png")
	ImgSite, _, _   = ebitenutil.NewImageFromFileSystem(res.Img, "website.png")
)

type Game struct {
	WW int
	WH int

	WX int
	WY int

	ConfigFile []byte

	PicType string
	RecType string
}

func NewGame() *Game {
	x, y := ebiten.ScreenSizeInFullscreen()
	cf, _ := os.ReadFile("config.json")
	pt := gjson.GetBytes(cf, "pictype").String()
	rt := gjson.GetBytes(cf, "rectype").String()
	return &Game{
		WW:         256,
		WH:         36,
		WX:         x - 296,
		WY:         y - 100,
		ConfigFile: cf,
		PicType:    pt,
		RecType:    rt,
	}
}

func (g *Game) PressButton(x, _ int) {
	if x < 36 { //shot
		ShotTest()
	} else if x < 72 { //rec
		dlgs.Info("Sorry", "Please wait the feature!")
	} else if x < 108 { //pause
		dlgs.Info("Sorry", "Please wait the feature!")
	} else if x < 108+36 { //folder

		dlgs.Info("Sorry", "Please wait the feature!")
	} else if x < 108+36+36 { //site
		err := browser.OpenURL("https://github.com/EbitenPot/screenpot")
		if err != nil {
			dlgs.Warning("Warning", fmt.Sprintf("openBrowser: unsupported operating system: %v", runtime.GOOS))
		}
	} else if x < 108+36+36+36 { //about

		dlgs.Info("About", `
Author: EldersJavas(eldersjavas@gmail.com)
Github: github.com/EbitenPot/screenpot
License: GPL v2.0		
		`)
	} else if x < 108+36+36+36+36 { //close
		g.Exit()
	} else {
		return
	}
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.PressButton(ebiten.CursorPosition())
	}
	if ebiten.IsWindowBeingClosed() {
		g.Exit()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 144, G: 139, B: 198, A: 0xff})
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.WW, g.WH
}

func (g *Game) Exit() {
	question, err := dlgs.Question("Exit", "Are you ready to exit?", false)
	if question != false || err != nil {
		os.Exit(0)
	}
}
