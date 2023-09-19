package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/programzheng/rent-house-crawler/config"
	"github.com/programzheng/rent-house-crawler/internal/cache"
	"github.com/programzheng/rent-house-crawler/pkg/helper"
)

type CookiesAndCsrfToken struct {
	Cookies   []*http.Cookie
	CsrfToken string
}

type CookiesAndCsrfTokenProxy struct {
	ProxyCookies []*ProxyCookie `json:"cookies"`
	CsrfToken    string         `json:"csrf_token"`
}

type ProxyCookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HomeResponse struct {
	Status  int    `json:"status"`
	Data    Data   `json:"data"`
	Records string `json:"records"`
	IsRecom int    `json:"is_recom"`
}

type Data struct {
	// TopData  []TopData `json:"topData"`
	Biddings []Bidding `json:"biddings"`
	Data     []Datum   `json:"data"`
	Page     string    `json:"page"`
}

type TopData struct {
	Title       string      `json:"title"`
	Type        int         `json:"type"`
	PostID      int         `json:"post_id"`
	Price       string      `json:"price"`
	PriceUnit   string      `json:"price_unit"`
	PhotoList   []string    `json:"photo_list"`
	SectionName string      `json:"section_name"`
	StreetName  string      `json:"street_name"`
	RentTag     []RentTag   `json:"rent_tag"`
	Area        string      `json:"area"`
	Surrounding Surrounding `json:"surrounding"`
	Community   string      `json:"community"`
	RoomStr     string      `json:"room_str"`
	IsVideo     int         `json:"is_video"`
	Preferred   int         `json:"preferred"`
	Kind        int         `json:"kind"`
}

type Bidding struct {
	// define the struct fields for biddings, if needed
}

type Datum struct {
	Title            string      `json:"title"`
	Type             string      `json:"type"`
	PostID           int         `json:"post_id"`
	KindName         string      `json:"kind_name"`
	RoomStr          string      `json:"room_str"`
	FloorStr         string      `json:"floor_str"`
	Community        string      `json:"community"`
	Price            string      `json:"price"`
	PriceUnit        string      `json:"price_unit"`
	PhotoList        []string    `json:"photo_list"`
	SectionName      string      `json:"section_name"`
	StreetName       string      `json:"street_name"`
	Location         string      `json:"location"`
	RentTag          []RentTag   `json:"rent_tag"`
	Area             interface{} `json:"area"`
	RoleName         string      `json:"role_name"`
	Contact          string      `json:"contact"`
	RefreshTime      string      `json:"refresh_time"`
	YesterdayHit     int         `json:"yesterday_hit"`
	IsVIP            int         `json:"is_vip"`
	IsCombine        int         `json:"is_combine"`
	Hurry            int         `json:"hurry"`
	IsSocial         int         `json:"is_socail"`
	Surrounding      Surrounding `json:"surrounding"`
	DiscountPriceStr string      `json:"discount_price_str"`
	CasesID          interface{} `json:"cases_id"`
	IsVideo          int         `json:"is_video"`
	Preferred        int         `json:"preferred"`
	CID              int         `json:"cid"`
}
type RentTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Surrounding struct {
	Type     string `json:"type"`
	Desc     string `json:"desc"`
	Distance string `json:"distance"`
}

type HomeDetailResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type HomeDetailData struct {
	Breadcrumb    []HomeDetailDataBreadcrumb  `json:"breadcrumb"`
	Title         string                      `json:"title"`
	Deposit       string                      `json:"deposit"`
	Kind          int                         `json:"kind"`
	Relieved      int                         `json:"relieved"`
	RegionID      int                         `json:"regionId"`
	SectionID     int                         `json:"sectionId"`
	ShareInfo     HomeDetailDataShareInfo     `json:"shareInfo"`
	DealText      string                      `json:"dealText"`
	DealTime      int                         `json:"dealTime"`
	Browse        HomeDetailDataBrowse        `json:"browse"`
	Tags          []HomeDetailDataTag         `json:"tags"`
	Price         string                      `json:"price"`
	PriceUnit     string                      `json:"priceUnit"`
	NavData       []HomeDetailDataNavData     `json:"navData"`
	Info          []HomeDetailDataInfo        `json:"info"`
	Publish       HomeDetailDataPublish       `json:"publish"`
	PositionRound HomeDetailDataPositionRound `json:"positionRound"`
}

