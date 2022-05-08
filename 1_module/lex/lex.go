package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type AVLNode struct {
	key    string
	Value  int
	height int
	left   *AVLNode
	right  *AVLNode
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLNode) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *AVLNode) rebalanceTree() *AVLNode {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

func (n *AVLNode) add(key string, value int) *AVLNode {
	if n == nil {
		return &AVLNode{key, value, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key, value)
	} else if key > n.key {
		n.right = n.right.add(key, value)
	} else {
		n.Value = value
	}
	return n.rebalanceTree()
}

func (n *AVLNode) search(key string) (int, bool) {
	if n == nil {
		return 0, false
	}
	if key < n.key {
		return n.left.search(key)
	} else if key > n.key {
		return n.right.search(key)
	} else {
		return n.Value, true
	}
}

type AssocArray interface {
	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}

type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Assign(s string, x int) {
	t.root = t.root.add(s, x)
}

func (t *AVLTree) Lookup(s string) (x int, exists bool) {
	return t.root.search(s)

}

func lex(sentence string, array AssocArray) []int {
	ans := make([]int, 0, 0)
	tokens := make([]string, 0, 0)
	for i := 0; i < len(sentence); i++ {
		if sentence[i] != ' ' {
			pos2 := i
			for pos2 < len(sentence) && sentence[pos2] != ' ' {
				pos2++
			}
			tokens = append(tokens, sentence[i:pos2])
			i = pos2
		}
	}

	counter := 1
	for i := range tokens {
		elem, ok := array.Lookup(tokens[i])
		if ok {
			ans = append(ans, elem)
		} else {
			array.Assign(tokens[i], counter)
			ans = append(ans, counter)
			counter++
		}
	}
	return ans
}

func main() {
	var s string
	tree := &AVLTree{}
	t := AssocArray(tree)
	s, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	ans := lex(s, t)
	fmt.Println(ans)
}
