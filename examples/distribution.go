package examples

import (
	"fmt"
	"math/rand"

	"github.com/neurova/go3d/stats"
	"github.com/neurova/go3d/utils"
)

func PDFExample() {
	randomData := []float64{}
	var i int
	for i < 50 {
		randomFloat := rand.Float64()
		randomData = append(randomData, randomFloat)
		i++
	}

	pdf, uniqueValues := stats.PDF(randomData)
	fmt.Println("PDF values (y): ", pdf)
	fmt.Println("Unique values (x): ", uniqueValues)
	fmt.Println("Sum of values: ", utils.SumFloat64Slice(pdf))

}
