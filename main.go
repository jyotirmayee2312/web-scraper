package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/html"
)

func main() {
	// URL of the website you want to scrape
	url := "https://www.w3schools.com/"

	// Send an HTTP GET request to the website
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse the HTML response
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Call the scraping function
	scrapeData(doc)
}

func scrapeData(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "h1" {
		// Print the text content of the h1 tag
		fmt.Println(n.FirstChild.Data)
	}

	// Recursively process child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		scrapeData(c)
	}
}
