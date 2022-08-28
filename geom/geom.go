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

		v := m.RowView(r)

		dist[0] = mat.Norm(v, 2) // 2 = euclidean norm

	} else {

		for i := 0; i <= r-1; i++ {

			v := m.RowView(i)

			dist[i] = mat.Norm(v, 2) // 2 = euclidean norm
		}
	}
	return dist
}
