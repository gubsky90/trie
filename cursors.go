package trie

type nodeCursorsArray []nodeCursor

func (c nodeCursorsArray) Move(prefix []byte) Cursor {
	return c.move(prefix)
}

func (c nodeCursorsArray) move(prefix []byte) nodeCursorsArray {
	var res nodeCursorsArray
	for _, cur := range c {
		cur.Move(prefix)
		if !cur.Empty() {
			res = append(res, cur)
		}
	}

	return res
}

func (c nodeCursorsArray) Empty() bool {
	return len(c) == 0
}

func (c nodeCursorsArray) Handle(hf HandlerFunc) {
	for _, cur := range c {
		cur.Handle(hf)
	}
}

func (c nodeCursorsArray) FindNextByte(nextByte byte) Cursor {
	var res nodeCursorsArray
	for _, cur := range c {
		res = append(res, cur.findNextByte(nextByte)...)
	}

	return res
}
