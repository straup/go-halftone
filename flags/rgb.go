package flags

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

type RGBColor []color.Color

func (c *RGBColor) String() string {
	return fmt.Sprintf("%v", *c)
}

func (c *RGBColor) Set(value string) error {

	rgb := strings.Split(value, ",")

	if len(rgb) != 3 {
		return errors.New("Invalid R,G,B count")
	}

	r, err := strconv.Atoi(rgb[0])

	if err != nil {
		return err
	}

	g, err := strconv.Atoi(rgb[1])

	if err != nil {
		return err
	}

	b, err := strconv.Atoi(rgb[2])

	if err != nil {
		return err
	}

	clr := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(1),
	}

	*c = append(*c, clr)
	return nil
}
