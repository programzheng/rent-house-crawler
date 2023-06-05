// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailbrowse"
)

// Rent591HomeDetailBrowse is the model entity for the Rent591HomeDetailBrowse schema.
type Rent591HomeDetailBrowse struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Pc holds the value of the "pc" field.
	Pc int `json:"pc,omitempty"`
	// Mobile holds the value of the "mobile" field.
	Mobile int `json:"mobile,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Rent591HomeDetailBrowseQuery when eager-loading is set.
	Edges        Rent591HomeDetailBrowseEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Rent591HomeDetailBrowseEdges holds the relations/edges for other nodes in the graph.
type Rent591HomeDetailBrowseEdges struct {
	// Rent591homeDetails holds the value of the rent591home_details edge.
	Rent591homeDetails []*Rent591HomeDetail `json:"rent591home_details,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// Rent591homeDetailsOrErr returns the Rent591homeDetails value or an error if the edge
// was not loaded in eager-loading.
func (e Rent591HomeDetailBrowseEdges) Rent591homeDetailsOrErr() ([]*Rent591HomeDetail, error) {
	if e.loadedTypes[0] {
		return e.Rent591homeDetails, nil
	}
	return nil, &NotLoadedError{edge: "rent591home_details"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Rent591HomeDetailBrowse) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rent591homedetailbrowse.FieldID, rent591homedetailbrowse.FieldPc, rent591homedetailbrowse.FieldMobile:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Rent591HomeDetailBrowse fields.
func (rdb *Rent591HomeDetailBrowse) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rent591homedetailbrowse.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rdb.ID = int(value.Int64)
		case rent591homedetailbrowse.FieldPc:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pc", values[i])
			} else if value.Valid {
				rdb.Pc = int(value.Int64)
			}
		case rent591homedetailbrowse.FieldMobile:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				rdb.Mobile = int(value.Int64)
			}
		default:
			rdb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Rent591HomeDetailBrowse.
// This includes values selected through modifiers, order, etc.
func (rdb *Rent591HomeDetailBrowse) Value(name string) (ent.Value, error) {
	return rdb.selectValues.Get(name)
}

// QueryRent591homeDetails queries the "rent591home_details" edge of the Rent591HomeDetailBrowse entity.
func (rdb *Rent591HomeDetailBrowse) QueryRent591homeDetails() *Rent591HomeDetailQuery {
	return NewRent591HomeDetailBrowseClient(rdb.config).QueryRent591homeDetails(rdb)
}

// Update returns a builder for updating this Rent591HomeDetailBrowse.
// Note that you need to call Rent591HomeDetailBrowse.Unwrap() before calling this method if this Rent591HomeDetailBrowse
// was returned from a transaction, and the transaction was committed or rolled back.
func (rdb *Rent591HomeDetailBrowse) Update() *Rent591HomeDetailBrowseUpdateOne {
	return NewRent591HomeDetailBrowseClient(rdb.config).UpdateOne(rdb)
}

// Unwrap unwraps the Rent591HomeDetailBrowse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rdb *Rent591HomeDetailBrowse) Unwrap() *Rent591HomeDetailBrowse {
	_tx, ok := rdb.config.driver.(*txDriver)
	if !ok {
		panic("ent: Rent591HomeDetailBrowse is not a transactional entity")
	}
	rdb.config.driver = _tx.drv
	return rdb
}

// String implements the fmt.Stringer.
func (rdb *Rent591HomeDetailBrowse) String() string {
	var builder strings.Builder
	builder.WriteString("Rent591HomeDetailBrowse(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rdb.ID))
	builder.WriteString("pc=")
	builder.WriteString(fmt.Sprintf("%v", rdb.Pc))
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(fmt.Sprintf("%v", rdb.Mobile))
	builder.WriteByte(')')
	return builder.String()
}

// Rent591HomeDetailBrowses is a parsable slice of Rent591HomeDetailBrowse.
type Rent591HomeDetailBrowses []*Rent591HomeDetailBrowse
