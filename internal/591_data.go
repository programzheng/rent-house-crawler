package internal

type Response struct {
	Status  int    `json:"status"`
	Data    Data   `json:"data"`
	Records string `json:"records"`
	IsRecom int    `json:"is_recom"`
	// other fields...
}

type Data struct {
	TopData  []TopData `json:"topData"`
	Biddings []Bidding `json:"biddings"`
	Data     []Datum   `json:"data"`
	Page     string    `json:"page"`
	// other fields...
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
	// other fields...
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
	Area             string      `json:"area"`
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
	CasesID          int         `json:"cases_id"`
	IsVideo          int         `json:"is_video"`
	Preferred        int         `json:"preferred"`
	CID              int         `json:"cid"`
	// other fields...
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
