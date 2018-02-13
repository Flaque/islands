package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenter(t *testing.T) {
	x, y := center(16, 16)

	assert.Equal(t, x, y)
	assert.Equal(t, x, 8.0)
}

func TestRadius(t *testing.T) {
	rad := radius(0, 8, 16, 16)
	assert.Equal(t, float64(8), float64(rad))
}

func TestSlantHeight(t *testing.T) {
	rad := radius(0, 8, 16, 16)
	assert.Equal(t, 12.81, toFixed(slantHeight(rad, 10), 2))
}
