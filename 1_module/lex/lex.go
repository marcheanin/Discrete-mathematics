package main

import (
	"bufio"
	"fmt"
	"os"
)

type AssocArray interface {
	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}

type Map map[string]int

func (m Map) Assign(s string, x int) {
	m[s] = x
}

func (m Map) Lookup(s string) (x int, exists bool) {
	var elem, ok = m[s]
	return elem, ok
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
	m := make(Map)
	s, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	ans := lex(s, m)
	fmt.Println(ans)
}
