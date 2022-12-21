package chatgpt

import (
	"github.com/gomarkdown/markdown"
)

func md2html(mdStr string) []byte {
	mdByte := []byte(mdStr)
	return markdown.ToHTML(mdByte, nil, nil)
}
