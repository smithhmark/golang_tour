package main

import (
	"golang.org/x/tour/wc"
        "strings"
)

func WordCount(s string) map[string]int {
        flds := strings.Fields(s)
        m := make(map[string]int)
        for _, fld := range(flds) {
                cnt := m[fld]
                m[fld] = cnt + 1
        }
	return m
}

func main() {
	wc.Test(WordCount)
}
