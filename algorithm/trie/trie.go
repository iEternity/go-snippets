package trie

type TrieNode struct {
	children map[rune]*TrieNode
	hasData bool
}

type Trie struct {
	root *TrieNode
}

func NewNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func New() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root

	for i, ch := range word {
		if cur.children[ch] == nil {
			cur.children[ch] = NewNode()
		}
		cur = cur.children[ch]

		if i == len(word) - 1 {
			cur.hasData = true
		}
	}
}

func (t *Trie) Find(word string) bool {
	cur := t.root
	for _, ch := range word {
		if cur.children[ch] == nil {
			return false
		}
		cur = cur.children[ch]

		if !cur.hasData {
			return false
		}
	}

	return true
}