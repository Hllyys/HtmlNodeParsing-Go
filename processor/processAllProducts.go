package processor

import (
	"HtmlNodeParsingGo/models"

	"golang.org/x/net/html"
)

// Ürün listesini döndürür
func ProcessAllProducts(n *html.Node) []models.Product {
	var products []models.Product

	// HTML içinde <li> etiketleri olan ürünler
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "li" {
			product := ProcessProductNode(node)
			if product.Name != "" {
				products = append(products, product)
			}
		}
		// Tüm çocukları gez
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(n)
	return products
}
