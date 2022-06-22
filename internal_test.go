package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_prefixLength(t *testing.T) {
	assert.Equal(t, 2, prefixLength(Prefix{1, 1}))
	assert.Equal(t, 0, prefixLength(Prefix{}))
	assert.Equal(t, 7, prefixLength(Prefix{1, 1, 1, 1, 1, 1, 1}))
}

func Test_equalPrefixLength(t *testing.T) {
	assert.Equal(t, 0, equalPrefixLength([]byte(""), []byte("")))
	assert.Equal(t, 0, equalPrefixLength([]byte("abc"), []byte("")))
	assert.Equal(t, 0, equalPrefixLength([]byte(""), []byte("abc")))
	assert.Equal(t, 3, equalPrefixLength([]byte("abc"), []byte("abc")))
	assert.Equal(t, 3, equalPrefixLength([]byte("abc"), []byte("abcd")))
	assert.Equal(t, 3, equalPrefixLength([]byte("abcd"), []byte("abc")))
	assert.Equal(t, 3, equalPrefixLength([]byte("abcd"), []byte("abcv")))
	assert.Equal(t, 0, equalPrefixLength([]byte("rabcd"), []byte("sabcd")))
}
