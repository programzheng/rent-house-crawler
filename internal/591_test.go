package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrawl(t *testing.T) {
	// 建立測試伺服器
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 回傳測試用的範例 HTML 內容
		htmlContent := []byte(`
			<div class="listInfo">
				<div class="houseInfo-title">House 1</div>
				<div class="price">10000</div>
			</div>
			<div class="listInfo">
				<div class="houseInfo-title">House 2</div>
				<div class="price">20000</div>
			</div>
		`)
		w.Write(htmlContent)
	}))
	defer testServer.Close()

	// 使用測試伺服器的 URL 進行測試
	titles, prices, err := Crawl(testServer.URL)
	if err != nil {
		t.Fatalf("爬取失敗URL%s：%v", testServer.URL, err)
	}

	// 驗證結果是否符合預期
	expectedTitles := []string{"House 1", "House 2"}
	expectedPrices := []string{"10000", "20000"}
	for i, title := range titles {
		if title != expectedTitles[i] {
			t.Errorf("標題不符合預期。預期：%s，實際：%s", expectedTitles[i], title)
		}
	}
	for i, price := range prices {
		if price != expectedPrices[i] {
			t.Errorf("價格不符合預期。預期：%s，實際：%s", expectedPrices[i], price)
		}
	}
}