type HomeDetailDataBreadcrumb struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Query string `json:"query"`
	Link  string `json:"link"`
}

type HomeDetailDataShareInfo struct {
	Url   string `json:"url"`
	From  string `json:"from"`
	Title string `json:"title"`
}

type HomeDetailDataBrowse struct {
	Pc     int `json:"pc"`
	Mobile int `json:"mobile"`
}

type HomeDetailDataTag struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type HomeDetailDataNavData struct {
	Title  string `json:"title"`
	Key    string `json:"key"`
	Active int    `json:"active"`
}

type HomeDetailDataInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Key   string `json:"key"`
}

type HomeDetailDataPublish struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Key        string `json:"key"`
	PostTime   string `json:"postTime"`
	UpdateTime string `json:"updateTime"`
}

type HomeDetailDataPositionRound struct {
	Title         string                               `json:"title"`
	Key           string                               `json:"key"`
	Active        int                                  `json:"active"`
	CommunityName string                               `json:"communityName"`
	CommunityID   int                                  `json:"communityId"`
	Address       string                               `json:"address"`
	Lat           string                               `json:"lat"`
	Lng           string                               `json:"lng"`
	Data          []HomeDetailDataPositionRoundData    `json:"data"`
	MapData       []HomeDetailDataPositionRoundMapData `json:"mapData"`
}

type HomeDetailDataPositionRoundData struct {
	Name     string                                    `json:"name"`
	Key      string                                    `json:"key"`
	Children []HomeDetailDataPositionRoundDataChildren `json:"children"`
}

type HomeDetailDataPositionRoundDataChildren struct {
	Type         string `json:"type"`
	Name         string `json:"name"`
	Distance     int    `json:"distance"`
	DistanceText string `json:"distanceText"`
}

type HomeDetailDataPositionRoundMapData struct {
	Name     string                                       `json:"name"`
	Key      string                                       `json:"key"`
	Children []HomeDetailDataPositionRoundMapDataChildren `json:"children"`
}

type HomeDetailDataPositionRoundMapDataChildren struct {
	Name     string                                            `json:"name"`
	Key      string                                            `json:"key"`
	Children []HomeDetailDataPositionRoundDataChildrenChildren `json:"children"`
}

type HomeDetailDataPositionRoundDataChildrenChildren struct {
	Name                string `json:"name"`
	Lat                 string `json:"lat"`
	Lng                 string `json:"lng"`
	Distance            int    `json:"distance"`
	TradingAreaID       int    `json:"trading_area_id"`
	TradingAreaDistance int    `json:"trading_area_distance"`
}

// singleton cookiesAndCsrfToken
var cookiesAndCsrfToken *CookiesAndCsrfToken

const RECORD_OFFSET = 30

func GetDataAndRecordTotalCountByCrawl(url string, startRecordCount int) ([]Datum, int, error) {
	response := crawl(url, startRecordCount)
	recordTotalCountString := strings.ReplaceAll(response.Records, ",", "")
	recordTotalCount, err := strconv.Atoi(recordTotalCountString)
	if err != nil {
		log.Fatalf("591_crawler.go GetDataAndRecordTotalCountByCrawl error: %v", err)
	}
	return response.Data.Data, recordTotalCount, nil
}

