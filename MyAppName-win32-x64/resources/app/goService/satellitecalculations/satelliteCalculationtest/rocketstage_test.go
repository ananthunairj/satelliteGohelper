package satellitecalculationtest

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"github.com/Anandhu3301/satelliteGohelper/satellitecalculations"
)


var helpers = struct {
TestRoundFloatNumbers func(float64, int) float64
}{
TestRoundFloatNumbers: func(num float64, precision int) float64 {
// Implement a simple rounding function for testing
format := "%." + strconv.Itoa(precision) + "f"
str := fmt.Sprintf(format, num)
result, _ := strconv.ParseFloat(str, 64)
return result
}, 
}

func TestRemainingMassCalculator(t *testing.T) {
tests := []struct {
massSlice    []float64
index        int
fuelburnrate float64
expected     []float64
}{
{[]float64{100.0, 200.0, 300.0}, 1, 50.0, []float64{100.0, 150.0, 300.0}},
{[]float64{100.0, 200.0, 300.0}, 0, 25.0, []float64{75.0, 200.0, 300.0}},
{[]float64{100.0, 200.0, 300.0}, 2, 100.0, []float64{100.0, 200.0, 200.0}},
}

for _, tt := range tests {
err := satellitecalculations.RemainingMassCalculator(tt.massSlice, tt.index, tt.fuelburnrate)
if err != nil {
t.Errorf("Unexpected error: %v", err)
}
if !reflect.DeepEqual(tt.massSlice, tt.expected) {
t.Errorf("RemainingMassCalculator(%v, %d, %f) = %v; want %v", tt.massSlice, tt.index, tt.fuelburnrate, tt.massSlice, tt.expected)
}
}
}
