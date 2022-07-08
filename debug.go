package trie

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func (node *Node) String() string {
	buf := bytes.NewBuffer(nil)
	node.Print(buf)
	return buf.String()
}

func (node *Node) Print(w io.Writer) {
	if node == nil {
		fmt.Println("<nil>")
	} else {
		node.print(" ", w)
	}
}

func (node *Node) print(prefix string, w io.Writer) {
	p := node.Prefix[:prefixLength(node.Prefix)]
	fmt.Fprint(w, prefix+">"+string(p))

	node.Values(func(value Comparable) {
		fmt.Fprint(w, "(", value, ")")
	})

	fmt.Fprintln(w)
	node.Child(func(node *Node) {
		node.print(prefix+strings.Repeat(" ", len(p)), w)
	})
}

func (p Pointer) Print(w io.Writer) {
	if p.current == nil {
		fmt.Fprintf(w, "ERROR %#v\n", p)
	} else {
		fmt.Fprintln(w, string(p.current.Prefix[:]), "offset", p.offset)
	}
}

func (c Cursor) Print(w io.Writer) {
	for _, p := range c {
		p.Print(w)
	}
}
