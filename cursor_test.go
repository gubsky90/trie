package trie

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Node_Cursor(t *testing.T) {
	node := &Node{}

	node.Insert([]byte("Monaco"))
	node.Insert([]byte("Montenegro"))
	node.Insert([]byte("Montserrat"))

	cur := node.Cursor().Move([]byte("Mo")).Move([]byte("nt")).Move([]byte("e"))

	assert.Equal(t, 1, cur.(*nodeCursor).offset)
	assert.Equal(t, Prefix{'e', 'n', 'e', 'g', 'r', 'o'}, cur.(*nodeCursor).current.Prefix)
}

func Test_Node_MultiCursor(t *testing.T) {
	node := &Node{}

	node.Insert([]byte("Monaco"))
	node.Insert([]byte("Montenegro"))
	node.Insert([]byte("Montserrat"))

	cur := node.Cursor().Move([]byte("Mo"))

	cur1 := cur.Move([]byte("nte"))
	assert.Equal(t, 1, cur1.(*nodeCursor).offset)
	assert.Equal(t, Prefix{'e', 'n', 'e', 'g', 'r', 'o'}, cur1.(*nodeCursor).current.Prefix)

	cur2 := cur.Move([]byte("na"))
	assert.Equal(t, 1, cur2.(*nodeCursor).offset)
	assert.Equal(t, Prefix{'a', 'c', 'o'}, cur2.(*nodeCursor).current.Prefix)
}

func Test_Node_FindNextByte(t *testing.T) {
	root := &Node{}
	for _, kw := range []string{
		"aa",
		"aaaa",
		"aaaaa a",
		"aaaaa aaa a",
	} {
		root.Insert([]byte(kw), String(kw))
	}

	cur := root.Cursor()

	cur = cur.FindNextByte(' ')
	cur.Handle(func(v interface{}) {
		fmt.Println("First", v)
	})

	cur = cur.FindNextByte(' ')
	cur.Handle(func(v interface{}) {
		fmt.Println("Second", v)
	})

	root.Print(os.Stdout)
}
