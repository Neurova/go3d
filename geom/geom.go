// package geom designed to enhance the interface of 3D point data
// license that can be found in the LICENSE file.
package geom

import (
	"gonum.org/v1/gonum/mat"
)

// magnitude function takes the rowwise euclidean distance of 2D or 3D points
// or vectors
func magnitude(m *mat.Dense) []float64 {

	if m.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	r, c := m.Dims()

	// return variable with the same length as m has rows
	dist := make([]float64, r)

	// Only accepts max of 3D (XYZ) data
	if c > 3 {
		panic(mat.ErrColLength)
	}

	if r == 1 {

		var v mat.Vector = m.RowView(r)

		dist[0] = mat.Norm(v, 2) // 2 = euclidean norm

	} else {

		for i := 0; i <= r-1; i++ {

			var v mat.Vector = m.RowView(i)

			dist[i] = mat.Norm(v, 2) // 2 = euclidean norm
		}
	}
	return dist
}

// normalize function applies rowwise normalization of 2D or 3D data
// to magnatiude 1
func normalize(m *mat.Dense) *mat.Dense {

	if m.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	r, c := m.Dims()

	var index = make([]float64, c)

	// return variable with the same shape as m
	norm := mat.NewDense(r, c, nil)

	var mag []float64 = magnitude(m)

	if r == 1 {

		var row []float64 = m.RawMatrix().Data

		var normRow = make([]float64, c)

		for idx, value := range row {

			normRow[idx] = value / mag[0] // if only 1 row, then mag only has a zero index

		}

		norm.SetRow(0, normRow)

	} else {

		for i := 0; i <= r-1; i++ {

			var normRow = make([]float64, c)

			var row []float64 = mat.Row(index, i, m)

			for idx, value := range row {

				normRow[idx] = value / mag[i]

			}

			norm.SetRow(i, normRow)
		}
	}

	return norm
}
