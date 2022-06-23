package trie

type (
	Prefix [7]byte
	Flags  byte

	Comparable interface {
		Compare(interface{}) int
	}

	HandlerFunc func(value interface{})

	Node struct {
		Next   []Comparable
		Flags  Flags
		Prefix Prefix
	}
)

const (
	PrefixLength = len(Prefix{})

	HasChild  Flags = 1
	HasValues Flags = 2
)

func (node *Node) Cursor() Cursor {
	return &nodeCursor{current: node}
}

func (node *Node) FindAll(prefix []byte, hf HandlerFunc) {
	if existNode, _ := node.move(prefix); existNode != nil {
		existNode.Handle(hf)
	}
}

func (node *Node) move(prefix []byte) (cur, parent *Node) {
	cur = node
	for len(prefix) > 0 {
		var pl int
		parent = cur
		cur, pl = cur.findPrefix(prefix)
		if cur == nil {
			return nil, nil
		}

		prefix = prefix[pl:]
	}

	return
}

func (node *Node) Compare(interface{}) int {
	return 0
}

func (node *Node) Handle(hf HandlerFunc) {
	for _, n := range node.Next {
		if c, ok := n.(*Node); ok {
			c.Handle(hf)
		} else {
			hf(n)
		}
	}
}

func (node *Node) Delete(prefix []byte, values ...Comparable) {
	existNode, parent := node.move(prefix)

	for i, n := range parent.Next {
		if n == existNode {
			parent.Next = append(parent.Next[:i], parent.Next[i+1:]...)
			return
		}
	}
}

func (node *Node) Insert(prefix []byte, values ...Comparable) {
	existNode, pl := node.findPrefix(prefix)
	if existNode == nil {
		var newNode Node

		if s := copy(newNode.Prefix[:PrefixLength], prefix); s == PrefixLength && len(prefix) > PrefixLength {
			newNode.Insert(prefix[PrefixLength:], values...)
		} else {
			newNode.Next = append(newNode.Next, values...)
		}

		node.Next = append(node.Next, &newNode)
	} else {
		switch epl := prefixLength(existNode.Prefix); {
		case epl == pl:
			if pl == len(prefix) {
				existNode.Next = append(existNode.Next, values...)
			} else {
				existNode.Insert(prefix[pl:], values...)
			}
		case epl > pl:
			existNext := existNode.Next
			existNode.Next = nil

			existNode.Insert(existNode.Prefix[pl:], existNext...)

			for i := pl; i < len(existNode.Prefix); i++ {
				existNode.Prefix[i] = 0
			}

			existNode.Insert(prefix[pl:], values...)
		default:
			existNode.Insert(prefix[pl:], values...)
		}
	}
}

func (node *Node) Root() bool {
	return node.Prefix == Prefix{}
}

func (node *Node) Values(hf func(value Comparable)) {
	for _, n := range node.Next {
		if _, ok := n.(*Node); !ok {
			hf(n)
		}
	}
}

func (node *Node) Child(hf func(node *Node)) {
	for _, n := range node.Next {
		if n, ok := n.(*Node); ok {
			hf(n)
		}
	}
}

func (node *Node) findPrefix(prefix []byte) (*Node, int) {
	for _, n := range node.Next {
		if n, ok := n.(*Node); ok {
			if l := equalPrefixLength(n.Prefix[:], prefix); l > 0 {
				return n, l
			}
		}
	}

	return nil, 0
}

func prefixLength(v Prefix) (i int) {
	for i < len(v) && v[i] != 0 {
		i++
	}

	return
}

func equalPrefixLength(a, b []byte) (i int) {
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}

	return
}
