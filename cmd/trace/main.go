package main

import (
	"flag"
	"github.com/aaronland/go-image-tools/util"
	"github.com/dennwc/gotrace"
	_ "github.com/dennwc/gotrace/bindings"
	"image"
	"image/color"
	"log"
	"os"
)

func rgbaToGray(img image.Image) *image.Gray {

	bounds := img.Bounds()
	gray := image.NewGray(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	
	return gray
}

func blackOrWhite(g color.Gray) color.Color {

	if g.Y < 127 {
		return color.NRGBA{
			R: uint8(255),
			G: uint8(255),
			B: uint8(255),
			A: uint8(255),
		}
	}

	return color.NRGBA{
		R: uint8(0),
		G: uint8(0),
		B: uint8(0),
		A: 0,
	}

}

func ThresholdDither(gray *image.Gray) image.Image {

	bounds := gray.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	dithered := image.NewRGBA(bounds)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			var c = blackOrWhite(gray.GrayAt(i, j))
			dithered.Set(i, j, c)
		}
	}
	
	return dithered
}

func main() {

	source := flag.String("source", "", "...")

	flag.Parse()

	im, _, err := util.DecodeImage(*source)

	if err != nil {
		log.Fatal(err)
	}

	var gray = rgbaToGray(im)
	gray2 := ThresholdDither(gray)

	bm := gotrace.NewBitmapFromImage(gray2, nil)

	params := gotrace.Defaults
	params.TurdSize = 5
	params.OptTolerance = 0.5
	
	paths, _ := gotrace.Trace(bm, &params)

	gotrace.WriteSvg(os.Stdout, gray.Bounds(), paths, "")
}
