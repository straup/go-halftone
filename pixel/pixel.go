package pixel

import (
	"github.com/aaronland/go-image-tools/util"
	"image"
	"image/color"
	"path/filepath"
	"sync"
)

type PixelFunc func(int, int, color.Color) (color.Color, error)

func MakeReplacePixelFunc(match_c color.Color, replace_c color.Color) (PixelFunc, error) {

     f := func(x int, y int, c color.Color) (color.Color, error) {

		cr, cg, cb, _ := c.RGBA()
		mr, mg, mb, _ := match_c.RGBA()

		if cr == mr && cg == mg && cb == mb {

			c = replace_c
		}

		return c, nil
     }

     return f, nil
}

func MakeTransparentPixelFunc(matches ...color.Color) (PixelFunc, error) {

	f := func(x int, y int, c color.Color) (color.Color, error) {

		cr, cg, cb, _ := c.RGBA()

		for _, m := range matches {

			mr, mg, mb, _ := m.RGBA()

			if cr == mr && cg == mg && cb == mb {

				c = color.NRGBA{
					R: uint8(cr / 257),
					G: uint8(cg / 257),
					B: uint8(cg / 257),
					A: 0,
				}

				break
			}
		}

		return c, nil
	}

	return f, nil
}

func ProcessPath(path string, cb PixelFunc) (image.Image, error) {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	im, _, err := util.DecodeImage(abs_path)

	if err != nil {
		return nil, err
	}

	return ProcessImage(im, cb)
}

func ProcessImage(im image.Image, cb PixelFunc) (image.Image, error) {

	bounds := im.Bounds()
	max := bounds.Max

	width := max.X
	height := max.Y

	pr := image.NewNRGBA(image.Rect(0, 0, width, height))

	wg := new(sync.WaitGroup)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			wg.Add(1)

			go func(x int, y int, c color.Color) {

				defer wg.Done()

				new_c, _ := cb(x, y, c)
				pr.Set(x, y, new_c)

			}(x, y, im.At(x, y))
		}
	}

	wg.Wait()

	return pr, nil
}
