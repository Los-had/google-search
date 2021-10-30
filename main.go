package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Welcome to google search on your terminal!")
}

func GenerateURL(search_term string) string {
	search_term = strings.Replace(search_term, " ", "+", -1)

	return search_term
}