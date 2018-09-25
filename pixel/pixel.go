package pixel

import (
	"github.com/aaronland/go-image-tools/util"
	"image"
	"image/color"
	"path/filepath"
	"sync"
)

type PixelFunc func(int, int, color.Color) (color.Color, error)

func MakeTransparentPixelFunc(r int, g int, b int) (PixelFunc, error) {

	f := func(x int, y int, c color.Color) (color.Color, error) {

		r32, g32, b32, _ := c.RGBA()

		r8 := uint8(r32 / 257)
		g8 := uint8(g32 / 257)
		b8 := uint8(b32 / 257)

		if r8 == uint8(r) && g8 == uint8(g) && b8 == uint8(g) {

			c = color.NRGBA{
				R: r8,
				G: g8,
				B: b8,
				A: 0,
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
