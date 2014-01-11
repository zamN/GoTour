package main

import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkHelper(t, ch)
	close(ch)
}

func WalkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkHelper(t.Left, ch)
	ch <- t.Value
	WalkHelper(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	t1Node, alive1 := <-ch1
	t2Node, alive2 := <-ch2

	for alive1 && alive2 {
		if t1Node != t2Node {
			return false
		}
		t1Node, alive1 = <-ch1
		t2Node, alive2 = <-ch2
	}

	return !(alive1 && alive2)
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
