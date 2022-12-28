// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/gen2brain/x264-go"
	"github.com/kbinani/screenshot"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Rectest() error {
	file, err := os.Create("screen.mp4")
	if err != nil {
		return err
	}

	bounds := screenshot.GetDisplayBounds(0)

	opts := &x264.Options{

		Width:     bounds.Dx(),
		Height:    bounds.Dy(),
		FrameRate: 3,
		Tune:      "fastdecode",
		Preset:    "ultrafast",
		Profile:   "main",
		LogLevel:  x264.LogError,
	}

	enc, err := x264.NewEncoder(file, opts)
	if err != nil {
		return err
	}

	defer enc.Close()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(time.Second / time.Duration(3))

	start := time.Now()
	frame := 0

	for range ticker.C {
		select {
		case <-s:
			err := enc.Flush()
			if err != nil {
				return err
			}

			err = file.Close()
			if err != nil {
				return err
			}

			os.Exit(0)
		default:
			frame++
			log.Printf("frame: %v", frame)
			img, err := screenshot.CaptureRect(bounds)
			if err != nil {
				return err
			}
			/*			var twr io.Writer
						var trr io.Reader
						err = jpeg.Encode(twr, img, &jpeg.Options{Quality: 10})
						if err != nil {
							return err
						}
						_, err = io.Copy(twr, trr)
						if err != nil {
							return err
						}
						dec, err := jpeg.Decode(trr)
						if err != nil {
							return err
						}*/
			err = enc.Encode(img)
			if err != nil {
				return err
			}
			log.Printf("t: %v", time.Since(start))
			start = time.Now()
		}
	}
	return nil
}
