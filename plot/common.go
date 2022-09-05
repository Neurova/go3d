package plot

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/neurova/go3d/utils"
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
    var data = JSON.parse("{{.}}");
    console.log(data);
    var layout = {
	  	height: Math.floor(height / 1.25),
    };
	var config = {responsive: true};
    Plotly.newPlot("plot", [data], layout, config);
  </script>
</body>
`

func plotlyHandler(w http.ResponseWriter, r *http.Request, data map[string]any) {
	t, err := template.New("plotly").Parse(baseHTML)
	if err != nil {
		log.Fatal(err)
	}
	JSONString := utils.Jsonify(data)
	err = t.Execute(w, JSONString)
	if err != nil {
		log.Fatal(err)
	}
}

func plot() {
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

type mapCreator func(...any) map[string]any
