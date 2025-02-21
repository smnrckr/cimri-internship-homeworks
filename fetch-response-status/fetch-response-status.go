package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func main() {
	var wg sync.WaitGroup
	maxGoroutines := 5

	semaphore := make(chan struct{}, maxGoroutines)

	limit := make(chan website_info, maxGoroutines)

	website_urls := [...]string{"https://www.cimri.com", "https://www.trendyol.com", "https://www.hepsiburada.com", "https://www.n11.com", "https://www.amazon.com", "https://www.idefix.com", "https://www.dr.com.tr", "https://www.nadirkitap.com", "https://www.kitapyurdu.com"}

	go func() {
		for i, url := range website_urls {
			wg.Add(1)
			//5de bloklanmalı
			fmt.Println(i)
			semaphore <- struct{}{}
			go func(url string) {
				defer func() { <-semaphore }()
				defer wg.Done()
				fetch_url(url, limit, &wg)
			}(url)
		}

	}()

	go func() {
		wg.Wait()
		close(limit)
	}()

	err := save(limit)
	if err != nil {
		panic(err)
	}
	fmt.Println("bitti")
}

type website_info struct {
	url   string
	title string
}

func fetch_url(url string, limit chan website_info, wg *sync.WaitGroup) {
	//result := make(map[string]string)
	result := website_info{}

	fmt.Println("........")
	time.Sleep(time.Second)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result.title = getTitle(resp.Body)
	result.url = url

	limit <- result
	fmt.Println("fetch url end")
}

func getTitle(r io.Reader) string {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	title, _ := traverse(doc)
	return title

}
func traverse(doc *html.Node) (string, bool) {
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "title" {
			return c.FirstChild.Data, true
		}
		res, ok := traverse(c)
		if ok {
			return res, true
		}
	}
	return "", false

}

func save(liste chan website_info) error {
	f, err := os.Create("url_files.txt")
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	for result := range liste {
		s := fmt.Sprintf("%s : %s", result.url, result.title)
		fmt.Fprintln(f, s)

	}
	f.Close()

	return nil
}
