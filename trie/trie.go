package trie

type Trie struct {
	root *node
}

type node struct {
	isEnd    bool
	children map[rune]*node
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func newNode() *node {
	return &node{
		isEnd:    false,
		children: map[rune]*node{},
	}
}

func (t *Trie) Insert(textInRune []rune) {
	cur := t.root
	for _, r := range textInRune {
		if next, ok := cur.children[r]; ok {
			cur = next
		} else {
			next := newNode()
			cur.children[r] = next
			cur = next
		}
	}
	cur.isEnd = true
}

func (t *Trie) LongestMatchedPrefix(textInRune []rune) (prefix []rune) {
	cur := t.root
	for i, r := range textInRune {
		if next, ok := cur.children[r]; ok {
			cur = next
			if cur.isEnd {
				prefix = textInRune[:i+1]
			}
		} else {
			return
		}
	}
	return
}
