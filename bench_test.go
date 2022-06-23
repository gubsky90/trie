package trie

import (
	"testing"

	"github.com/gubsky90/trie/testdata"
)

func BenchmarkInsert(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		root := &Node{}
		for _, kw := range testdata.Countries {
			root.Insert([]byte(kw))
		}
	}
}

func BenchmarkFindAll(b *testing.B) {
	b.ReportAllocs()

	data := testdata.Countries

	root := &Node{}
	for _, kw := range data {
		root.Insert([]byte(kw))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root.FindAll([]byte(data[i%len(data)]), func(i interface{}) {})
	}
}
