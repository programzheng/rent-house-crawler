package internal

import (
	"context"
	"log"
	"time"

	"github.com/programzheng/rent-house-crawler/config"
	"github.com/programzheng/rent-house-crawler/ent"
	"github.com/programzheng/rent-house-crawler/internal/repository"
	"github.com/programzheng/rent-house-crawler/pkg/helper"
)

var RegionIDs = map[string]int{
	"臺北市": 1,
	"基隆市": 2,
	"新北市": 3,
	"新竹市": 4,
	"新竹縣": 5,
	"桃園市": 6,
	"苗栗縣": 7,
	"台中市": 8,
	"彰化縣": 10,
	"南投縣": 11,
	"嘉義市": 12,
	"嘉義縣": 13,
	"雲林縣": 14,
	"台南市": 15,
	"高雄市": 17,
	"屏東縣": 19,
	"宜蘭縣": 21,
	"台東縣": 22,
	"花蓮縣": 23,
	"澎湖縣": 24,
	"金門縣": 25,
	"連江縣": 26,
}

func SaveHomeData(ctx context.Context) {
	urls := config.Cfg.GetStringSlice("crawlers.591.urls")
	allData := make([]Datum, 0)
	for _, url := range urls {
		_, recordTotalCount, err := GetDataAndRecordTotalCountByCrawl(url, 0)
		if err != nil {
			log.Fatalf("Failed to GetDataByCrawl: %v", err)
		}
		results := make(chan struct{}, (recordTotalCount+RECORD_OFFSET)/RECORD_OFFSET)
		for offset := 0; offset < recordTotalCount+RECORD_OFFSET; offset += RECORD_OFFSET {
			go func(url string, offset int) {
				data, _, err := GetDataAndRecordTotalCountByCrawl(url, offset)
				if err != nil {
					log.Printf("Failed to GetDataByCrawl: %v", err)
					return
				}
				allData = append(allData, data...)

				results <- struct{}{}
			}(url, offset)
			<-results
		}
	}
	SaveRent591Home(ctx, &allData)
}

func SaveRent591Home(ctx context.Context, datum *[]Datum) {
	rp, _ := repository.NewRepository()

	// err := rp.StartTransaction(ctx)
	// if err != nil {
	// 	log.Fatalf("repository StartTransaction error: %v", err)
	// }

	for _, data := range *datum {
		tp, err := helper.ConvertCommaSeparatedStringToInt(data.Type)
		if err != nil {
			log.Fatalf("data.Type ConvertCommaSeparatedStringToInt error: %v", err)
		}

		price, err := helper.ConvertCommaSeparatedStringToInt(data.Price)
		if err != nil {
			log.Fatalf("data.Price ConvertCommaSeparatedStringToInt error: %v", err)
		}

		switch data.CasesID.(type) {
		case string, float64:
			data.CasesID = 0
		}

		r5h, err := rp.UpsertRent591Home(ctx, ent.Rent591Home{
			Title:            data.Title,
			Type:             tp,
			PostID:           data.PostID,
			KindName:         data.KindName,
			RoomStr:          data.RoomStr,
			FloorStr:         data.FloorStr,
			Community:        data.Community,
			Price:            price,
			PriceUnit:        data.PriceUnit,
			PhotoList:        data.PhotoList,
			SectionName:      data.SectionName,
			StreetName:       data.StreetName,
			Location:         data.Location,
			Area:             data.Area.(string),
			RoleName:         data.RoleName,
			Contact:          data.Contact,
			RefreshTime:      data.RefreshTime,
			YesterdayHit:     data.YesterdayHit,
			IsVip:            data.IsVIP,
			IsCombine:        data.IsCombine,
			Hurry:            data.Hurry,
			IsSocial:         data.IsSocial,
			DiscountPriceStr: data.DiscountPriceStr,
			CasesID:          data.CasesID.(int),
			IsVideo:          data.IsVideo,
			Preferred:        data.Preferred,
			Cid:              data.CID,
		})
		if err != nil {
			log.Fatalf("repository CreateRent591Home error: %v", err)
		}
		// fmt.Printf("UpsertRent591Home:%v", r5h)

		r5tsi := []*ent.Rent591HomeTag{}
		for _, rt := range data.RentTag {
			r5tsi = append(r5tsi, &ent.Rent591HomeTag{
				ID:   rt.ID,
				Name: rt.Name,
			})
		}
		_, err = rp.UpdateOrCreateRent591TagsByRent591Home(ctx, r5h, r5tsi)
		if err != nil {
			log.Fatalf("repository CreateRent591TagsByRent591Home error: %v", err)
		}
		// fmt.Printf("UpdateOrCreateRent591TagsByRent591Home:%v", r5ts)

		_, err = rp.UpsertRent591SurroundingByRent591Home(ctx, r5h, ent.Rent591HomeSurrounding{
			Type:     data.Surrounding.Type,
			Desc:     data.Surrounding.Desc,
			Distance: data.Surrounding.Distance,
		})
		if err != nil {
			log.Fatalf("repository UpsertRent591SurroundingByRent591Home error: %v", err)
		}
		// fmt.Printf("UpsertRent591SurroundingByRent591Home:%v", r5s)
	}
}

