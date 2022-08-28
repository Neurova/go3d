package plot

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	// "github.com/neurova/go3d/templates"
	"github.com/pkg/browser"
)

const port = ":8080"
const url = "http://localhost" + port

func jsonifyMap(input map[string]any) string {
	jsonString, err := json.MarshalIndent(input, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonString)
}

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

// handle sending user input to plotly front end
func scatter3dHandler(w http.ResponseWriter, r *http.Request, x, y, z []float64) {
	t, err := template.New("plotly").Parse(BaseHTML)
	if err != nil {
		log.Fatal(err)
	}
	scatter3dMap := createScatter3dMap(x, y, z)
	scatter3dJSONString := jsonifyMap(scatter3dMap)
	t.Execute(w, scatter3dJSONString)
}

func Scatter3d(x, y, z []float64) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		scatter3dHandler(w, r, x, y, z)
	})

	fmt.Println("Opening at", url)
	err := browser.OpenURL(url)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