func crawl(urlString string, recordCount int) HomeResponse {
	// Get response body at cache
	cacheKey := fmt.Sprintf("%s_%v", helper.ConvertUrlToMd5(urlString), recordCount)
	bodyBytes, err := getCrawlResponseBytesAtCache(cacheKey)
	if len(bodyBytes) == 0 || err != nil {
		parsedURL, err := url.Parse(urlString)
		if err != nil {
			log.Fatalf("591_crawler.go Failed url.Parse error: %v", err)
		}
		queryParams := parsedURL.Query()
		if recordCount > 0 {
			queryParams.Set("firstRow", strconv.Itoa(recordCount))
		}

		_, err = getCookiesAndCsrfToken(config.Cfg.GetString("crawlers.591.cookie-and-csrf-token-url"))
		if err != nil {
			log.Fatalf("591_crawler.go getCookiesAndCsrfToken error: %v", err)
		}

		// Create HTTP client
		client := &http.Client{
			Timeout: time.Second * 10,
		}

		// Create GET request
		req, err := http.NewRequest("GET", urlString, nil)
		if err != nil {
			log.Fatalf("591_crawler.go Failed to create GET request: %v", err)
		}
		req.URL.RawQuery = queryParams.Encode()

		// Set user agent to header
		req.Header.Set("User-Agent", browser.Random())
		// Set the CSRF token
		req.Header.Set("X-CSRF-TOKEN", cookiesAndCsrfToken.CsrfToken)
		// Set the obtained cookies in the request
		for _, cookie := range cookiesAndCsrfToken.Cookies {
			req.AddCookie(cookie)
		}

		time.Sleep(time.Second * 1)

		// Send GET request
		response, err := client.Do(req)
		if err != nil {
			log.Fatalf("591_crawler.go Failed to send GET request: %v", err)
		}
		defer response.Body.Close()

		// Read the response body
		bodyBytes, err = io.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("591_crawler.go Failed to read response body: %v", err)
		}

		if config.Cfg.GetBool("crawlers.591.cache") {
			// Store the response body in cache for 12 hour
			_, err = saveCrawlResponseBytesToCache(cacheKey, bodyBytes, 12*time.Hour)
			if err != nil {
				log.Fatalf("591_crawler.go Failed to save crawl response bytes to cache: %v", err)
			}
		}
	}
	// Create an instance of the Response struct
	var responseData HomeResponse

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		log.Fatalf("591_crawler.go crawl Failed to unmarshal JSON: %s error: %v", string(bodyBytes), err)
	}

	return responseData
}

func getCookiesAndCsrfToken(url string) (*CookiesAndCsrfToken, error) {
	if config.Cfg.GetBool("crawlers.591.debug") {
		defer func() {
			log.Printf("getCookiesAndCsrfToken cookiesAndCsrfToken:%v\n", cookiesAndCsrfToken)
		}()
	}

	// singleton
	if cookiesAndCsrfToken != nil {
		return cookiesAndCsrfToken, nil
	}

	if config.Cfg.GetBool("crawlers.591.proxy") {
		return getCookiesAndCsrfTokenByEdpByProxy(url)
	}

	return getCookiesAndCsrfTokenByEdp(url)
}

func getCookiesAndCsrfTokenByEdpByProxy(url string) (*CookiesAndCsrfToken, error) {
	// Create HTTP client
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("591_crawler.go getCookiesAndCsrfTokenByEdpByProxy Failed to create GET request: %v", err)
	}
	// Set user agent to header
	req.Header.Set("User-Agent", browser.Random())
	// Send GET request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("591_crawler.go getCookiesAndCsrfTokenByEdpByProxy Failed to send GET request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var cookiesAndCsrfTokenProxy *CookiesAndCsrfTokenProxy
	err = json.Unmarshal(bodyBytes, &cookiesAndCsrfTokenProxy)
	if err != nil {
		log.Printf("591_crawler.go getCookiesAndCsrfTokenByEdpByProxy &cookiesAndCsrfTokenProxy Failed to unmarshal JSON: %v", err)
	}

	var requestCookies []*http.Cookie
	for _, cookieProxy := range cookiesAndCsrfTokenProxy.ProxyCookies {
		cookie := &http.Cookie{
			Name:  cookieProxy.Name,
			Value: cookieProxy.Value,
		}
		requestCookies = append(requestCookies, cookie)
	}
	cookiesAndCsrfToken = &CookiesAndCsrfToken{
		Cookies:   requestCookies,
		CsrfToken: cookiesAndCsrfTokenProxy.CsrfToken,
	}

	return cookiesAndCsrfToken, nil
}

func getCookiesAndCsrfTokenByEdp(url string) (*CookiesAndCsrfToken, error) {
	// Create Chromedp context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Execute Chromedp tasks to obtain cookies
	var requestCookies []*http.Cookie
	var htmlContent string
	var csrfToken string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
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
		return nil, err
	}

	cookiesAndCsrfToken = &CookiesAndCsrfToken{
		Cookies:   requestCookies,
		CsrfToken: csrfToken,
	}

	return cookiesAndCsrfToken, nil
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