func UpsertAllHomeByDetailData(ctx context.Context) {
	rp, _ := repository.NewRepository()

	r5hs, err := rp.GetRent591HomePostIDs(ctx)
	if err != nil {
		log.Fatalf("UpsertAllHomeDetailData rp.GetRent591HomePostIDs error: %v", err)
	}
	detailDataUrl := config.Cfg.GetString("crawlers.591.detail-data-url")
	for _, r5h := range r5hs {
		dd, err := GetHomeDetailDataByPostIDAtCrawl(detailDataUrl, r5h.PostID)
		if err != nil {
			log.Printf("GetHomeDetailDataByPostIDAtCrawl post ID:%d error: %v", r5h.PostID, err)
			continue
		}
		if dd == nil {
			rp.DeleteRent591HomeByID(ctx, r5h.ID)
		} else {
			if len(dd.Breadcrumb) > 0 {
				err = rp.SetRent591HomeRegionByID(ctx, r5h.ID, dd.Breadcrumb[0].Name)
				if err != nil {
					log.Printf("SetRent591HomeRegionByPostID post ID:%d, error: %v", r5h.PostID, err)
					continue
				}
				r5hdID, err := rp.UpsertRent591HomeDetailByRent591Home(ctx, r5h, createEntRent591HomeDetail(dd))
				if err != nil {
					log.Printf("UpsertRent591HomeDetailByRent591Home post ID:%d, Retn591HomeDetail ID:%d, error: %v", r5h.PostID, r5hdID, err)
					continue
				}
				err = rp.UpsertRent591HomeDetailBreadcrumbsByRent591HomeDetailID(ctx, r5hdID, createEntRent591HomeDetailBreadcrumb(dd.Breadcrumb))
				if err != nil {
					log.Printf("UpsertRent591HomeDetailBreadcrumbsByRent591HomeDetailID post ID:%d, Retn591HomeDetail ID:%d, error: %v", r5h.PostID, r5hdID, err)
					continue
				}
				_, err = rp.UpsertRent591HomeDetailPositionRoundByRent591HomeDetailID(ctx, r5hdID, createEntRent591HomeDetailPositionRound(dd.PositionRound))
				if err != nil {
					log.Printf("UpsertRent591HomeDetailPositionRoundByRent591HomeDetailID post ID:%d, Retn591HomeDetail ID:%d, error: %v", r5h.PostID, r5hdID, err)
					continue
				}
				_, err = rp.UpsertRent591HomeDetailPublishByRent591HomeDetailID(ctx, r5hdID, createEntRent591HomeDetailPublish(&dd.Publish))
				if err != nil {
					log.Printf("UpsertRent591HomeDetailPublishByRent591HomeDetailID post ID:%d, Retn591HomeDetail ID:%d, error: %v", r5h.PostID, r5hdID, err)
					continue
				}
			}
		}
	}
}

func createEntRent591HomeDetail(dd *HomeDetailData) *ent.Rent591HomeDetail {
	price, err := helper.ConvertCommaSeparatedStringToInt(dd.Price)
	if err != nil {
		log.Fatalf("createEntRent591HomeDetail ConvertCommaSeparatedStringToInt error: %v", err)
	}

	return &ent.Rent591HomeDetail{
		Title:     dd.Title,
		Deposit:   dd.Deposit,
		Kind:      dd.Kind,
		Relieved:  dd.Relieved,
		RegionID:  dd.RegionID,
		SectionID: dd.SectionID,
		DealText:  dd.DealText,
		DealTime:  dd.DealTime,
		Price:     price,
		PriceUnit: dd.PriceUnit,
	}
}

func createEntRent591HomeDetailBreadcrumb(ddbs []HomeDetailDataBreadcrumb) *ent.Rent591HomeDetailBreadcrumbs {
	r5hdbs := ent.Rent591HomeDetailBreadcrumbs{}
	for _, ddb := range ddbs {
		r5hdbs = append(r5hdbs, &ent.Rent591HomeDetailBreadcrumb{
			Name: ddb.Name,
			// original id
			PostID: ddb.ID,
			Query:  ddb.Query,
			Link:   ddb.Link,
		})
	}
	return &r5hdbs
}

func createEntRent591HomeDetailPositionRound(ddpr HomeDetailDataPositionRound) *ent.Rent591HomeDetailPositionRound {
	lat, err := helper.ConvertStringToFloat64(ddpr.Lat)
	if err != nil {
		log.Fatalf("createEntRent591HomeDetailPositionRound Lat ConvertStringToFloat64 error: %v", err)
	}

	lng, err := helper.ConvertStringToFloat64(ddpr.Lng)
	if err != nil {
		log.Fatalf("createEntRent591HomeDetailPositionRound Lng ConvertStringToFloat64 error: %v", err)
	}

	return &ent.Rent591HomeDetailPositionRound{
		Title:         ddpr.Title,
		Key:           ddpr.Key,
		Active:        ddpr.Active,
		CommunityName: ddpr.CommunityName,
		CommunityID:   ddpr.CommunityID,
		Address:       ddpr.Address,
		Lat:           lat,
		Lng:           lng,
	}
}

func createEntRent591HomeDetailPublish(ddp *HomeDetailDataPublish) *ent.Rent591HomeDetailPublish {
	return &ent.Rent591HomeDetailPublish{
		PostID:     ddp.ID,
		Name:       ddp.Name,
		Key:        ddp.Key,
		PostTime:   ddp.PostTime,
		UpdateTime: ddp.UpdateTime,
	}
}

func GetRent591HomesFilterCreatedAtByStartTimestampAndEndTimestamp(ctx context.Context, startTimestamp *time.Time, endTimestamp *time.Time) ([]*ent.Rent591Home, error) {
	rp, err := repository.NewRepository()
	if err != nil {
		return nil, err
	}
	return rp.GetRent591HomesFilterCreatedAtByStartTimestampAndEndTimestamp(ctx, startTimestamp, endTimestamp)
}

func GetNewRent591HomesByCity(ctx context.Context, city *string) ([]*ent.Rent591Home, error) {
	rp, err := repository.NewRepository()
	if err != nil {
		return nil, err
	}
	return rp.GetNewRent591HomesByRegionID(ctx, 1)
}
