package main

import (
	"HtmlNodeParsingGo/controller"
	"HtmlNodeParsingGo/parser"
	"HtmlNodeParsingGo/processor"
	"encoding/json"
	"fmt"
)

func main() {
	url := "https://www.scrapingcourse.com/ecommerce/"

	// HTML verisini çek
	htmlData, err := controller.FetchHtml(url)
	if err != nil {
		fmt.Println("HTML alma hatası:", err)
		return
	}

	// HTML parse et
	doc, err := parser.ParseHTML(htmlData)
	if err != nil {
		fmt.Println("Parse hatası:", err)
		return
	}

	// Ürünleri işle
	products := processor.ProcessAllProducts(doc)

	// JSON çıktısı ver
	jsonOutput, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		fmt.Println("JSON'a çevirme hatası:", err)
		return
	}

	fmt.Println(string(jsonOutput))
}
