package message

//Text 文本消息
type Text struct {
	CommonContent
	TextContent string `json:"text"`
	Safe        int    `json:"safe"`
}

// TextContent 文本内容
type TextContent struct {
	Content string `json:"content"`
}

//NewText 初始化文本消息
func NewText(content string) *Text {
	text := new(Text)
	text.TextContent = TextContent{
		Content: content,
	}
	return text
}
