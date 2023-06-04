package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
)

func Crawl(url string) ([]string, []string, error) {
	// 建立新的 Colly 收集器
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	var titles []string
	var prices []string

	c.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Printf("%v", e)
	})

	// 定義在抓取到目標元素時執行的回呼函式
	c.OnHTML(".listInfo", func(e *colly.HTMLElement) {
		title := e.ChildText(".houseInfo-title")
		price := e.ChildText(".price")
		titles = append(titles, title)
		prices = append(prices, price)
	})

	err := c.Visit(url)
	if err != nil {
		return nil, nil, err
	}

	return titles, prices, nil
}

func EdpRun(url string) {
	options := []chromedp.ExecAllocatorOption{
		// chromedp.Flag("headless", false), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	e, _ := chromedp.NewExecAllocator(context.Background(), options...)
	// 建立 Chromedp 的上下文
	chromeContext, cancel := chromedp.NewContext(e)
	chromedp.Run(chromeContext, make([]chromedp.Action, 0, 1)...)
	defer cancel()
	timeoutCtx, cancel := context.WithTimeout(chromeContext, 20*time.Second)
	defer cancel()
	// 使用 Chromedp 抓取整個頁面的 HTML
	var htmlContent string
	err := chromedp.Run(timeoutCtx, chromedp.Tasks{
		chromedp.Navigate(url), // 替換為目標網站的 URL
		chromedp.Sleep(10 * time.Second),
		chromedp.WaitVisible("body", chromedp.ByQuery), // 等待 body 元素可見
		// chromedp.ActionFunc(func(ctx context.Context) error {
		// 	cookies, err := network.GetAllCookies().Do(ctx)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	for i, cookie := range cookies {
		// 		log.Printf("chrome cookie %d: %+v", i, cookie)
		// 	}

		// 	return nil
		// }),
		chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
	},
	)
	if err != nil {
		log.Fatalf("頁面抓取失敗：%v", err)
	}
	fmt.Printf("%v", htmlContent)

	// 建立新的 Colly 收集器
	c := colly.NewCollector()

	// 在 Colly 的回呼函式中解析 HTML
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Printf("標題：%s\n", title)
	})

	// 將抓取到的 HTML 傳遞給 Colly 解析
	err = c.Visit("data:text/html," + htmlContent)
	if err != nil {
		log.Fatalf("HTML 解析失敗：%v", err)
	}

	fmt.Println("解析完成")
}
