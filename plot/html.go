package plot

const BaseHTML = `
<!DOCTYPE html>
<head>
  <script src="https://cdn.plot.ly/plotly-2.14.0.min.js"></script>
</head>
<body>
  <div id="plot"></div>
  <script>
    var data = JSON.parse("{{.}}");
    console.log(data);
    var layout = {
      margin: {
        l: 0,
        r: 0,
        b: 0,
        t: 0,
      },
      height: 850,
    };
    Plotly.newPlot("plot", [data], layout);
  </script>
</body>
`
