package main

import "golang.org/x/tour/pic"

import (
        //"fmt"
        "math"
        "image"
        "image/color"
)

type Image struct{
        Hi string
}
/*
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
*/
func (i Image) ColorModel() color.Model {
        return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
        return image.Rect(0,0,256,256)
}

func (i Image) At(x,y int) color.Color {
        bnds := i.Bounds()
        dx := bnds.Max.X
        dy := bnds.Max.Y
        sx := float64(x)/float64(dx/3)
        sy := float64(y)/float64(dy/3)
        hyp := math.Sqrt(sx*sx+sy*sy)
        sin := math.Sin(hyp)
        scaled := float64(math.MaxUint8) * sin
        //scaledx := float64(math.MaxUint8) * math.Sin(float64(x))
        //scaledy := float64(math.MaxUint8) * math.Sin(float64(y))
        //val := math.Sin(scaledx*scaledy)
        //data := uint8(float64(math.MaxUint8) * val )
        data := uint8(scaled)
        if x == y {
                data = math.MaxUint8
        }
        //fmt.Printf("x:%v y:%v v:%v d:%v\n", x, y, val, data)
        return color.RGBA{data,data,data,math.MaxUint8}
}

func main() {
	m := Image{}
        //fmt.Println(m.At(0,0))
        //fmt.Println(m.At(2,3))
        //fmt.Println(m.At(4,4))
        //fmt.Println(m.At(4,5))
	pic.ShowImage(m)
}

