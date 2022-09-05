package stats

import (
	"sort"

	"github.com/neurova/go3d/utils"
)

// returns probability, unique values
func PDF(data []float64) ([]float64, []float64) {
	n := len(data)
	// data := utils.VectorToSlice(v)
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	uniqueValues := utils.Unique(data)
	probability := []float64{}
	for _, uniqueValue := range uniqueValues {
		num := utils.NumOccurences(uniqueValue, data)
		probability = append(probability, float64(num)/float64(n))
	}
	return uniqueValues, probability
}
