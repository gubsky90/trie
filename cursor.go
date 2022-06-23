package trie

type Cursor interface {
	Empty() bool
	Move(prefix []byte) Cursor
	FindNextByte(nextByte byte) Cursor
	Handle(hf HandlerFunc)
}

type nodeCursor struct {
	current, parent *Node
	offset          int
}

func (c nodeCursor) Empty() bool {
	return c.current == nil
}

func (c *nodeCursor) Move(prefix []byte) Cursor {
	return c.move(prefix)
}
func (c *nodeCursor) Clone() Cursor {
	return c.clone()
}

func (c *nodeCursor) clone() *nodeCursor {
	return &nodeCursor{
		offset:  c.offset,
		current: c.current,
		parent:  c.parent,
	}
}

func (c *nodeCursor) move(prefix []byte) *nodeCursor {
	newCursor := c.clone()
	for len(prefix) > 0 {
		if newCursor.current == nil {
			return nil
		}

		if l := equalPrefixLength(newCursor.current.Prefix[newCursor.offset:], prefix); l > 0 {
			newCursor.offset += l
			prefix = prefix[l:]
		} else {
			newCursor.parent = newCursor.current
			newCursor.current, newCursor.offset = newCursor.current.findPrefix(prefix)
			prefix = prefix[newCursor.offset:]
		}
	}

	return newCursor
}

func (c nodeCursor) FindNextByte(nextByte byte) Cursor {
	if !c.current.Root() {
		for i, b := range c.current.Prefix[c.offset:] {
			if b == nextByte {
				newCursor := c.clone()
				newCursor.offset += i
				return newCursor
			}
		}
	}

	var cursors nodeCursorsArray
	c.current.Child(func(node *Node) {
		cursors = append(cursors, nodeCursor{current: node}.findNextByte(nextByte)...)
	})

	return cursors
}

func (c *nodeCursor) Handle(hf HandlerFunc) {
	c.current.Handle(hf)
}
