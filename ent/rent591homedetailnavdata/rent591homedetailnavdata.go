// Code generated by ent, DO NOT EDIT.

package rent591homedetailnavdata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the rent591homedetailnavdata type in the database.
	Label = "rent591home_detail_nav_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// EdgeRent591homeDetails holds the string denoting the rent591home_details edge name in mutations.
	EdgeRent591homeDetails = "rent591home_details"
	// Table holds the table name of the rent591homedetailnavdata in the database.
	Table = "rent591home_detail_nav_data"
	// Rent591homeDetailsTable is the table that holds the rent591home_details relation/edge. The primary key declared below.
	Rent591homeDetailsTable = "rent591home_detail_nav_data_rent591home_details"
	// Rent591homeDetailsInverseTable is the table name for the Rent591HomeDetail entity.
	// It exists in this package in order to avoid circular dependency with the "rent591homedetail" package.
	Rent591homeDetailsInverseTable = "rent591home_details"
)

// Columns holds all SQL columns for rent591homedetailnavdata fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldKey,
	FieldActive,
}

var (
	// Rent591homeDetailsPrimaryKey and Rent591homeDetailsColumn2 are the table columns denoting the
	// primary key for the rent591home_details relation (M2M).
	Rent591homeDetailsPrimaryKey = []string{"rent591home_detail_nav_data_id", "rent591home_detail_id"}
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

// OrderOption defines the ordering options for the Rent591HomeDetailNavData queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByRent591homeDetailsCount orders the results by rent591home_details count.
func ByRent591homeDetailsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRent591homeDetailsStep(), opts...)
	}
}

// ByRent591homeDetails orders the results by rent591home_details terms.
func ByRent591homeDetails(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRent591homeDetailsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRent591homeDetailsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Rent591homeDetailsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, Rent591homeDetailsTable, Rent591homeDetailsPrimaryKey...),
	)
}
