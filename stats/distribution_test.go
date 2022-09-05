package stats

import (
	"reflect"
	"testing"

	"github.com/neurova/go3d/utils"
)

func TestPDF(t *testing.T) {
	expectedProbability := []float64{0.25, 0.25, 0.25, 0.25}
	data := []float64{1, 1, 2, 2, 3, 3, 4, 4}
	expectedValues := utils.Unique(data)
	expectedLen := len(utils.Unique(data))

	values, probability := PDF(data)
	probabilitySum := utils.SumFloat64Slice(probability)

	if !reflect.DeepEqual(probability, expectedProbability) {
		t.Errorf("TestPDF failed. Expected %v but recieved %v", expectedProbability, probability)
	}

	if probabilitySum != 1.0 {
		t.Errorf("TestPDF failed. Sum (%v) != 1.0", probabilitySum)
	}

	if len(probability) != expectedLen {
		t.Errorf("TestPDF failed. PDF has %v values but there are %v unique values", len(probability), expectedLen)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("TestPDF failed, PDF returned %v unique values but expected %v", values, expectedValues)
	}
}
