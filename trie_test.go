package trie_test

import (
	"testing"

	"github.com/gubsky90/trie"
	"github.com/gubsky90/trie/testdata"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	tree := &trie.Trie{}
	for _, kw := range testdata.Countries {
		tree.Insert(kw, trie.String(kw))
	}

	var res []string
	tree.FindAll("Mon", func(value interface{}) {
		res = append(res, string(value.(trie.String)))
	})

	assert.ElementsMatch(t, []string{
		"Monaco",
		"Mongolia",
		"Montenegro",
		"Montserrat",
	}, res)
}
