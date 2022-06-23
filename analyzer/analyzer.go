package analyzer

import (
	"github.com/gubsky90/trie"
)

type (
	Analyzer struct {
		root *trie.Node
	}
)

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		root: &trie.Node{},
	}
}

func (a *Analyzer) Insert(kw string, value ...trie.Comparable) {
	a.root.Insert([]byte(kw), value...)
}

func (a *Analyzer) Do(q string) (res []interface{}) {
	cursor := a.root.Cursor()

	for i := 0; i < len(q); i++ {
		old := cursor
		cursor = old.Move([]byte{q[i]})

		for _, p := range cursor {
			p.Handle(func(value interface{}) {
				res = append(res, value)
			})
		}

		cursor = append(cursor, a.root.Pointer())

		if q[i] == ' ' {
			cursor = append(cursor, old...)
		}
	}

	return
}
