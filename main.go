package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strconv"
	"strings"
	"github.com/gocolly/colly"
)

type SearchResult struct {
	Link string
	Name string
	Description string
}

func main() {
	fmt.Println("Welcome to google search on your terminal!")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Search: ")
	scanner.Scan()
	input := scanner.Text()
	
	// Search on google
	url := "www.google.com/search?q=" + GenerateURL(input) + "&ie=UTF-8&oe=UTF-8"
	url2 := "google.com/search?q=" + GenerateURL(input) + "&ie=UTF-8&oe=UTF-8"
	i := 0
	//results := make([]SearchResult, 0)

	c := colly.NewCollector(
		colly.AllowedDomains(url, url2),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/61.0.3163.100 Safari/537.36"),
	)

	c.OnHTML(".g div ", func(e *colly.HTMLElement) {
		//return
		i++
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatal("Error:", err)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Scraping", request.URL.String())
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(i, "Results was founded")
	})

	c.Visit(url)
}

func GenerateURL(search_term string) string {
	search_term = strings.Replace(search_term, " ", "+", -1)

	return search_term
}
