package utils

import (
	"fmt"
	"strconv"
)

func Atoi(s string) int {
	convertedInt, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("error converting string to int: %v", err)
		return 0
	}
	return convertedInt
}

func Atof(s string) float64 {
	convertedFloat, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("error converting string to float64, %v", err)
		return 0.0
	}
	return convertedFloat
}