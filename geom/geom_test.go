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

	x := []float64{1, 2}
	y := []float64{3, 4}
	z := []float64{5, 6}

	ss0 := math.Round(math.Sqrt(x[0]*x[0]+y[0]*y[0]+z[0]*z[0])*dig) / dig
	ss1 := math.Round(math.Sqrt(x[1]*x[1]+y[1]*y[1]+z[1]*z[1])*dig) / dig

	a := mat.NewDense(2, 3, nil)
	a.SetCol(0, x)
	a.SetCol(1, y)
	a.SetCol(2, z)

	mag := magnitude(a)

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

	normA := normalize(a)

	fmt.Println("A", a)
	fmt.Println("Magnitude A", magnitude(a))
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

// 5.91 row 1 mangnitude
// .169 + .507 +
