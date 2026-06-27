package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Results struct {
	Site   string
	Status string
}

func Fetch(site string, wg *sync.WaitGroup, ch chan Results) {
	defer wg.Done()

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error encountered:", err)
		return
	}

	defer resp.Body.Close()

	ch <- Results{
		Site:   site,
		Status: resp.Status,
	}
}

func main() {
	results := make(chan Results)

	sites := []string{
		"https://google.com",
		"https://youtube.com",
		"https://golang.com",
	}

	var wg sync.WaitGroup

	for _, site := range sites {
		wg.Add(1)
		go Fetch(site, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result.Site, "->", result.Status)
	}

}