func getCrawlResponseBytesAtCache(key string) ([]byte, error) {
	if config.Cfg.GetBool("crawlers.591.debug") {
		log.Printf("getCrawlResponseBytesAtCache key:%s\n", key)
	}
	cd, err := cache.GetCacheDriver("")
	if err != nil {
		return []byte{}, err
	}
	return cd.GetBytes(key)
}

func saveCrawlResponseBytesToCache(key string, response []byte, expiration time.Duration) (string, error) {
	cd, err := cache.GetCacheDriver("")
	if err != nil {
		return "", err
	}
	return cd.Set(key, response, expiration)
}

func deleteCrawlResponseBytesByKey(key string) (int64, error) {
	cd, err := cache.GetCacheDriver("")
	if err != nil {
		return 0, err
	}
	return cd.Del(key)
}

func GetHomeDetailDataByPostIDAtCrawl(url string, postID int) (*HomeDetailData, error) {
	hdrd := crawlHomeDetailByPostID(url, postID)
	return hdrd, nil
}

func crawlHomeDetailByPostID(urlString string, postID int) *HomeDetailData {
	// Get response body at cache
	cacheKey := fmt.Sprintf("%s_%d", helper.ConvertUrlToMd5(urlString), postID)
	bodyBytes, err := getCrawlResponseBytesAtCache(cacheKey)
	if len(bodyBytes) == 0 || err != nil {
		parsedURL, err := url.Parse(urlString)
		if err != nil {
			log.Fatalf("591_crawler.go Failed url.Parse error: %v", err)
		}
		queryParams := parsedURL.Query()
		queryParams.Set("id", strconv.Itoa(postID))

		_, err = getCookiesAndCsrfToken(config.Cfg.GetString("crawlers.591.cookie-and-csrf-token-url"))
		if err != nil {
			log.Fatalf("591_crawler.go getCookiesAndCsrfToken error: %v", err)
		}

		// Create HTTP client
		client := &http.Client{
			Timeout: time.Second * 10,
		}

		// Create GET request
		req, err := http.NewRequest("GET", urlString, nil)
		if err != nil {
			log.Fatalf("591_crawler.go Failed to create GET request: %v", err)
		}
		req.URL.RawQuery = queryParams.Encode()

		// Set user agent to header
		req.Header.Set("User-Agent", browser.Random())
		// Set the CSRF token
		req.Header.Set("X-CSRF-TOKEN", cookiesAndCsrfToken.CsrfToken)
		// Set the obtained cookies in the request
		for _, cookie := range cookiesAndCsrfToken.Cookies {
			if cookie.Name == config.Cfg.GetString("crawlers.591.detail-cookie-token-key") {
				req.Header.Set("deviceid", cookie.Value)
			}
			req.AddCookie(cookie)
		}
		// Set device
		req.Header.Set("device", "pc")

		delaySeconds := 1
		if config.Cfg.GetBool("crawlers.591.debug") {
			log.Printf("crawlHomeDetailByPostID post id:%d, waiting %d seconds...\n", postID, delaySeconds)
		}
		time.Sleep(time.Duration(delaySeconds) * time.Second)

		// Send GET request
		response, err := client.Do(req)
		if err != nil {
			log.Fatalf("591_crawler.go Failed to send GET request: %v", err)
		}
		defer response.Body.Close()

		// Read the response body
		bodyBytes, err = io.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("591_crawler.go Failed to read response body: %v", err)
		}

		// Store the response body in cache for 12 hour
		_, err = saveCrawlResponseBytesToCache(cacheKey, bodyBytes, 12*time.Hour)
		if err != nil {
			log.Fatalf("591_crawler.go Failed to save crawl response bytes to cache: %v", err)
		}
	}
	// Create an instance of the Response struct
	var response HomeDetailResponse

	// Unmarshal the JSON data into the HomeDetailResponse struct
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Printf("591_crawler.go &response Failed to unmarshal JSON: %v", err)
	}
	switch response.Data.(type) {
	case map[string]interface{}:
		dataByte, _ := json.Marshal(response.Data)
		var responseData HomeDetailData
		err = json.Unmarshal(dataByte, &responseData)
		if err != nil {
			log.Printf("591_crawler.go &responseData Failed to unmarshal JSON: %v", err)
		}
		return &responseData
	default:
		return nil
	}
}
