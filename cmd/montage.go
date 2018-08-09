package main

// THIS DOES NOT WORK YET

import (
	"flag"
	"github.com/aaronland/go-image-tools/util"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	out := flag.String("out", "", "...")

	flag.Parse()

	ext := filepath.Ext(*out)
	format := strings.Replace(ext, ".", "", -1)

	wr, err := os.OpenFile(*out, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal()
	}

	files := flag.Args()
	count := len(files)

	images := make([]image.Image, count)

	w := 0
	h := 0

	for i, path := range files {

		im, _, err := util.DecodeImage(path)

		if err != nil {
			log.Fatal(err)
		}

		w += util.Width(im)
		h += util.Height(im)

		images[i] = im
	}

	log.Println("FINAL", w, h)
	
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	rgbaBlack := color.NRGBA{0, 0, 0, 0}

	draw.Draw(m, m.Bounds(), &image.Uniform{rgbaBlack}, image.ZP, draw.Src)

	x := 0
	y := 0

	for _, im := range images {

		w := util.Width(im)
		h := util.Height(im)

		log.Println("DRAW AT", x, y, w, h)
		
		draw.Draw(m, image.Rect(x, y, w, h), im, image.ZP, draw.Src)

		x += w
		y += h
	}

	err = util.EncodeImage(m, format, wr)

	if err != nil {
		log.Fatal(err)
	}

	wr.Close()
}
