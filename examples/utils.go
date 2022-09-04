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
