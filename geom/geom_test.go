package geom

import (
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
