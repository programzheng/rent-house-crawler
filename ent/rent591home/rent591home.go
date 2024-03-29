// Code generated by ent, DO NOT EDIT.

package rent591home

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the rent591home type in the database.
	Label = "rent591home"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldPostID holds the string denoting the post_id field in the database.
	FieldPostID = "post_id"
	// FieldKindName holds the string denoting the kind_name field in the database.
	FieldKindName = "kind_name"
	// FieldRoomStr holds the string denoting the room_str field in the database.
	FieldRoomStr = "room_str"
	// FieldFloorStr holds the string denoting the floor_str field in the database.
	FieldFloorStr = "floor_str"
	// FieldCommunity holds the string denoting the community field in the database.
	FieldCommunity = "community"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldPriceUnit holds the string denoting the price_unit field in the database.
	FieldPriceUnit = "price_unit"
	// FieldPhotoList holds the string denoting the photo_list field in the database.
	FieldPhotoList = "photo_list"
	// FieldRegionName holds the string denoting the region_name field in the database.
	FieldRegionName = "region_name"
	// FieldSectionName holds the string denoting the section_name field in the database.
	FieldSectionName = "section_name"
	// FieldStreetName holds the string denoting the street_name field in the database.
	FieldStreetName = "street_name"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldArea holds the string denoting the area field in the database.
	FieldArea = "area"
	// FieldRoleName holds the string denoting the role_name field in the database.
	FieldRoleName = "role_name"
	// FieldContact holds the string denoting the contact field in the database.
	FieldContact = "contact"
	// FieldRefreshTime holds the string denoting the refresh_time field in the database.
	FieldRefreshTime = "refresh_time"
	// FieldYesterdayHit holds the string denoting the yesterday_hit field in the database.
	FieldYesterdayHit = "yesterday_hit"
	// FieldIsVip holds the string denoting the is_vip field in the database.
	FieldIsVip = "is_vip"
	// FieldIsCombine holds the string denoting the is_combine field in the database.
	FieldIsCombine = "is_combine"
	// FieldHurry holds the string denoting the hurry field in the database.
	FieldHurry = "hurry"
	// FieldIsSocial holds the string denoting the is_social field in the database.
	FieldIsSocial = "is_social"
	// FieldDiscountPriceStr holds the string denoting the discount_price_str field in the database.
	FieldDiscountPriceStr = "discount_price_str"
	// FieldCasesID holds the string denoting the cases_id field in the database.
	FieldCasesID = "cases_id"
	// FieldIsVideo holds the string denoting the is_video field in the database.
	FieldIsVideo = "is_video"
	// FieldPreferred holds the string denoting the preferred field in the database.
	FieldPreferred = "preferred"
	// FieldCid holds the string denoting the cid field in the database.
	FieldCid = "cid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeRent591homeTags holds the string denoting the rent591home_tags edge name in mutations.
	EdgeRent591homeTags = "rent591home_tags"
	// EdgeRent591homeSurroundings holds the string denoting the rent591home_surroundings edge name in mutations.
	EdgeRent591homeSurroundings = "rent591home_surroundings"
	// EdgeRent591homeDetails holds the string denoting the rent591home_details edge name in mutations.
	EdgeRent591homeDetails = "rent591home_details"
	// Table holds the table name of the rent591home in the database.
	Table = "rent591homes"
	// Rent591homeTagsTable is the table that holds the rent591home_tags relation/edge. The primary key declared below.
	Rent591homeTagsTable = "rent591home_tag_rent591homes"
	// Rent591homeTagsInverseTable is the table name for the Rent591HomeTag entity.
	// It exists in this package in order to avoid circular dependency with the "rent591hometag" package.
	Rent591homeTagsInverseTable = "rent591home_tags"
	// Rent591homeSurroundingsTable is the table that holds the rent591home_surroundings relation/edge.
	Rent591homeSurroundingsTable = "rent591home_surroundings"
	// Rent591homeSurroundingsInverseTable is the table name for the Rent591HomeSurrounding entity.
	// It exists in this package in order to avoid circular dependency with the "rent591homesurrounding" package.
	Rent591homeSurroundingsInverseTable = "rent591home_surroundings"
	// Rent591homeSurroundingsColumn is the table column denoting the rent591home_surroundings relation/edge.
	Rent591homeSurroundingsColumn = "rent591home_rent591home_surroundings"
	// Rent591homeDetailsTable is the table that holds the rent591home_details relation/edge.
	Rent591homeDetailsTable = "rent591home_details"
	// Rent591homeDetailsInverseTable is the table name for the Rent591HomeDetail entity.
	// It exists in this package in order to avoid circular dependency with the "rent591homedetail" package.
	Rent591homeDetailsInverseTable = "rent591home_details"
	// Rent591homeDetailsColumn is the table column denoting the rent591home_details relation/edge.
	Rent591homeDetailsColumn = "rent591home_rent591home_details"
)

