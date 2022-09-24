package main

import (
	"log"

	"github.com/neurova/go3d/plot"
	"github.com/neurova/go3d/stats"
)

func main() {
	origData := []float64{1.0, 1.0, 2.0, 2.0, 2.0, 2.0, 3.0, 3.0, 3.0, 3.0, 3.0, 3.0, 4.0, 4.0, 4.0, 5.0, 5.0, 5.0, 6.0}
	values, probability := stats.PDF(origData)
	trace := make(map[string]any)
	trace["x"] = values
	trace["y"] = probability
	fig := plot.NewPDFFigure([]map[string]any{trace}, nil)
	err := plot.Show(fig)
	if err != nil {
		log.Fatal(err)
	}
}
