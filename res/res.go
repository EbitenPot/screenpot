// Created by EldersJavas(EldersJavas&gmail.com)

package res

import (
	"embed"
	_ "embed"
	_ "image/png"
	"io/fs"
)

//go:embed img
var embimg embed.FS

var Img fs.FS

func init() {
	f, err := fs.Sub(embimg, "img")
	if err != nil {
		panic(err)
	}
	Img = f
}
