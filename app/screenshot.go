// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/kbinani/screenshot"
	"image/png"
	"log"
	"os"
)

type ScreenShot struct {
	Path      string
	PicType   string
	ScreenNum int8
}

func ShotTest() {
	//n := screenshot.NumActiveDisplays()

	//for i := 0; i < n; i++ {
	bounds := screenshot.GetDisplayBounds(0)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%d_%dx%d.png", 0, bounds.Dx(), bounds.Dy())
	file, _ := os.Create(fileName)
	defer file.Close()
	//png.Encode(file, img)
	err = png.Encode(file, img)
	if err != nil {
		log.Println("cannot encode png")
		return
	}
	fmt.Printf("#%d : %v \"%s\"\n", 0, bounds, fileName)
	//}
	err = beeep.Notify("Screen Pot", "Screen has successful saved!", "assets/information.png")
	if err != nil {
		log.Println("cannot send notice")
		return
	}
}
