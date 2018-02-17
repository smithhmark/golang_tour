package main

import "golang.org/x/tour/tree"

import (
        "fmt"
)

//func inorder(t *tree.Tree) int {
func inorder(t *tree.Tree, ch chan int) int {
        lcnt := 0
        rcnt := 0
        if t.Left != nil {
                lcnt = inorder(t.Left, ch)
        }
        ch <- t.Value
        //fmt.Println("v:",t.Value)
        if t.Right != nil {
                rcnt = inorder(t.Right, ch)
        }
        return lcnt + rcnt + 1
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
        cnt := inorder(t, ch)
        close(ch)
        fmt.Println("I visited:", cnt, "starting at:", t.Value)
}


// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
        leftch := make(chan int)
        rightch := make(chan int)
        go Walk(t1, leftch)
        go Walk(t2, rightch)
        for lval := range(leftch) {
                rval, ok := <- rightch
                if ok == false {
                        return false
                }
                if lval != rval {
                        return false
                } 
        }
        return true
}

func main() {
        t1 := tree.New(1)
        t2 := tree.New(1)
        s := Same(t1, t2)
        fmt.Println("the same?", s)
        t3 := tree.New(10)
        s = Same(t1, t3)
        fmt.Println("the same?", s)
}

