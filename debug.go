package trie

import (
	"fmt"
	"io"
	"strings"
)

func (t *Node) Print(w io.Writer) {
	t.print("", w)
}

func (t *Trie) Print(w io.Writer) {
	t.root.Print(w)
}

func (t *Node) print(prefix string, w io.Writer) {
	p := t.Prefix[:prefixLength(t.Prefix)]
	fmt.Fprint(w, prefix+string(p))
	for _, v := range t.Values() {
		fmt.Print("(", v, ")")
	}
	fmt.Fprintln(w)
	for _, n := range t.Nodes() {
		n.print(prefix+strings.Repeat(" ", len(p)), w)
	}
}
