package main

import "image/color"

type Biome struct {
	level float64
	name  string
	color color.RGBA
	bump  int
}

func (b Biome) RealLevel(seaLevel float64) float64 {
	return b.level - seaLevel
}

func (b Biome) isHigherThan(level float64) bool {
	return level < b.level
}

var biomes = []Biome{
	sea, sand, dirt, green, darkGreen, stone, peak,
}

var sea Biome = Biome{
	0.0,
	"sea",
	color.RGBA{32, 76, 101, 255},
	0,
}

var sand Biome = Biome{
	0.1,
	"sand",
	color.RGBA{132, 120, 84, 255},
	0,
}

var dirt Biome = Biome{
	0.3,
	"dirt",
	color.RGBA{96, 70, 59, 255},
	-2,
}

var green Biome = Biome{
	0.5,
	"green",
	color.RGBA{10, 135, 84, 255},
	-3,
}

var darkGreen Biome = Biome{
	0.7,
	"darkGreen",
	color.RGBA{20, 145, 94, 255},
	-4,
}

var stone Biome = Biome{
	0.9,
	"stone",
	color.RGBA{93, 101, 97, 255},
	-5,
}

var peak Biome = Biome{
	1.1,
	"peak",
	color.RGBA{217, 220, 214, 255},
	-6,
}

func GetBiome(level float64) Biome {
	for i := len(biomes) - 1; i >= 0; i-- {
		if level >= biomes[i].level {
			return biomes[i]
		}
	}

	// Should never happen
	return sea
}
