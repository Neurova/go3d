// package geom designed to enhance the interface of 3D point data
// license that can be found in the LICENSE file.
package geom

import (
	"fmt"
	"math"

	"github.com/neurova/go3d/utils"
	"gonum.org/v1/gonum/mat"
)

var approvedUnits = []string{"deg", "degrees", "rad", "radians"}

// convert radians to degrees
func toDegrees(a []float64) []float64 {

	var degrees = make([]float64, len(a))

	for i, v := range a {
		degrees[i] = v * 180 / math.Pi
	}

	return degrees
}

// convert degrees to radians
func toRadians(a []float64) []float64 {

	var radians = make([]float64, len(a))

	for i, v := range a {
		radians[i] = v * math.Pi / 180
	}

	return radians
}

// magnitude function takes the rowwise euclidean distance of 2D or 3D points
// or vectors
func Magnitude(m *mat.Dense) []float64 {

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

		var v mat.Vector = m.RowView(0)

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
func Normalize(m *mat.Dense) *mat.Dense {

	if m.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	r, c := m.Dims()

	var index = make([]float64, c)

	// return variable with the same shape as m
	norm := mat.NewDense(r, c, nil)

	var mag []float64 = Magnitude(m)

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

// Finds the cross product between two single or stacked vectors
func Cross(a *mat.Dense, b *mat.Dense) *mat.Dense {

	if a.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	if b.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	a_r, a_c := a.Dims()
	b_r, b_c := b.Dims()

	if a_r != b_r {
		panic(mat.ErrShape)
	}

	if a_c != b_c {
		panic(mat.ErrShape)
	}

	m := mat.NewDense(a_r, a_c, nil) // can chose either a or b since they are required to be the same shape

	for i := 0; i <= a_r-1; i++ {

		var rowA mat.Vector = a.RowView(i)
		var rowB mat.Vector = b.RowView(i)

		var crossRow = make([]float64, a_c)

		crossRow[0] = rowA.AtVec(1)*rowB.AtVec(2) - rowA.AtVec(2)*rowB.AtVec(1)
		crossRow[1] = rowA.AtVec(2)*rowB.AtVec(0) - rowA.AtVec(0)*rowB.AtVec(2)
		crossRow[2] = rowA.AtVec(0)*rowB.AtVec(1) - rowA.AtVec(1)*rowB.AtVec(0)

		m.SetRow(i, crossRow)
	}

	return m
}

// Find the angle between 2 vectors
func Angle(m1 *mat.Dense, m2 *mat.Dense, units string, is_normalized bool) []float64 {

	var degUnits = []string{"deg", "degrees"}
	var mOneNorm *mat.Dense
	var mTwoNorm *mat.Dense

	if m1.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	if m2.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	m1_r, m1_c := m1.Dims()
	m2_r, m2_c := m2.Dims()

	if m1_r != m2_r {
		panic(mat.ErrShape)
	}

	if m1_c != m2_c {
		panic(mat.ErrShape)
	}

	if !utils.IsIn(units, approvedUnits) {
		fmt.Printf("unexpected value for units: got: %v  wanted either: %v", units, approvedUnits)
	}

	var dot = make([]float64, m1_r)
	var angles = make([]float64, m1_r)

	if is_normalized {

		mOneNorm = m1
		mTwoNorm = m2

	} else {
		mOneNorm = Normalize(m1)
		mTwoNorm = Normalize(m2)
	}

	if m1_r == 1 {

		rowM1 := mOneNorm.RowView(0)
		rowM2 := mTwoNorm.RowView(0)

		dot[0] = mat.Dot(rowM1, rowM2)

	} else {

		for i := 0; i < m1_r; i++ {

			rowM1 := mOneNorm.RowView(i)
			rowM2 := mTwoNorm.RowView(i)

			dot[i] = mat.Dot(rowM1, rowM2)
		}
	}

	for i, v := range dot {

		angles[i] = math.Acos(math.Min(math.Max(v, -1), 1))
	}

	if utils.IsIn(units, degUnits) {

		angles = toDegrees(angles)

	}

	return angles

}

// computes the scalar needed to project m1 onto m2
func ScalarProjection(m1 *mat.Dense, m2 *mat.Dense, normalized bool) []float64 {

	if m1.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	if m2.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	m1_r, m1_c := m1.Dims()
	m2_r, m2_c := m2.Dims()

	if m1_r != m2_r {
		panic(mat.ErrShape)
	}

	if m1_c != m2_c {
		panic(mat.ErrShape)
	}

	var project = make([]float64, m1_r)
	var mTwoNorm *mat.Dense

	if normalized {
		mTwoNorm = m2
	} else {
		mTwoNorm = Normalize(m2)
	}

	if m1_r == 1 {

		v1 := m1.RowView(0)
		v2 := mTwoNorm.RowView(0)

		project[0] = mat.Dot(v1, v2)

	} else {

		for i := 0; i < m1_r; i++ {

			v1 := m1.RowView(i)
			v2 := mTwoNorm.RowView(i)

			project[i] = mat.Dot(v1, v2)
		}
	}

	return project
}

// projects vector(s) m1 onto vector(s) m2
func ProjectVector(m1 *mat.Dense, m2 *mat.Dense, normalized bool) *mat.Dense {

	if m1.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	if m2.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	m1_r, m1_c := m1.Dims()
	m2_r, m2_c := m2.Dims()

	if m1_r != m2_r {
		panic(mat.ErrShape)
	}

	if m1_c != m2_c {
		panic(mat.ErrShape)
	}

	mProject := mat.NewDense(m1_r, m1_c, nil)

	var mTwoNorm *mat.Dense

	if normalized {
		mTwoNorm = m2
	} else {
		mTwoNorm = Normalize(m2)
	}

	scalar := ScalarProjection(m1, mTwoNorm, true)

	if m1_r == 1 {

		onto := mTwoNorm.RawRowView(0)

		var proj = make([]float64, m1_c)

		for i := 0; i < m1_c; i++ {
			proj[i] = onto[i] * scalar[0]
		}

		mProject.SetRow(0, proj)

	} else {

		for i := 0; i < m1_r; i++ {

			onto := mTwoNorm.RawRowView(i)

			var proj = make([]float64, m1_c)

			for ii := 0; ii < m1_c; ii++ {
				proj[ii] = onto[ii] * scalar[i]
			}

			mProject.SetRow(i, proj)
		}
	}

	return mProject

}

// implements the rodrigues rotation formula to rotate m around axis a
func Rotate(m *mat.Dense, a *mat.Dense, angle float64, units string, normalized bool) *mat.Dense {

	if m.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	if a.IsEmpty() {
		panic(mat.ErrZeroLength)
	}

	mr, mc := m.Dims()
	ar, ac := a.Dims()

	if mr != ar {
		panic(mat.ErrShape)
	}

	if mc != ac {
		panic(mat.ErrShape)
	}

	if !utils.IsIn(units, approvedUnits) {
		fmt.Printf("unexpected value for units: got: %v  wanted either: %v", units, approvedUnits)
	}

	if utils.IsIn(units, []string{"degrees", "deg"}) {
		angle = angle * math.Pi / 180
	}

	var dots = make([]float64, mc)

	var c = mat.NewDense(mr, mc, nil)
	var s = mat.NewDense(mr, mc, nil)
	var d = mat.NewDense(mr, mc, nil)
	var r mat.Dense

	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	if !normalized {
		a = Normalize(a)
	}

	for i := 0; i < mr; i++ {

		var vcos = make([]float64, mc)

		row := m.RawRowView(i)

		for ii, v := range row {
			vcos[ii] = cosine * v
		}

		c.SetRow(i, vcos)
	}

	cross := Cross(a, m)

	for i := 0; i < mr; i++ {

		var vsin = make([]float64, mc)

		row := cross.RawRowView(i)

		for ii, v := range row {
			vsin[ii] = sine * v
		}

		s.SetRow(i, vsin)
	}

	for i := 0; i < mr; i++ {
		dots[i] = mat.Dot(a.RowView(i), m.RowView(i))
	}

	for i := 0; i < mr; i++ {

		var vdots = make([]float64, mc)

		row := a.RawRowView(i)

		for ii, v := range row {
			vdots[ii] = (1 - cosine) * dots[i] * v
		}

		d.SetRow(i, vdots)
	}

	r.Add(c, s)
	r.Add(&r, d)

	return &r

}
