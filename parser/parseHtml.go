package parser

import (
	"strings"

	"golang.org/x/net/html"
)

func ParseHTML(htmlStr string) (*html.Node, error) {
	return html.Parse(strings.NewReader(htmlStr))
}
