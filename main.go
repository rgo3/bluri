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

	flag.StringVar(&file, "f", "/tmp/screenshot.png", "specify the image you want to blur")
	flag.Float64Var(&blurFactor, "b", 1.5, "specify the blur factor")

	start := time.Now()
	blur(file, blurFactor)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}

func blur(file string, blurFactor float64) {
	src, err := imaging.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open image: %v", err)
	}

	src = imaging.Resize(src, 300, 300, imaging.Lanczos)

	src = imaging.Blur(src, blurFactor)

	src = imaging.Resize(src, 1920, 1080, imaging.Lanczos)

	err = imaging.Save(src, file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to save image: %v", err)
	}
}
