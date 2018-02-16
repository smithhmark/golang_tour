package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
    z := 1.0
	ep := 0.0001
	cnt := 0
	step := (z*z - x)/(2*z)
	for ; math.Abs(step) > ep; z -= step {
	   step = (z*z - x)/(2*z)
	   cnt += 1
	}
	fmt.Println(cnt)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

