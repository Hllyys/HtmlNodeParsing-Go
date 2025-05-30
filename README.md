This project is a simple scraper application written in Go that fetches product data from an e-commerce webpage and processes it. The HTML content of the page is handled using the net/http and golang.org/x/net/html packages.

The goal is to extract from each <li> element the product's:

Title (from the <h2> tag)

Image URL (from the <img src="..." /> tag)

and return it as a structured object.
