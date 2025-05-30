package processor

import (
	"HtmlNodeParsingGo/models"
	"strings"

	"golang.org/x/net/html"
)

// ProcessProductNode: <li> elementi içinden tek bir ürün bilgisi çıkartır
func ProcessProductNode(n *html.Node) models.Product {
	var product models.Product

	var extract func(*html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "h2":
				if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
					product.Name = strings.TrimSpace(n.FirstChild.Data)
				}
			case "img":
				for _, a := range n.Attr {
					if a.Key == "src" {
						product.ImageURL = a.Val
					}
				}
			}
		}

		// Alt düğümleri gezer
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}

	extract(n)
	return product
}
