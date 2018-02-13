package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"time"

	opensimplex "github.com/ojrac/opensimplex-go"
)

func save(name string, img *image.RGBA) {
	outputFile, err := os.Create(name)
	if err != nil {
		fmt.Println("Error saving file")
	}

	png.Encode(outputFile, img)
	outputFile.Close()
}

func blank(width int, height int) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	c := sea.color
	draw.Draw(m, m.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	return m
}

func noiseValues(width int, height int, density float64) [][]float64 {
	noise := opensimplex.NewWithSeed(int64(time.Now().Nanosecond()))

	// Make an empty 2d array
	vals := make([][]float64, height)
	for i := 0; i < height; i++ {
		vals[i] = make([]float64, width)
	}

	// Fill with floats
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			vals[x][y] = noise.Eval2(float64(x)/density, float64(y)/density)
		}
	}

	return vals
}

// Applys a smooth function by scaling the noise with a cone's slant height
func smoothNoiseViaCone(noise [][]float64, width int, height int, scalar float64) [][]float64 {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			rad := radius(float64(x), float64(y), float64(width), float64(height)) / float64(width*4)

			noise[x][y] /= rad * scalar
		}
	}

	return noise
}

func drawBiome(img *image.RGBA, noise [][]float64, seaLevel float64, biome Biome) {
	for x := 0; x < len(noise); x++ {
		for y := 0; y < len(noise[x]); y++ {

			noiseVal := noise[x][y]
			if biome.isHigherThan(noiseVal - seaLevel) {
				continue
			}

			img.SetRGBA(x, y+biome.bump, biome.color)
		}
	}
}

func main() {
	width := 64
	height := 64
	density := 15.0
	seaLevel := 0.0
	dropoff := 10.0

	img := blank(width, height)
	noise := noiseValues(width, height, density)
	noise = smoothNoiseViaCone(noise, width, height, dropoff)

	for _, biome := range biomes {
		drawBiome(img, noise, seaLevel, biome)
	}

	save("test.png", img)
}
