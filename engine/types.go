package engine

// 对指定url使用指定方法去处理
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// 接收处理结果
type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}
