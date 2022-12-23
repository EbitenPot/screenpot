// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
)

func ShotTest() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		//png.Encode(file, img)
		png.Encode(file, img)
		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
	beeep.Notify("screenpot", "Screen has saved.", "assets/information.png")
}
