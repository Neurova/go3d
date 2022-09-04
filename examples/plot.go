package examples

import (
	"math"

	"github.com/neurova/go3d/plot"
)

func Scatter3dExample() {
	// initialize slices of float64 for x, y, and z coordinates
	x := []float64{}
	y := []float64{}
	z := []float64{}

	// generate a circle
	for i := 0; i < 100000; i++ {
		x = append(x, math.Sin(-1*float64(i)*10.0))
		y = append(y, math.Sin(float64(i)*10.0))
		z = append(z, math.Cos(float64(i)*10.0))
	}
	// plot
	plot.Scatter3d(x, y, z)
}
