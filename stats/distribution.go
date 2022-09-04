package stats

import (
	"sort"

	"github.com/neurova/go3d/utils"
	"gonum.org/v1/gonum/mat"
)

func PDF(v mat.Vector) []float64 {
	n := v.Len()
	sortedValues := utils.VectorToSlice(v)
	sort.Slice(sortedValues, func(i, j int) bool {
		return sortedValues[i] < sortedValues[j]
	})

	uniqueValues := utils.Unique(sortedValues)
	pdf := []float64{}
	for _, uniqueValue := range uniqueValues {
		num := utils.NumOccurences(uniqueValue, sortedValues)
		pdf = append(pdf, float64(num)/float64(n))
	}
	return pdf
}
