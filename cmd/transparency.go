package main

import (
	"flag"
	"github.com/aaronland/go-image-tools/pixel"
	"github.com/aaronland/go-image-tools/util"
	"image/color"
	"log"
	"os"
)

func main() {

	flag.Parse()

	wh := color.RGBA{
		R: uint8(255),
		G: uint8(255),
		B: uint8(255),
		A: uint8(1),
	}

	cb, err := pixel.MakeTransparentPixelFunc(wh)

	if err != nil {
		log.Fatal(err)
	}

	for _, path := range flag.Args() {

		im, err := pixel.ProcessPath(path, cb)

		if err != nil {
			log.Fatal(err)
		}

		fh, err := os.OpenFile("tr3.png", os.O_RDWR|os.O_CREATE, 0644)

		if err != nil {
			log.Fatal(err)
		}

		err = util.EncodeImage(im, "png", fh)

		if err != nil {
			log.Fatal(err)
		}

	}
}
