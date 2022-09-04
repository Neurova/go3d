package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/neurova/go3d/utils"
	"gonum.org/v1/gonum/mat"
)

func TestPDF(t *testing.T) {
	expectedPDF := []float64{0.25, 0.25, 0.25, 0.25}
	data := []float64{1, 1, 2, 2, 3, 3, 4, 4}
	uniqueValues := utils.Unique(data)
	expectedLen := len(utils.Unique(data))
	dataVector := mat.NewVecDense(len(data), data)
	fmt.Println(uniqueValues)

	pdf := PDF(dataVector)
	pdfSum := utils.SumFloat64Slice(pdf)

	if !reflect.DeepEqual(pdf, expectedPDF) {
		t.Errorf("TestPDF failed. Expected %v but recieved %v", expectedPDF, pdf)
	}

	if pdfSum != 1.0 {
		t.Errorf("TestPDF failed. Sum (%v) != 1.0", pdfSum)
	}

	if len(pdf) != expectedLen {
		t.Errorf("TestPDF failed. PDF has %v values but there are %v unique values", len(pdf), expectedLen)
	}
}
