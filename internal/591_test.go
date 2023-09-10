package internal

import (
	"context"
	"testing"
)

func TestSaveRent591Home(t *testing.T) {
	datum := &[]Datum{
		{
			Title:     "3樓獨立套房出租",
			Type:      "1",
			PostID:    1,
			KindName:  "獨立套房",
			RoomStr:   "",
			FloorStr:  "3F/4F",
			Community: "",
			Price:     "1,000",
			PriceUnit: "元/月",
			PhotoList: []string{
				"1.jpg",
				"2.jpg",
				"3.jpg",
				"4.jpg",
				"5.jpg",
			},
			SectionName:      "內湖區",
			StreetName:       "測試路三段",
			Location:         "內湖區-測試路x段xx巷",
			Area:             "6",
			RoleName:         "代理人",
			Contact:          "王先生",
			RefreshTime:      "10小時內",
			YesterdayHit:     517,
			IsVIP:            0,
			IsCombine:        2,
			Hurry:            0,
			IsSocial:         0,
			DiscountPriceStr: "",
			CasesID:          1.0,
			IsVideo:          0,
			Preferred:        0,
			CID:              0,
		},
		{
			Title:     "2樓獨立套房出租",
			Type:      "1",
			PostID:    2,
			KindName:  "獨立套房",
			RoomStr:   "",
			FloorStr:  "2F/4F",
			Community: "",
			Price:     "1,113",
			PriceUnit: "元/月",
			PhotoList: []string{
				"1.jpg",
				"2.jpg",
				"3.jpg",
				"4.jpg",
				"5.jpg",
			},
			SectionName:      "中和區",
			StreetName:       "測試路四段",
			Location:         "中和區-測試路x段xx巷",
			Area:             "6",
			RoleName:         "代理人",
			Contact:          "林先生",
			RefreshTime:      "3小時內",
			YesterdayHit:     111,
			IsVIP:            0,
			IsCombine:        2,
			Hurry:            0,
			IsSocial:         0,
			DiscountPriceStr: "",
			CasesID:          "",
			IsVideo:          0,
			Preferred:        0,
			CID:              0,
		},
	}

	SaveRent591Home(context.Background(), datum)
}
