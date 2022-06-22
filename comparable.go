package trie

import "strings"

type (
	String string
)

func (a String) Compare(b interface{}) int {
	if b, ok := b.(string); ok {
		return strings.Compare(string(a), b)
	}

	return 0
}
