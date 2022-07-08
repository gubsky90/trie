package trie

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Node_findPrefix(t *testing.T) {
	node := Node{
		Next: []Comparable{
			&Node{Prefix: Prefix{'s', 'u', 'f', 'i', 'x'}},
			&Node{Prefix: Prefix{'p', 'r', 'e', 'f', 'i', 'x'}},
		},
	}

	{
		r, l := node.findPrefix([]byte("pre"))
		if assert.NotNil(t, r) {
			assert.Equal(t, Prefix{'p', 'r', 'e', 'f', 'i', 'x'}, r.Prefix)
			assert.Equal(t, 3, l)
		}
	}

	{
		r, l := node.findPrefix([]byte("sufixes"))
		if assert.NotNil(t, r) {
			assert.Equal(t, Prefix{'s', 'u', 'f', 'i', 'x'}, r.Prefix)
			assert.Equal(t, 5, l)
		}
	}

	{
		r, l := node.findPrefix([]byte("any"))
		assert.Nil(t, r)
		assert.Equal(t, 0, l)
	}
}

func Test_Node_Insert(t *testing.T) {
	node := &Node{}

	node.Insert([]byte("Monaco"))
	node.Insert([]byte("Montenegro"))
	node.Insert([]byte("Montserrat"))

	assert.Equal(t, &Node{
		Next: []Comparable{
			&Node{Prefix: Prefix{'M', 'o', 'n'}, Next: []Comparable{
				&Node{Prefix: Prefix{'a', 'c', 'o'}},
				&Node{Prefix: Prefix{'t'}, Next: []Comparable{
					&Node{Prefix: Prefix{'e', 'n', 'e', 'g', 'r', 'o'}},
					&Node{Prefix: Prefix{'s', 'e', 'r', 'r', 'a', 't'}},
				}},
			}},
		},
	}, node)
}

func Test_Node_Insert2(t *testing.T) {
	node := &Node{}

	for _, kw := range []string{
		"his",
		"hi",
		"she",
		"he",
	} {
		node.Insert([]byte(kw), String(kw))
	}

	assert.Equal(t, &Node{
		Next: []Comparable{
			&Node{Prefix: Prefix{'h'}, Next: []Comparable{
				&Node{Prefix: Prefix{'i'}, Next: []Comparable{
					&Node{Prefix: Prefix{'s'}, Next: []Comparable{
						String("his"),
					}},
					String("hi"),
				}},
				&Node{Prefix: Prefix{'e'}, Next: []Comparable{
					String("he"),
				}},
			}},
			&Node{Prefix: Prefix{'s', 'h', 'e'}, Next: []Comparable{
				String("she"),
			}},
		},
	}, node)
}

func Test_Node_Delete(t *testing.T) {
	node := &Node{}

	node.Insert([]byte("Monaco"))
	node.Insert([]byte("Montenegro"))
	node.Insert([]byte("Montserrat"))

	node.Delete([]byte("Mont"))

	assert.Equal(t, &Node{
		Next: []Comparable{
			&Node{Prefix: Prefix{'M', 'o', 'n'}, Next: []Comparable{
				&Node{Prefix: Prefix{'a', 'c', 'o'}},
			}},
		},
	}, node)
}

func Test_Node_Delete2(t *testing.T) {
	root := &Node{}

	for _, kw := range []string{
		"his",
		"hi",
		"she",
		"he",
	} {
		root.Insert([]byte(kw), String(kw))
	}

	assert.Equal(t,
		` >
 >h
  >i(hi)
   >s(his)
  >e(he)
 >she(she)
`, root.String())

	root.Delete([]byte("hi"))

	assert.Equal(t,
		` >
 >h
  >is(his)
  >e(he)
 >she(she)
`, root.String())
}

func Test_Node_InsertExample(t *testing.T) {
	root := &Node{}

	for _, kw := range []string{
		"his",
		"hi",
		"she",
		"he",
	} {
		root.Insert([]byte(kw), String(kw))
	}

	root.Delete([]byte("hi"))
	root.Print(os.Stdout)

	cur := root.Cursor().Move([]byte("h"))
	cur.Handle(true, func(value interface{}) {
		fmt.Println(value)
	})
}
