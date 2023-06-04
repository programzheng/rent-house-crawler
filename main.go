package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	// Create Chromedp context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Execute Chromedp tasks to obtain cookies
	var requestCookies []*http.Cookie
	var htmlContent string
	var csrfToken string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://rent.591.com.tw"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetCookies().Do(ctx)
			if err != nil {
				return err
			}
			for _, c := range cookies {
				cookie := &http.Cookie{
					Name:  c.Name,
					Value: c.Value,
				}

				requestCookies = append(requestCookies, cookie)
			}

			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Get the HTML of the current page
			err := chromedp.EvaluateAsDevTools(`document.documentElement.outerHTML`, &htmlContent).Do(ctx)
			if err != nil {
				return err
			}

			// Parse the CSRF token from the HTML
			csrfToken = parseCSRFToken(htmlContent)
			return nil
		}),
	)
	if err != nil {
		log.Fatalf("Chromedp error: %v", err)
	}

	// Create HTTP client
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Create GET request
	req, err := http.NewRequest("GET", "https://rent.591.com.tw/home/search/rsList?is_format_data=1&is_new_list=1&type=1&kind=2&multiPrice=5000_10000", nil)
	if err != nil {
		log.Fatalf("Failed to create GET request: %v", err)
	}
	// Set the CSRF token
	req.Header.Set("X-CSRF-TOKEN", csrfToken)
	// Set the obtained cookies in the request
	for _, cookie := range requestCookies {
		req.AddCookie(cookie)
	}

	// Send GET request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send GET request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Convert byte slice to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

// Parse the CSRF token from the HTML
func parseCSRFToken(html string) string {
	const metaTag = `<meta name="csrf-token" content="`
	startIndex := strings.Index(html, metaTag)
	if startIndex == -1 {
		return ""
	}

	startIndex += len(metaTag)
	endIndex := strings.Index(html[startIndex:], `"`)
	if endIndex == -1 {
		return ""
	}

	return html[startIndex : startIndex+endIndex]
}
