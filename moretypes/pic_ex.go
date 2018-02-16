package main

import "golang.org/x/tour/pic"

import (
        "math"
)

func Pic(dx, dy int) [][]uint8 {
        data := make([][]uint8, dy)
        for ri := range(data) {
                data[ri] = make([]uint8, dx)
        }
        for y := 0 ; y < dy ; y++ {
                for x := 0 ; x < dx ; x++ {
                        scaledx := float64(x)/float64(dx)
                        scaledy := float64(y)/float64(dy)
                        val := math.Sin(scaledx*scaledy)
                        data[y][x] = uint8(256 * val)
                        if x == y {
                                data[y][x] = 250
                        }
                }
        }
        return data
}

func main() {
    pic.Show(Pic)
}

