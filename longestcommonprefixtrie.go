package main

import (
	"fmt"
)

type Trie struct {
	children map[rune]*Trie
	isLeaf   bool
}

func (t *Trie) Insert(str string) {
	curr := t
	for _, r := range str {
		if _, ok := curr.children[r]; !ok {
			curr.children[r] = &Trie{children: make(map[rune]*Trie)}
		}
		curr = curr.children[r]
	}
	curr.isLeaf = true
}

func (t *Trie) LCP() string {
	res := ""
	curr := t
	for len(curr.children) == 1 && !curr.isLeaf {
		for r := range curr.children {
			res += string(r)
			curr = curr.children[r]
		}
	}
	return res
}

func main() {
	trie := &Trie{children: make(map[rune]*Trie)}
	strs := []string{"hello", "heaven", "heavy", "hen"}
	for _, str := range strs {
		trie.Insert(str)
	}
	fmt.Println(trie.LCP())
}
// Beats 100 percent leet code solutions - Runtime of Oms.
