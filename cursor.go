package trie

type Cursor []Pointer

func (c Cursor) Empty() bool {
	return len(c) == 0
}

func (c Cursor) Clone() Cursor {
	return append([]Pointer{}, c...)
}

func (c Cursor) Move(prefix []byte) Cursor {
	var cursor Cursor
	for _, p := range c {
		if p, ok := p.Move(prefix); ok {
			cursor = append(cursor, p)
		}
	}

	return cursor
}

func (c Cursor) Handle(hf HandlerFunc) {
	for _, p := range c {
		p.Handle(hf)
	}
}

func (c Cursor) FindNextByte(nextByte byte) Cursor {
	var cursor Cursor
	for _, p := range c {
		cursor = append(cursor, p.MoveToByte(nextByte)...)
	}

	return cursor
}
