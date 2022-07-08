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

	if assert.Equal(t, 1, len(cur)) {
		ptr := cur[0]
		assert.Equal(t, 1, ptr.offset)
		assert.Equal(t, Prefix{'e', 'n', 'e', 'g', 'r', 'o'}, ptr.current.Prefix)
	}
}

func Test_Node_MultiCursor(t *testing.T) {
	node := &Node{}

	node.Insert([]byte("Monaco"))
	node.Insert([]byte("Montenegro"))
	node.Insert([]byte("Montserrat"))

	ptr, _ := node.Pointer().Move([]byte("Mo"))

	ptr1, _ := ptr.Move([]byte("nte"))
	assert.Equal(t, 1, ptr1.offset)
	assert.Equal(t, Prefix{'e', 'n', 'e', 'g', 'r', 'o'}, ptr1.current.Prefix)

	ptr2, _ := ptr.Move([]byte("na"))
	assert.Equal(t, 1, ptr2.offset)
	assert.Equal(t, Prefix{'a', 'c', 'o'}, ptr2.current.Prefix)
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
	cur.Handle(false, func(v interface{}) {
		fmt.Println("First", v)
	})
	cur.Print(os.Stdout)

	cur = cur.FindNextByte(' ')
	cur.Handle(false, func(v interface{}) {
		fmt.Println("Second", v)
	})
	cur.Print(os.Stdout)

	root.Print(os.Stdout)
}
