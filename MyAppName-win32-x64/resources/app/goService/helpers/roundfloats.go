package helpers

import "math"

func RoundFloatNumbers(number float64, precision uint) float64 {
	ratio :=  number * math.Pow10(int(precision));
	return math.Round(number * ratio) / ratio;
}