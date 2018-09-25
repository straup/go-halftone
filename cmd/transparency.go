package main

import (
	"flag"
	"github.com/aaronland/go-image-tools/flags"
	"github.com/aaronland/go-image-tools/pixel"
	"github.com/aaronland/go-image-tools/util"
	"log"
	"os"
)

func main() {

	var colors flags.RGBColor
	flag.Var(&colors, "color", "...")

	flag.Parse()

	cb, err := pixel.MakeTransparentPixelFunc(colors...)

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
