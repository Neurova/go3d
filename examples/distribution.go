package examples

import (
	"fmt"
	"math/rand"

	"github.com/neurova/go3d/stats"
	"github.com/neurova/go3d/utils"
	"gonum.org/v1/gonum/mat"
)

func PDFExample() {
	randomData := []float64{}
	var i int
	for i < 50 {
		randomFloat := rand.Float64()
		randomData = append(randomData, randomFloat)
		i++
	}

	randomVector := mat.NewVecDense(len(randomData), randomData)
	pdf := stats.PDF(randomVector)
	fmt.Println("PDF values: ", pdf)
	fmt.Println("Sum of values: ", utils.SumFloat64Slice(pdf))

}
