package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// type Tree struct {
//    Left  *Tree
//    Value int
//    Right *Tree
//}

func WalkR(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkR(t.Left, ch)
	}
	if t.Right != nil {
		WalkR(t.Right, ch)
	}
	ch <- t.Value
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkR(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, c1ok := <-c1
		v2, c2ok := <-c2
		if v1 != v2 {
			return false
		}
		if !c1ok || !c2ok {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	s := Same(tree.New(1), tree.New(2))
	fmt.Printf("Must be false: %v\n", s)
	t := tree.New(1)
	s = Same(t, t)
	fmt.Printf("Must be true: %v\n", s)
}
