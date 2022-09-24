package plot

func NewScatter3DFigure(data []map[string]any, layout *map[string]any) Figure {
	for _, trace := range data {
		trace["mode"] = "markers"
		trace["type"] = "scatter3d"
	}
	return Figure{data, layout}
}
