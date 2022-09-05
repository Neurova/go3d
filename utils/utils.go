package utils

import (
	"encoding/json"
	"log"

	"gonum.org/v1/gonum/mat"
)

func Jsonify(input map[string]any) string {
	jsonByte, err := json.MarshalIndent(input, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonByte)
}

func NumOccurences[T comparable](check T, slice []T) int {
	var occurences int
	for _, value := range slice {
		if value == check {
			occurences += 1
		}
	}
	return occurences
}

func IsIn[T comparable](check T, slice []T) bool {
	for _, value := range slice {
		if value == check {
			return true
		}
	}
	return false
}

func Unique[T comparable](slice []T) []T {
	uniqueValues := []T{}
	for _, value := range slice {
		if !IsIn(value, uniqueValues) {
			uniqueValues = append(uniqueValues, value)
		}
	}
	return uniqueValues
}

func VectorToSlice(v mat.Vector) []float64 {
	n := v.Len()
	s := []float64{}
	for i := 0; i < n; i++ {
		s = append(s, v.AtVec(i))
	}
	return s
}

func SumIntSlice(slice []int) int {
	var sum int
	for _, value := range slice {
		sum += value
	}
	return sum
}

func SumFloat64Slice(slice []float64) float64 {
	var sum float64
	for _, value := range slice {
		sum += value
	}
	return sum
}
