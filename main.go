package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"encoding/base64"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please specify a file with a list of URLs")
		return
	}

	urlsFile, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	urls := strings.Split(string(urlsFile), "\n")

	maxWorkers := 5 // Maximum number of concurrent workers

	var wg sync.WaitGroup
	urlChan := make(chan string)

	// Start worker goroutines
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, urlChan)
	}

	// Push URLs into the channel
	for _, u := range urls {
		u = strings.TrimSpace(u)
		if u == "" {
			continue
		}
		urlChan <- u
	}
	close(urlChan)

	wg.Wait()

	fmt.Println("Screenshot capture complete")
}

func worker(ctx context.Context, wg *sync.WaitGroup, urlChan <-chan string) {
	defer wg.Done()

	for urlStr := range urlChan {
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
			log.Printf("Invalid URL: %s - %s\n", urlStr, err)
			continue
		}

		// hostname := parsedURL.Hostname()

		var buf []byte

		err = chromedp.Run(ctx,
			chromedp.Navigate(urlStr),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
			chromedp.CaptureScreenshot(&buf),
		)
		if err != nil {
			log.Printf("Failed to capture screenshot for URL: %s - %s\n", urlStr, err)
			continue
		}

		full_url:=parsedURL.String()
		bytes :=[]byte(full_url)
		encoded := base64.StdEncoding.EncodeToString(bytes)
		encoded = strings.ReplaceAll(encoded, `"`, "")
		encoded = strings.TrimRight(encoded, "=")


		err = ioutil.WriteFile(fmt.Sprintf("%s.jpg", encoded), buf, 0644)
		if err != nil {
			log.Printf("Failed to save screenshot for URL: %s - %s\n", urlStr, err)
		} else {
			log.Printf("Screenshot saved for URL: %s\n", urlStr)
		}
	}
}
