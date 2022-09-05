package plot

import (
	"net/http"
)

// take in x, y, and z coordinates as slices of float64s
// return a map ready to be JSONified for 3d scatter plotting
func createScatter3dMap(x, y, z []float64) map[string]any {
	scatter3dMap := make(map[string]any)
	scatter3dMap["mode"] = "markers"
	scatter3dMap["type"] = "scatter3d"
	scatter3dMap["marker"] = map[string]any{"size": 2, "symbol": "circle", "opacity": 0.8}
	scatter3dMap["x"] = x
	scatter3dMap["y"] = y
	scatter3dMap["z"] = z
	return scatter3dMap

}

func Scatter3d(x, y, z []float64) {
	m := createScatter3dMap(x, y, z)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		plotlyHandler(w, r, m)
	})

	plot()
}
