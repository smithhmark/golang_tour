package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
        a := -1
        b := 1
        fib := func () int {
                c := a + b
                a = b
                b = c
                return c
        }
        return fib
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
