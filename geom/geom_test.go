package geom

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"

	"math"
)

func TestMagnitude(t *testing.T) {

	t.Parallel()

	// Precision test of 5
	const dig = 100000

	const cols int = 3
	const rows int = 2

	x := []float64{1, 2}
	y := []float64{3, 4}
	z := []float64{5, 6}

	ss0 := math.Round(math.Sqrt(x[0]*x[0]+y[0]*y[0]+z[0]*z[0])*dig) / dig
	ss1 := math.Round(math.Sqrt(x[1]*x[1]+y[1]*y[1]+z[1]*z[1])*dig) / dig

	a := mat.NewDense(rows, cols, nil)

	a.SetCol(0, x)
	a.SetCol(1, y)
	a.SetCol(2, z)

	mag := Magnitude(a)

	if len(mag) != 2 {
		t.Errorf("unexpected length of magnitude: got: %d  want: 2", len(mag))
	}

	if math.Round(mag[0]*dig)/dig != ss0 {
		t.Errorf("unexpected value for first magnituded: got: %e  want: %e", mag[0], ss0)
	}

	if math.Round(mag[1]*dig)/dig != ss1 {
		t.Errorf("unexpected value for second magnituded: got: %e  want: %e", mag[1], ss1)
	}

}

func TestNormalize(t *testing.T) {

	t.Parallel()

	const proofMagnitude float64 = 1
	const cols int = 3
	const rows int = 2

	x := []float64{1, 2}
	y := []float64{3, 4}
	z := []float64{5, 6}

	a := mat.NewDense(rows, cols, nil)
	a.SetCol(0, x)
	a.SetCol(1, y)
	a.SetCol(2, z)

	normA := Normalize(a)

	fmt.Println("A", a)
	fmt.Println("Magnitude A", Magnitude(a))
	fmt.Println("Normalized A", normA)

	var rowOneMag float64
	var rowTwoMag float64

	var rowOne mat.Vector = normA.RowView(0)
	var rowTwo mat.Vector = normA.RowView(1)

	rowOneMag = math.Sqrt(rowOne.AtVec(0)*rowOne.AtVec(0) + rowOne.AtVec(1)*rowOne.AtVec(1) + rowOne.AtVec(2)*rowOne.AtVec(2))
	rowTwoMag = math.Sqrt(rowTwo.AtVec(0)*rowTwo.AtVec(0) + rowTwo.AtVec(1)*rowTwo.AtVec(1) + rowTwo.AtVec(2)*rowTwo.AtVec(2))

	r, c := normA.Dims()

	if r != rows {
		t.Errorf("undexpected number of columns:  got %d  wanted: %d", c, cols)
	}

	if c != cols {
		t.Errorf("undexpected number of rows:  got %d  wanted: %d", r, rows)
	}

	if math.Round(rowOneMag) != proofMagnitude {
		t.Errorf("unexpected magnitude for row one:  got %f  wanted: 1", rowOneMag)
	}

	if math.Round(rowTwoMag) != proofMagnitude {
		t.Errorf("unexpected magnitude for row two:  got %f  wanted: 1", rowTwoMag)
	}

}

func TestCross(t *testing.T) {

	t.Parallel()

	const cols int = 3
	const rows int = 2
	const rowOneProof float64 = 0
	const rowTwoProof float64 = 2

	x1 := []float64{1, 2}
	y1 := []float64{3, 4}
	z1 := []float64{5, 6}

	x2 := []float64{1, 2}
	y2 := []float64{3, 3}
	z2 := []float64{5, 5}

	a := mat.NewDense(rows, cols, nil)
	b := mat.NewDense(rows, cols, nil)

	a.SetCol(0, x1)
	a.SetCol(1, y1)
	a.SetCol(2, z1)

	b.SetCol(0, x2)
	b.SetCol(1, y2)
	b.SetCol(2, z2)

	crossProduct := Cross(a, b)

	fmt.Println(crossProduct)

	var rowOne mat.Vector = crossProduct.RowView(0)
	var rowTwo mat.Vector = crossProduct.RowView(1)

	var rowOneSum float64 = rowOne.AtVec(0) + rowOne.AtVec(1) + rowOne.AtVec(2)
	var rowTwoSum float64 = rowTwo.AtVec(0) + rowTwo.AtVec(1) + rowTwo.AtVec(2)

	r, c := crossProduct.Dims()

	if r != rows {
		t.Errorf("unexpected number of rows:  got: %d  wanted: %d", r, rows)
	}

	if c != cols {
		t.Errorf("unexpected number of columns:  got: %d  wanted: %d", c, cols)
	}

	if rowOneSum != rowOneProof {
		t.Errorf("unexpected sum of row one values:  got: %f  wanted: %f", rowOneSum, rowOneProof)
	}

	if rowTwoSum != rowTwoProof {
		t.Errorf("unexpected sum of row two values:  got: %f  wanted: %f", rowTwoSum, rowTwoProof)
	}
}

func TestAngle(t *testing.T) {

	t.Parallel()

	// Precision test of 4
	const dig = 10000

	const cols int = 3
	const rows int = 2
	const angleOneProof float64 = 0
	const angleTwoProof float64 = 4.3066

	x1 := []float64{1, 2}
	y1 := []float64{3, 4}
	z1 := []float64{5, 6}

	x2 := []float64{1, 2}
	y2 := []float64{3, 3}
	z2 := []float64{5, 5}

	a := mat.NewDense(rows, cols, nil)
	b := mat.NewDense(rows, cols, nil)

	a.SetCol(0, x1)
	a.SetCol(1, y1)
	a.SetCol(2, z1)

	b.SetCol(0, x2)
	b.SetCol(1, y2)
	b.SetCol(2, z2)

	aNorm := Normalize(a)
	bNorm := Normalize(b)

	anglesPostNorm := Angle(a, b, "deg", false)
	anglesPreNorm := Angle(aNorm, bNorm, "deg", true)

	if anglesPostNorm[0] != anglesPreNorm[0] {
		t.Errorf("unexpected difference between pre and post normalized matrices in row one")
	}

	if math.Round(anglesPostNorm[1]*dig)/dig != math.Round(anglesPreNorm[1]*dig)/dig {
		t.Errorf("unexpected difference between pre and post normalized matrices in row two")
	}

	if anglesPostNorm[0] != angleOneProof {
		t.Errorf("unexoected angle for row one:  got %v  wanted: %v", anglesPostNorm[0], angleOneProof)
	}

	if math.Round(anglesPostNorm[1]*dig)/dig != angleTwoProof {
		t.Errorf("unexoected angle for row one:  got %v  wanted: %v", anglesPostNorm[1], angleTwoProof)
	}
}
