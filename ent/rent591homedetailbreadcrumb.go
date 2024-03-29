// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailbreadcrumb"
)

// Rent591HomeDetailBreadcrumb is the model entity for the Rent591HomeDetailBreadcrumb schema.
type Rent591HomeDetailBreadcrumb struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// PostID holds the value of the "post_id" field.
	PostID int `json:"post_id,omitempty"`
	// Query holds the value of the "query" field.
	Query string `json:"query,omitempty"`
	// Link holds the value of the "link" field.
	Link string `json:"link,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Rent591HomeDetailBreadcrumbQuery when eager-loading is set.
	Edges        Rent591HomeDetailBreadcrumbEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Rent591HomeDetailBreadcrumbEdges holds the relations/edges for other nodes in the graph.
type Rent591HomeDetailBreadcrumbEdges struct {
	// Rent591homeDetails holds the value of the rent591home_details edge.
	Rent591homeDetails []*Rent591HomeDetail `json:"rent591home_details,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// Rent591homeDetailsOrErr returns the Rent591homeDetails value or an error if the edge
// was not loaded in eager-loading.
func (e Rent591HomeDetailBreadcrumbEdges) Rent591homeDetailsOrErr() ([]*Rent591HomeDetail, error) {
	if e.loadedTypes[0] {
		return e.Rent591homeDetails, nil
	}
	return nil, &NotLoadedError{edge: "rent591home_details"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Rent591HomeDetailBreadcrumb) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rent591homedetailbreadcrumb.FieldID, rent591homedetailbreadcrumb.FieldPostID:
			values[i] = new(sql.NullInt64)
		case rent591homedetailbreadcrumb.FieldName, rent591homedetailbreadcrumb.FieldQuery, rent591homedetailbreadcrumb.FieldLink:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Rent591HomeDetailBreadcrumb fields.
func (rdb *Rent591HomeDetailBreadcrumb) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rent591homedetailbreadcrumb.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rdb.ID = int(value.Int64)
		case rent591homedetailbreadcrumb.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rdb.Name = value.String
			}
		case rent591homedetailbreadcrumb.FieldPostID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field post_id", values[i])
			} else if value.Valid {
				rdb.PostID = int(value.Int64)
			}
		case rent591homedetailbreadcrumb.FieldQuery:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field query", values[i])
			} else if value.Valid {
				rdb.Query = value.String
			}
		case rent591homedetailbreadcrumb.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				rdb.Link = value.String
			}
		default:
			rdb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Rent591HomeDetailBreadcrumb.
// This includes values selected through modifiers, order, etc.
func (rdb *Rent591HomeDetailBreadcrumb) Value(name string) (ent.Value, error) {
	return rdb.selectValues.Get(name)
}

// QueryRent591homeDetails queries the "rent591home_details" edge of the Rent591HomeDetailBreadcrumb entity.
func (rdb *Rent591HomeDetailBreadcrumb) QueryRent591homeDetails() *Rent591HomeDetailQuery {
	return NewRent591HomeDetailBreadcrumbClient(rdb.config).QueryRent591homeDetails(rdb)
}

// Update returns a builder for updating this Rent591HomeDetailBreadcrumb.
// Note that you need to call Rent591HomeDetailBreadcrumb.Unwrap() before calling this method if this Rent591HomeDetailBreadcrumb
// was returned from a transaction, and the transaction was committed or rolled back.
func (rdb *Rent591HomeDetailBreadcrumb) Update() *Rent591HomeDetailBreadcrumbUpdateOne {
	return NewRent591HomeDetailBreadcrumbClient(rdb.config).UpdateOne(rdb)
}

// Unwrap unwraps the Rent591HomeDetailBreadcrumb entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rdb *Rent591HomeDetailBreadcrumb) Unwrap() *Rent591HomeDetailBreadcrumb {
	_tx, ok := rdb.config.driver.(*txDriver)
	if !ok {
		panic("ent: Rent591HomeDetailBreadcrumb is not a transactional entity")
	}
	rdb.config.driver = _tx.drv
	return rdb
}

// String implements the fmt.Stringer.
func (rdb *Rent591HomeDetailBreadcrumb) String() string {
	var builder strings.Builder
	builder.WriteString("Rent591HomeDetailBreadcrumb(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rdb.ID))
	builder.WriteString("name=")
	builder.WriteString(rdb.Name)
	builder.WriteString(", ")
	builder.WriteString("post_id=")
	builder.WriteString(fmt.Sprintf("%v", rdb.PostID))
	builder.WriteString(", ")
	builder.WriteString("query=")
	builder.WriteString(rdb.Query)
	builder.WriteString(", ")
	builder.WriteString("link=")
	builder.WriteString(rdb.Link)
	builder.WriteByte(')')
	return builder.String()
}

// Rent591HomeDetailBreadcrumbs is a parsable slice of Rent591HomeDetailBreadcrumb.
type Rent591HomeDetailBreadcrumbs []*Rent591HomeDetailBreadcrumb
