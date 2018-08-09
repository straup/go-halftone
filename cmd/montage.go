package main

// cloned from https://github.com/lucentminds/montaginator

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"bufio"
)

func main() {
	// Start defining your application here.
	var opt jpeg.Options

	//fmt.Println( os.Args )

	// Check to make sure we've got two arguments.
	if( len( os.Args ) < 3 ) {
		// Missing one or more arguments.
		fmt.Println( "Useage: montaginator /path/to/image/dir /path/to/output.jpg" )
		log.Fatalf( "error: Missing one or more arguments." )
	}

	// Determines the path to the images directory.
	cPathImages := os.Args[1]

	// Determines the path to the output image file.
	cOutputFile := os.Args[2]

	// Verify images directory
	lExists, err := exists(cPathImages)

	if err != nil {
		// Failed to verify images directory.
		log.Fatalf("exists failed: %v", err)
	}

	if !lExists {
		// Images directory does not exist.
		log.Fatalf("error: Directory \"%v\" does not exist.", cPathImages)
	}

	// Scan the directory for the image files.
	aFiles, _ := ioutil.ReadDir(cPathImages)

	// Determine the width and height of the first image.
	// All other images should be the same.
	nWidthOne, nHeightOne := getImageDimension(cPathImages + "/" + aFiles[0].Name())

	// Determines the total height the final montage image will be.
	nHeightAll := nHeightOne * len(aFiles)

	// Determines the main montage image object.
	imgMontage := image.NewRGBA(image.Rect(0, 0, nWidthOne, nHeightAll))

	// Determines the color black.
	rgbaBlack := color.NRGBA{0, 0, 0, 0}

	// Draw the black color as the background to the montage image.
	draw.Draw(imgMontage, imgMontage.Bounds(), &image.Uniform{rgbaBlack}, image.ZP, draw.Src)

	// Create a new output file to write to.
	out, err := os.Create( cOutputFile )
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Determines the current y position multiplier for the next image.
	nY := 0

	// Loop over each image in the images directory.
	for _, oFile := range aFiles {

		// Open the next image.
		src, _, err := decode(cPathImages + "/" + oFile.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Paste the next image into the main montage image below the previous
		// image.
		draw.Draw(imgMontage, image.Rect( 0, nHeightOne*nY, nWidthOne, nHeightAll ), src, image.ZP, draw.Src)
		
		nY++
	} // /for()

	// put quality to x%
	opt.Quality = 90

	// Save the final montage image file.
	err = jpeg.Encode(out, imgMontage, &opt) 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out.Close()
} // /main()

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
} // /getImageDimension()

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func decode(filename string) (image.Image, string, error) {
  	f, err := os.Open(filename)
  	if err != nil {
  		return nil, "", err
  	}
  	defer f.Close()
  	return image.Decode(bufio.NewReader(f))
  }
