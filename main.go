package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/disintegration/imaging"
)

func main() {
	var file string
	var blurFactor float64
	var outFile string

	flag.StringVar(&file, "f", "/tmp/screenshot.png", "specify the image you want to blur")
	flag.Float64Var(&blurFactor, "b", 1.5, "specify the blur factor")
	flag.StringVar(&outFile, "o", "/tmp/screenshot.png", "specify output file")
	flag.Parse()

	start := time.Now()
	blur(file, outFile, blurFactor)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}

func blur(file string, outFile string, blurFactor float64) {
	src, err := imaging.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open image: %v\n", err)
		return
	}

	width := src.Bounds().Dx()
	height := src.Bounds().Dy()

	src = imaging.Resize(src, int(width/10), int(height/10), imaging.Lanczos)
	src = imaging.Blur(src, blurFactor)
	src = imaging.Resize(src, width, height, imaging.Lanczos)

	err = imaging.Save(src, outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to save image: %v", err)
	}
}
