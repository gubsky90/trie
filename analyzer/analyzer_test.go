package analyzer

import (
	"testing"

	"github.com/gubsky90/trie/testdata"
)

func TestAnalyzer_Do(t *testing.T) {
	analyzer := NewAnalyzer()
	for _, kw := range testdata.Countries {
		analyzer.Insert(kw)
	}

	a := &Analyzer{}

	a.Do("hello from Ukraine!!!")
}
