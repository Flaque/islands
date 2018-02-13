package main

import (
	"fmt"
	"math"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func center(width float64, height float64) (float64, float64) {
	return width / 2, height / 2
}

func radius(x float64, y float64, width float64, height float64) float64 {
	cx, cy := center(width, height)
	return math.Sqrt(math.Pow(x-cx, 2.0) + math.Pow(y-cy, 2.0))
}

func slantHeight(r float64, h float64) float64 {
	return math.Sqrt(math.Pow(r, 2) + math.Pow(h, 2))
}

func scaledSlantHeight(r float64, h float64) float64 {
	fmt.Println(r, h)
	return slantHeight(r, h) / h
}
