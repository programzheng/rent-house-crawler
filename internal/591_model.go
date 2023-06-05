package internal

type Rent591Home struct {
	Title            string       `json:"title"`
	Type             string       `json:"type"`
	PostID           int          `json:"post_id"`
	KindName         string       `json:"kind_name"`
	RoomStr          string       `json:"room_str"`
	FloorStr         string       `json:"floor_str"`
	Community        string       `json:"community"`
	Price            string       `json:"price"`
	PriceUnit        string       `json:"price_unit"`
	PhotoList        []string     `json:"photo_list"`
	SectionName      string       `json:"section_name"`
	StreetName       string       `json:"street_name"`
	Location         string       `json:"location"`
	RentTag          []Rent591Tag `json:"rent_tag"`
	Area             interface{}  `json:"area"`
	RoleName         string       `json:"role_name"`
	Contact          string       `json:"contact"`
	RefreshTime      string       `json:"refresh_time"`
	YesterdayHit     int          `json:"yesterday_hit"`
	IsVIP            int          `json:"is_vip"`
	IsCombine        int          `json:"is_combine"`
	Hurry            int          `json:"hurry"`
	IsSocial         int          `json:"is_socail"`
	Surrounding      Surrounding  `json:"surrounding"`
	DiscountPriceStr string       `json:"discount_price_str"`
	CasesID          interface{}  `json:"cases_id"`
	IsVideo          int          `json:"is_video"`
	Preferred        int          `json:"preferred"`
	CID              int          `json:"cid"`
}

type Rent591Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Rent591Surrounding struct {
	Type     string `json:"type"`
	Desc     string `json:"desc"`
	Distance string `json:"distance"`
}
