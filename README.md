# HtmlNodeParsing-Go

Extracting HTML format from a web page is a modular Go application that parses the product.

## 📌Project Description

This project is a simple scraper application that uses the `Go` language to scrape the product from an e-commerce page. The HTML data on the page is processed with the `net/http` and `golang.org/x/net/html` packages.

The goal is to extract the product within its `<li>` tag:

- Title (`<h2>`)
- Image URL (`<img src="..." />`)

and return it as an object.
