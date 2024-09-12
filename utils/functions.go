package utils

import (
	"fmt"
	"strings"
)

func Includes(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func Float64SliceToString(slice []float64) string {
	// Create a slice of strings with the same length as the input slice
	strs := make([]string, 2)

	strs[0] = fmt.Sprintf("%f", slice[1])
	strs[1] = fmt.Sprintf("%f", slice[0])

	// Join the strings with commas
	return strings.Join(strs, ",")
}
