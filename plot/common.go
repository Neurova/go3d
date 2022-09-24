package plot

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/pkg/browser"
)

const port = ":8080"
const url = "http://localhost" + port
const baseHTML = `
<!DOCTYPE html>
<head>
  <script src="https://cdn.plot.ly/plotly-2.14.0.min.js"></script>
</head>
<body>
  <div id="plot"></div>
  <script>
	var height = window.screen.availHeight;
	var data = {{ .Data }}
	var layout = {{ .Layout }}
	if (layout === null) {
		layout = {"height": Math.floor(height / 1.25)};
	} else if (layout.height === undefined) {
		layout["height"] = Math.floor(heigh / 1.25)
	}
	var config = {responsive: true};
    Plotly.newPlot("plot", data, layout, config);
  </script>
</body>
`

type Figure struct {
	Data   []map[string]any // slice of maps containing data for each trace
	Layout *map[string]any  // optional map for defining the plot's layout
}

// NewFigure takes in a slice of maps. Each map is a trace for rendering
// and returns a pointer to the Figure object with a `nil` layout
// Figures with `nil` layout will have a `height` attribute assigned
// in JavaScript equal to available screen height / 1.25
func NewFigure(data []map[string]any) *Figure {
	return &Figure{data, nil}
}

// NewFigureWithLayout takes in a slice of maps for trace data, and a map
// defining the plots layout.
func NewFigureWithLayout(data []map[string]any, layout map[string]any) *Figure {
	return &Figure{data, &layout}
}

// Show function takes in a Figure struct and opens a browser tab at
// `localhost:8080` for rendering the figure. If the figure does not
// appear, please use the Developer Console in the browser to identify
// the issue. Most likely, an incorrect key: value pair was passed
// to plotly.js
func Show(f Figure) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = plotlyHandler(w, r, f)
	})

	err := show()
	return err
}

func show() error {
	fmt.Println("Opening at", url)
	err := browser.OpenURL(url)
	if err != nil {
		return err
	}

	err = http.ListenAndServe(port, nil)
	return err
}

func plotlyHandler(w http.ResponseWriter, r *http.Request, f Figure) error {
	t, err := template.New("plotly").Parse(baseHTML)
	if err != nil {
		return err
	}
	err = t.Execute(w, f)
	return err
}
