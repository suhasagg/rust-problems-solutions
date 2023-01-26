func wordBreak(s string, wordDict []string) bool {
     trie := NewTrie()
	 for _, word := range wordDict {
		trie.Insert(word)
	}
    return wordbreak(s, trie)
}

type TrieNode struct {
	children map[rune]*TrieNode
	endOfWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode), endOfWord: false}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		if _, ok := current.children[char]; !ok {
			current.children[char] = NewTrieNode()
		}
		current = current.children[char]
	}
	current.endOfWord = true
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		if _, ok := current.children[char]; !ok {
			return false
		}
		current = current.children[char]
	}
	return current.endOfWord
}

func wordbreak(s string, trie *Trie) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && trie.Search(s[j:i]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
