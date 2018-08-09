package util

import (
	"image"
)

func Width(im image.Image) int {
	b := im.Bounds()
	return b.Max.X
}

func Height(im image.Image) int {
	b := im.Bounds()
	return b.Max.Y
}
