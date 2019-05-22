package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"io"
	"image/png"
	"log"
	"os"
)

func main() {

	width := flag.Int("width", 1, "Image width, in pixels")
	height := flag.Int("height", 1, "Image height, in pixels")
	format := flag.String("format", "png", "Valid formats are: gif, png")
	stdout := flag.Bool("stdout", false, "Write to STDOUT")

	flag.Parse()

	switch *format {
	case "gif":
		// pass
	case "png":
		// pass
	default:
		log.Fatal("Invalid format")
	}

	im := image.NewNRGBA(image.Rect(0, 0, *width, *height))

	var wr io.Writer

	if *stdout {
		wr = os.Stdout
	} else {
		
		fname := fmt.Sprintf("empty.%s", *format)
		fh, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0644)
		
		if err != nil {
			log.Fatal(err)
		}

		defer fh.Close()
		wr = fh
	}

	var err error
	
	switch *format {
	case "gif":
		err = gif.Encode(wr, im, &gif.Options{})
	case "png":
		err = png.Encode(wr, im)
	default:
		// pass
	}

	if err != nil {
		log.Fatal(err)
	}

}
