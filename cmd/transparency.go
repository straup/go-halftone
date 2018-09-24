package main

import (
	"flag"
	"github.com/aaronland/go-image-tools/util"
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"
)

func main() {

	flag.Parse()

	for _, path := range flag.Args() {

		abs_path, err := filepath.Abs(path)

		if err != nil {
			log.Fatal(err)
		}

		im, _, err := util.DecodeImage(abs_path)

		if err != nil {
			log.Fatal(err)
		}

		bounds := im.Bounds()
		max := bounds.Max

		width := max.X
		height := max.Y

		new := image.NewNRGBA(image.Rect(0, 0, width, height))

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {

				c := im.At(x, y)
				r, g, b, _ := c.RGBA()
				
				r8 := uint8(r / 257)
				g8 := uint8(g / 257)
				b8 := uint8(b / 257)
				// a8 := uint8(a / 257)

				// log.Println("WHAT", x, y, r, b, g)

				if r8 == 255 && g8 == 255 && b8 == 255 {

					c = color.NRGBA{
						R: r8, 
						G: g8,
						B: b8,
						A: 0,
					}
				}
												
				new.Set(x, y, c)				
			}
		}

		fh, err := os.OpenFile("tr.png", os.O_RDWR|os.O_CREATE, 0644)

 		if err != nil {
			log.Fatal(err)
		}

		err = util.EncodeImage(new, "png", fh)

 		if err != nil {
			log.Fatal(err)
		}
		
	}
}
