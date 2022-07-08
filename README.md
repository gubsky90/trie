# Node example

```go
	root := &Node{}

	for _, kw := range []string{
		"his",
		"hi",
		"she",
		"he",
	} {
		root.Insert([]byte(kw), String(kw))
	}

	cur := root.Cursor().Move([]byte("h"))
	cur.Handle(true, func(value interface{}) {
		fmt.Println(value)
	})
```
