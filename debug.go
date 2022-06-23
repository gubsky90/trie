package trie

import (
	"fmt"
	"io"
	"strings"
)

func (node *Node) Print(w io.Writer) {
	if node == nil {
		fmt.Println("<nil>")
	} else {
		node.print("", w)
	}
}

func (node *Node) print(prefix string, w io.Writer) {
	p := node.Prefix[:prefixLength(node.Prefix)]

	fmt.Fprint(w, prefix+string(p))
	node.Values(func(value Comparable) {
		fmt.Print("(", value, ")")
	})

	fmt.Fprintln(w)
	node.Child(func(node *Node) {
		node.print(prefix+strings.Repeat(" ", len(p)), w)
	})
}
