package trie

type Pointer struct {
	current, parent *Node
	offset          int
}

func (p Pointer) Node() *Node {
	return p.current
}

func (p Pointer) Handle(deep bool, hf HandlerFunc) {
	if deep || p.Done() {
		p.current.Handle(deep, hf)
	}
}

func (p Pointer) Done() bool {
	return p.offset == prefixLength(p.current.Prefix)
}

func (p Pointer) Move(prefix []byte) (Pointer, bool) {
	for len(prefix) > 0 {
		if p.current == nil {
			break
		}

		if l := equalPrefixLength(p.current.Prefix[p.offset:], prefix); l > 0 {
			p.offset += l
			prefix = prefix[l:]
		} else {
			p.parent = p.current
			p.current, p.offset = p.current.findPrefix(prefix)
			prefix = prefix[p.offset:]
		}
	}

	return p, p.current != nil
}

func (p Pointer) MoveToByte(toByte byte) []Pointer {
	if !p.current.Root() {
		for i, b := range p.current.Prefix[p.offset:] {
			if b == toByte {
				p.offset += i + 1
				return []Pointer{p}
			}
		}
	}

	var pointers []Pointer
	p.current.Child(func(node *Node) {
		pointers = append(pointers, Pointer{current: node}.MoveToByte(toByte)...)
	})

	return pointers
}
