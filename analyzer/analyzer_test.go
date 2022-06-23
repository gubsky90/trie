package analyzer

import (
	"fmt"
	"testing"

	"github.com/gubsky90/trie"
)

func TestAnalyzer_Do(t *testing.T) {
	analyzer := NewAnalyzer()
	for _, kw := range []string{
		"his",
		"hi",
		"she",
		"he",
	} {
		analyzer.Insert(kw, trie.String(kw))
	}

	fmt.Println(analyzer.Do("hishe"))

	// analyzer.root.Print(os.Stdout)
}

func BenchmarkAnalyzer_Do(b *testing.B) {
	b.ReportAllocs()

	analyzer := NewAnalyzer()
	for _, kw := range []string{
		"his",
		"hi",
		"she",
		"he",
	} {
		analyzer.Insert(kw, trie.String(kw))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Do("hishe")
	}
}
