package plot

import (
	"net/http"
)

// take in x, y, and z coordinates as slices of float64s
// return a map ready to be JSONified for 3d scatter plotting
func createPDFMap(x, y []float64) map[string]any {
	m := make(map[string]any)
	m["type"] = "bar"
	m["mode"] = "lines"
	m["histnorm"] = "probability"
	m["x"] = x
	m["y"] = y
	return m

}

func PDF(x, y []float64) {
	m := createPDFMap(x, y)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		plotlyHandler(w, r, m)
	})

	plot()
}
