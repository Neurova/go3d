package utils

import (
	"gonum.org/v1/gonum/mat"
)

type Number interface {
	int | float64
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

func SumSlice[T Number](s []T) T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}
