package main

import (
	"log"
	"math"

	"github.com/neurova/go3d/plot"
)

func main() {
	// initialize slices of float64 for x, y, and z coordinates
	x := []float64{}
	y := []float64{}
	z := []float64{}

	// generate a circle
	for i := 0; i < 10000; i++ {
		x = append(x, math.Sin(-1*float64(i)*10.0))
		y = append(y, math.Sin(float64(i)*10.0))
		z = append(z, math.Cos(float64(i)*10.0))
	}

	trace := make(map[string]any)
	trace["x"] = x
	trace["y"] = y
	trace["z"] = z
	trace["mode"] = "markers"
	trace["type"] = "scatter3d"

	fig := plot.NewScatter3DFigure([]map[string]any{trace}, nil)

	err := plot.Show(fig)

	if err != nil {
		log.Fatal(err)
	}
}
