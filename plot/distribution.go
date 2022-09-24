package plot

func NewPDFFigure(data []map[string]any, layout *map[string]any) Figure {
	for _, trace := range data {
		trace["type"] = "bar"
		trace["mode"] = "lines"
		trace["histnorm"] = "probability"
	}
	return Figure{data, layout}
}
