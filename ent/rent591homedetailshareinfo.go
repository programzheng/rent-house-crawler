// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailshareinfo"
)

// Rent591HomeDetailShareInfo is the model entity for the Rent591HomeDetailShareInfo schema.
type Rent591HomeDetailShareInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// From holds the value of the "From" field.
	From string `json:"From,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Rent591HomeDetailShareInfoQuery when eager-loading is set.
	Edges        Rent591HomeDetailShareInfoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Rent591HomeDetailShareInfoEdges holds the relations/edges for other nodes in the graph.
type Rent591HomeDetailShareInfoEdges struct {
	// Rent591homeDetails holds the value of the rent591home_details edge.
	Rent591homeDetails []*Rent591HomeDetail `json:"rent591home_details,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// Rent591homeDetailsOrErr returns the Rent591homeDetails value or an error if the edge
// was not loaded in eager-loading.
func (e Rent591HomeDetailShareInfoEdges) Rent591homeDetailsOrErr() ([]*Rent591HomeDetail, error) {
	if e.loadedTypes[0] {
		return e.Rent591homeDetails, nil
	}
	return nil, &NotLoadedError{edge: "rent591home_details"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Rent591HomeDetailShareInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rent591homedetailshareinfo.FieldID:
			values[i] = new(sql.NullInt64)
		case rent591homedetailshareinfo.FieldURL, rent591homedetailshareinfo.FieldFrom, rent591homedetailshareinfo.FieldTitle:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Rent591HomeDetailShareInfo fields.
func (rdsi *Rent591HomeDetailShareInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rent591homedetailshareinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rdsi.ID = int(value.Int64)
		case rent591homedetailshareinfo.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				rdsi.URL = value.String
			}
		case rent591homedetailshareinfo.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field From", values[i])
			} else if value.Valid {
				rdsi.From = value.String
			}
		case rent591homedetailshareinfo.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				rdsi.Title = value.String
			}
		default:
			rdsi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Rent591HomeDetailShareInfo.
// This includes values selected through modifiers, order, etc.
func (rdsi *Rent591HomeDetailShareInfo) Value(name string) (ent.Value, error) {
	return rdsi.selectValues.Get(name)
}

// QueryRent591homeDetails queries the "rent591home_details" edge of the Rent591HomeDetailShareInfo entity.
func (rdsi *Rent591HomeDetailShareInfo) QueryRent591homeDetails() *Rent591HomeDetailQuery {
	return NewRent591HomeDetailShareInfoClient(rdsi.config).QueryRent591homeDetails(rdsi)
}

// Update returns a builder for updating this Rent591HomeDetailShareInfo.
// Note that you need to call Rent591HomeDetailShareInfo.Unwrap() before calling this method if this Rent591HomeDetailShareInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (rdsi *Rent591HomeDetailShareInfo) Update() *Rent591HomeDetailShareInfoUpdateOne {
	return NewRent591HomeDetailShareInfoClient(rdsi.config).UpdateOne(rdsi)
}

// Unwrap unwraps the Rent591HomeDetailShareInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rdsi *Rent591HomeDetailShareInfo) Unwrap() *Rent591HomeDetailShareInfo {
	_tx, ok := rdsi.config.driver.(*txDriver)
	if !ok {
		panic("ent: Rent591HomeDetailShareInfo is not a transactional entity")
	}
	rdsi.config.driver = _tx.drv
	return rdsi
}

// String implements the fmt.Stringer.
func (rdsi *Rent591HomeDetailShareInfo) String() string {
	var builder strings.Builder
	builder.WriteString("Rent591HomeDetailShareInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rdsi.ID))
	builder.WriteString("url=")
	builder.WriteString(rdsi.URL)
	builder.WriteString(", ")
	builder.WriteString("From=")
	builder.WriteString(rdsi.From)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(rdsi.Title)
	builder.WriteByte(')')
	return builder.String()
}

// Rent591HomeDetailShareInfos is a parsable slice of Rent591HomeDetailShareInfo.
type Rent591HomeDetailShareInfos []*Rent591HomeDetailShareInfo