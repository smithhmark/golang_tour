package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
        for ii := range(b) {
                b[ii] = 'A'
        }
        return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

