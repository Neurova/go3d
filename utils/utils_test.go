package utils

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestNumOccurrences(t *testing.T) {
	expectedOccurrences := getExpectedOccurrences()
	uniqueStringValues := getUniqueStringValues()
	uniqueFloatValues := getUniqueFloatvalues()
	uniqueIntValues := getUniqueIntValues()

	stringData := getStringData()
	stringOccurrences := []int{}
	floatData := getFloatData()
	floatOccurrences := []int{}
	intData := getIntData()
	intOccurrences := []int{}

	for _, value := range uniqueStringValues {
		n := NumOccurences(value, stringData)
		stringOccurrences = append(stringOccurrences, n)
	}

	for _, value := range uniqueFloatValues {
		n := NumOccurences(value, floatData)
		floatOccurrences = append(floatOccurrences, n)
	}

	for _, value := range uniqueIntValues {
		n := NumOccurences(value, intData)
		intOccurrences = append(intOccurrences, n)
	}

	if !reflect.DeepEqual(stringOccurrences, expectedOccurrences) {
		t.Errorf("TestNumOccurrences failed string test. Expected %v but recieved %v", expectedOccurrences, stringOccurrences)
	}

	if !reflect.DeepEqual(floatOccurrences, expectedOccurrences) {
		t.Errorf("TestNumOccurrences failed float64 test. Expected %v but recieved %v", expectedOccurrences, floatOccurrences)
	}

	if !reflect.DeepEqual(intOccurrences, expectedOccurrences) {
		t.Errorf("TestNumOccurrences failed int test. Expected %v but recieved %v", expectedOccurrences, intOccurrences)
	}

}

func TestIsIn(t *testing.T) {
	const intCheck = 1
	const floatCheck = 1.0
	const stringCheck = "2"

	stringData := getStringData()
	floatData := getFloatData()
	intData := getIntData()

	if !IsIn(intCheck, intData) {
		t.Errorf("TestIsIn failed int test. %v not found in %v", intCheck, intData)
	}

	if !IsIn(floatCheck, floatData) {
		t.Errorf("TestIsIn failed float test. %v not found in %v", floatCheck, floatData)
	}

	if !IsIn(stringCheck, stringData) {
		t.Errorf("TestIsIn failed string test. %v not found in %v", stringCheck, stringData)
	}
}

func TestUnique(t *testing.T) {
	uniqueStringValues := getUniqueStringValues()
	uniqueFloatValues := getUniqueFloatvalues()
	uniqueIntValues := getUniqueIntValues()

	stringData := getStringData()
	floatData := getFloatData()
	intData := getIntData()

	if !reflect.DeepEqual(Unique(stringData), uniqueStringValues) {
		t.Errorf("TestUnique failed string test. Expected %v but recieved %v", uniqueStringValues, Unique(stringData))
	}

	if !reflect.DeepEqual(Unique(floatData), uniqueFloatValues) {
		t.Errorf("TestUnique failed float test. Expected %v but recieved %v", uniqueFloatValues, Unique(floatData))
	}

	if !reflect.DeepEqual(Unique(intData), uniqueIntValues) {
		t.Errorf("TestUnique failed int test. Expected %v but recieved %v", uniqueIntValues, Unique(intData))
	}

}

func TestVectorToSlice(t *testing.T) {
	floatSlice := getFloatData()
	floatVector := mat.NewVecDense(len(floatSlice), floatSlice)

	if !reflect.DeepEqual(VectorToSlice(floatVector), floatSlice) {
		t.Errorf("TestVectorToslice failed. Expected %v but recieved %v", floatSlice, VectorToSlice(floatVector))
	}
}

func TestSumFloat64Slice(t *testing.T) {
	floatSlice := getFloatData()
	expectedSum := getExpectedFloat64DataSum()
	sum := SumSlice(floatSlice)

	if !reflect.DeepEqual(sum, expectedSum) {
		t.Errorf("TestSumFloat64Slice failed. Expected %v but recieved %v", expectedSum, sum)
	}
}

func TestSumIntSlice(t *testing.T) {
	intSlice := getIntData()
	expectedSum := getExpectedIntDataSum()
	sum := SumSlice(intSlice)

	if !reflect.DeepEqual(sum, expectedSum) {
		t.Errorf("TestSumIntSlice failed. Expected %v but recieved %v", expectedSum, sum)
	}

}

func getExpectedFloat64DataSum() float64 {
	return 13.0
}

func getExpectedIntDataSum() int {
	return 13
}

func getExpectedOccurrences() []int {
	return []int{2, 1, 3}
}

func getStringData() []string {
	return []string{"1", "1", "2", "3", "3", "3"}
}

func getFloatData() []float64 {
	return []float64{1.0, 1.0, 2.0, 3.0, 3.0, 3.0}
}

func getIntData() []int {
	return []int{1, 1, 2, 3, 3, 3}
}

func getUniqueStringValues() []string {
	return []string{"1", "2", "3"}
}

func getUniqueFloatvalues() []float64 {
	return []float64{1.0, 2.0, 3.0}
}

func getUniqueIntValues() []int {
	return []int{1, 2, 3}
}
