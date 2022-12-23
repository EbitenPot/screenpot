// Created by EldersJavas(EldersJavas&gmail.com)

package app

import "github.com/kbinani/screenshot"

func ScreenNumber() int {
	return screenshot.NumActiveDisplays()
}
