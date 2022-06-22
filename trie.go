package trie

type Trie struct {
	root *Node
}

func (t *Trie) Insert(kw string, values ...Comparable) {
	if t.root == nil {
		t.root = &Node{}
	}

	t.root.Insert([]byte(kw), values...)
}

func (t *Trie) FindAll(prefix string, hf HandlerFunc) {
	t.root.FindAll([]byte(prefix), hf)
}