// Columns holds all SQL columns for rent591home fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldType,
	FieldPostID,
	FieldKindName,
	FieldRoomStr,
	FieldFloorStr,
	FieldCommunity,
	FieldPrice,
	FieldPriceUnit,
	FieldPhotoList,
	FieldRegionName,
	FieldSectionName,
	FieldStreetName,
	FieldLocation,
	FieldArea,
	FieldRoleName,
	FieldContact,
	FieldRefreshTime,
	FieldYesterdayHit,
	FieldIsVip,
	FieldIsCombine,
	FieldHurry,
	FieldIsSocial,
	FieldDiscountPriceStr,
	FieldCasesID,
	FieldIsVideo,
	FieldPreferred,
	FieldCid,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// Rent591homeTagsPrimaryKey and Rent591homeTagsColumn2 are the table columns denoting the
	// primary key for the rent591home_tags relation (M2M).
	Rent591homeTagsPrimaryKey = []string{"rent591home_tag_id", "rent591home_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Rent591Home queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByPostID orders the results by the post_id field.
func ByPostID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostID, opts...).ToFunc()
}

// ByKindName orders the results by the kind_name field.
func ByKindName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKindName, opts...).ToFunc()
}

// ByRoomStr orders the results by the room_str field.
func ByRoomStr(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoomStr, opts...).ToFunc()
}

// ByFloorStr orders the results by the floor_str field.
func ByFloorStr(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFloorStr, opts...).ToFunc()
}

// ByCommunity orders the results by the community field.
func ByCommunity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommunity, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByPriceUnit orders the results by the price_unit field.
func ByPriceUnit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPriceUnit, opts...).ToFunc()
}

// ByRegionName orders the results by the region_name field.
func ByRegionName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionName, opts...).ToFunc()
}

// BySectionName orders the results by the section_name field.
func BySectionName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSectionName, opts...).ToFunc()
}

// ByStreetName orders the results by the street_name field.
func ByStreetName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStreetName, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByArea orders the results by the area field.
func ByArea(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArea, opts...).ToFunc()
}

// ByRoleName orders the results by the role_name field.
func ByRoleName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleName, opts...).ToFunc()
}

// ByContact orders the results by the contact field.
func ByContact(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContact, opts...).ToFunc()
}

// ByRefreshTime orders the results by the refresh_time field.
func ByRefreshTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefreshTime, opts...).ToFunc()
}

// ByYesterdayHit orders the results by the yesterday_hit field.
func ByYesterdayHit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYesterdayHit, opts...).ToFunc()
}

// ByIsVip orders the results by the is_vip field.
func ByIsVip(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsVip, opts...).ToFunc()
}

// ByIsCombine orders the results by the is_combine field.
func ByIsCombine(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsCombine, opts...).ToFunc()
}

// ByHurry orders the results by the hurry field.
func ByHurry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHurry, opts...).ToFunc()
}

// ByIsSocial orders the results by the is_social field.
func ByIsSocial(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsSocial, opts...).ToFunc()
}

// ByDiscountPriceStr orders the results by the discount_price_str field.
func ByDiscountPriceStr(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDiscountPriceStr, opts...).ToFunc()
}

// ByCasesID orders the results by the cases_id field.
func ByCasesID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCasesID, opts...).ToFunc()
}

// ByIsVideo orders the results by the is_video field.
func ByIsVideo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsVideo, opts...).ToFunc()
}

// ByPreferred orders the results by the preferred field.
func ByPreferred(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPreferred, opts...).ToFunc()
}

// ByCid orders the results by the cid field.
func ByCid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCid, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByRent591homeTagsCount orders the results by rent591home_tags count.
func ByRent591homeTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRent591homeTagsStep(), opts...)
	}
}

// ByRent591homeTags orders the results by rent591home_tags terms.
func ByRent591homeTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRent591homeTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRent591homeSurroundingsField orders the results by rent591home_surroundings field.
func ByRent591homeSurroundingsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRent591homeSurroundingsStep(), sql.OrderByField(field, opts...))
	}
}

// ByRent591homeDetailsField orders the results by rent591home_details field.
func ByRent591homeDetailsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRent591homeDetailsStep(), sql.OrderByField(field, opts...))
	}
}
func newRent591homeTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Rent591homeTagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, Rent591homeTagsTable, Rent591homeTagsPrimaryKey...),
	)
}
func newRent591homeSurroundingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Rent591homeSurroundingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, Rent591homeSurroundingsTable, Rent591homeSurroundingsColumn),
	)
}
func newRent591homeDetailsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Rent591homeDetailsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, Rent591homeDetailsTable, Rent591homeDetailsColumn),
	)
}
