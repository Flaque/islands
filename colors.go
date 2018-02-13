package main

import (
	"image/color"
)

func colorFromNoise(noiseVal float64, seaLevel float64) color.RGBA {

	// Blue
	if noiseVal < sea.RealLevel(seaLevel) {
		return sea.color
	}

	// pinkish
	if noiseVal < sand.RealLevel(seaLevel) {
		return sand.color
	}

	// Brown
	if noiseVal < dirt.RealLevel(seaLevel) {
		return dirt.color
	}

	// Green
	if noiseVal < green.RealLevel(seaLevel) {
		return green.color
	}

	// White
	return peak.color
}

func blackFromNoise(noiseVal float64, seaLevel float64) color.RGBA {
	if noiseVal > seaLevel {
		return color.RGBA{0, 0, 0, 255}
	}
	return color.RGBA{255, 255, 255, 255}
}

func grayFromNoise(noiseVal float64) color.RGBA {

	return color.RGBA{uint8((noiseVal) * 255), uint8((noiseVal) * 255), uint8((noiseVal) * 255), 255}
}
