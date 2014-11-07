package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    walk(t, ch)
    close(ch)
}

// depth-first recursive Tree walk, values sent on ch
func walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }
    walk(t.Left, ch)
    ch <- t.Value
    walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1, c2 := make(chan int), make(chan int)
    go Walk(t1, c1)
    go Walk(t2, c2)
    for {
        v1, ok1 := <- c1
        v2, ok2 := <- c2
        if ok1 != ok2 || v1 != v2 {
            return false
        }
        if !ok1 {
            break
        }
    }
    return true
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
