package trie

type Cursor []Pointer

func (c Cursor) Empty() bool {
	return len(c) == 0
}

func (c Cursor) Clone() Cursor {
	return append([]Pointer{}, c...)
}

func (c Cursor) Move(prefix []byte) Cursor {
	var i int
	for _, p := range c {
		if p, ok := p.Move(prefix); ok {
			c[i] = p
			i++
		}
	}

	return c[:i]
}

func (c Cursor) Handle(deep bool, hf HandlerFunc) {
	for _, p := range c {
		p.Handle(deep, hf)
	}
}

func (c Cursor) FindNextByte(nextByte byte) Cursor {
	var cursor Cursor
	for _, p := range c {
		cursor = append(cursor, p.MoveToByte(nextByte)...)
	}

	return cursor
}
