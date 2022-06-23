package trie

import (
	"testing"

	"github.com/gubsky90/trie/testdata"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	root := &Node{}
	for _, kw := range testdata.Countries {
		root.Insert([]byte(kw), String(kw))
	}

	var res []string
	root.FindAll([]byte("Mon"), func(value interface{}) {
		res = append(res, string(value.(String)))
	})

	assert.ElementsMatch(t, []string{
		"Monaco",
		"Mongolia",
		"Montenegro",
		"Montserrat",
	}, res)
}
