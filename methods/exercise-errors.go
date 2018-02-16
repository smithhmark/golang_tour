package main

import (
	"fmt"
        "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
        num := float64(e)
        return fmt.Sprintf("cannot Sqrt of negative number: %v", num)
}

func Sqrt(x float64) (float64, error) {
        if x < 0 {
                return 0, ErrNegativeSqrt(0)
        }
        z := 1.0
	ep := 0.0001
	cnt := 0
	step := (z*z - x)/(2*z)
	for ; math.Abs(step) > ep; z -= step {
                step = (z*z - x)/(2*z)
                cnt += 1
	}
	fmt.Println(cnt)
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

