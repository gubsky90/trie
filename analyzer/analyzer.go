package analyzer

import (
	"os"

	"github.com/gubsky90/trie"
)

type (
	Analyzer struct {
		root     *trie.Node
		splitter Splitter
	}

	Fragment struct {
		Bytes []byte
		Start int
		End   int
	}

	Splitter func(string) []Fragment
)

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		root: &trie.Node{},
	}
}

func (a *Analyzer) Insert(kw string, value ...trie.Comparable) {
	a.root.Insert([]byte(kw), value...)
}

func (a *Analyzer) Do(q string) {
	cursors := a.root.Cursor()

	for i := 0; i < len(q); i++ {
		old := cursors

		cursors = old.Move([]byte{q[i]})

		if q[i] == 'o' {
			cursors = append(cursors, old.Move([]byte{'a'})...)
		}

		if q[i] == ' ' {
			cursors = append(cursors, a.root.Cursor()...)
		}
	}

	cursors.Print(os.Stdout)
}
