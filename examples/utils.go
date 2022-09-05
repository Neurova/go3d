package examples

import (
	"fmt"

	"github.com/neurova/go3d/utils"
)

func UniqueExample() {
	data := []float64{1, 1, 2, 2, 3, 3}
	fmt.Println("Original: ", data)
	fmt.Println("Unique: ", utils.Unique(data))
}

func IsInExample() {
	check := 1.0
	data := []float64{1, 2, 3}
	isIn := utils.IsIn(check, data)
	fmt.Println("Original: ", data)
	fmt.Printf("%v is in %v? %v", check, data, isIn)
}
