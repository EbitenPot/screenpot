// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"fmt"
	"github.com/EbitenPot/screenpot/res"
	"github.com/gen2brain/dlgs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pkg/browser"
	"github.com/tidwall/gjson"
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

	PicType       string
	RecType       string
	FPath         string
	Defaultscreen int8
}

func NewGame() *Game {
	x, y := ebiten.ScreenSizeInFullscreen()

	/*
		for other language speakers:
		It was just a joke ↓↓↓
		If there are any problems,
		you can change to ASCII characters.
	*/
	//goland:noinspection NonAsciiCharacters
	图片格式 := "png"
	//goland:noinspection NonAsciiCharacters
	视频格式 := "mp4"
	//goland:noinspection NonAsciiCharacters
	文件夹, _ := os.Getwd()

	var ds int8
	cf, err := os.ReadFile("config.json")
	if err != nil {
		图片格式 = gjson.GetBytes(cf, "pictype").String()
		视频格式 = gjson.GetBytes(cf, "rectype").String()
		文件夹 = gjson.GetBytes(cf, "savefolder").String()
		ds = int8(gjson.GetBytes(cf, "defaultscreen").Int())

	}
	return &Game{
		WW:            256,
		WH:            36,
		WX:            x - 296,
		WY:            y - 100,
		ConfigFile:    cf,
		PicType:       图片格式,
		RecType:       视频格式,
		FPath:         文件夹 + "output/",
		Defaultscreen: ds,
	}
}

func (g *Game) PressButton(x, _ int) {
	if x < 36 { //shot
		ShotTest()
	} else if x < 36*2 { //rec
		//trt ,_:= NewTranslation(TranslationConfig{}).T("Please wait the features!", "en", "zh")
		//ttest, _ := trt.T("Please wait the features!", "en", "zh")
		//_, err := dlgs.Info("Sorry", "Please wait the features!")
		err := Rectest()
		if err != nil {
			return
		}
	} else if x < 36*3 { //pause
		_, err := dlgs.Info("Sorry", "Please wait the features!")
		if err != nil {
			return
		}
	} else if x < 36*4 { //folder
		err := browser.OpenFile(g.FPath)
		if err != nil {
			_, err := dlgs.Warning("Warning", fmt.Sprintf("openBrowser: unsupported operating system: %v", runtime.GOOS))
			if err != nil {
				return
			}
		}
	} else if x < 36*5 { //site
		err := browser.OpenURL("https://github.com/EbitenPot/screenpot")
		if err != nil {
			_, err := dlgs.Warning("Warning", fmt.Sprintf("openBrowser: unsupported operating system: %v", runtime.GOOS))
			if err != nil {
				return
			}
		}
	} else if x < 36*6 { //about
		_, err := dlgs.Info("About", `Author: EldersJavas(eldersjavas@gmail.com)
Github: github.com/EbitenPot/screenpot
License: GPL v2.0
Powered by Ebitengine v2.4
Enjoy!`)
		if err != nil {
			return
		}
	} else if x < 36*7 { //close
		g.Exit()
	} else {
		return
	}
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.WW, g.WH
}

func (g *Game) Exit() {
	question, err := dlgs.Question("Exit", "Are you ready to exit?", false)
	if question != false || err != nil {
		os.Exit(0)
	}
}
