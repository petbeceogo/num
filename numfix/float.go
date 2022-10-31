package numfix

import (
	"fmt"
	"math"
)

func FloatToString(num float64) string {
	return fmt.Sprintf("%f", num)
}

func FloatToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

func FloatToFixedString(num float64, precision int) string {
	return FloatToString(FloatToFixed(num, precision))
}
