// Created by EldersJavas(EldersJavas&gmail.com)

package res

import (
	"embed"
	"io/fs"
)

var (
	//go:embed img
	embimg embed.FS
)

var Img fs.FS

func init() {
	f, err := fs.Sub(embimg, "img")
	if err != nil {
		panic(err)
	}
	Img = f
}
