package examples

import (
	"math"

	"github.com/neurova/go3d/plot"
	"github.com/neurova/go3d/stats"
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

func PDFPlotExample() {
	data := []float64{1.0, 1.0, 2.0, 2.0, 2.0, 2.0, 3.0, 3.0, 3.0, 3.0, 3.0, 3.0, 4.0, 4.0, 4.0, 5.0, 5.0, 5.0, 6.0}
	values, probability := stats.PDF(data)
	plot.PDF(values, probability)
}
