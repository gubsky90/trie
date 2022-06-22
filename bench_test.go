package trie_test

import (
	"testing"

	"github.com/gubsky90/trie"
	"github.com/gubsky90/trie/testdata"
)

func BenchmarkInsert(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tree := &trie.Trie{}
		for _, kw := range testdata.Countries {
			tree.Insert(kw)
		}
	}
}

func BenchmarkFindAll(b *testing.B) {
	b.ReportAllocs()

	data := testdata.Countries

	tree := &trie.Trie{}
	for _, kw := range data {
		tree.Insert(kw)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.FindAll(data[i%len(data)], func(i interface{}) {})
	}
}
