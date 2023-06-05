// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailtag"
)

// Rent591HomeDetailTag is the model entity for the Rent591HomeDetailTag schema.
type Rent591HomeDetailTag struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PostID holds the value of the "post_id" field.
	PostID int `json:"post_id,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Rent591HomeDetailTagQuery when eager-loading is set.
	Edges        Rent591HomeDetailTagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Rent591HomeDetailTagEdges holds the relations/edges for other nodes in the graph.
type Rent591HomeDetailTagEdges struct {
	// Rent591homeDetails holds the value of the rent591home_details edge.
	Rent591homeDetails []*Rent591HomeDetail `json:"rent591home_details,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// Rent591homeDetailsOrErr returns the Rent591homeDetails value or an error if the edge
// was not loaded in eager-loading.
func (e Rent591HomeDetailTagEdges) Rent591homeDetailsOrErr() ([]*Rent591HomeDetail, error) {
	if e.loadedTypes[0] {
		return e.Rent591homeDetails, nil
	}
	return nil, &NotLoadedError{edge: "rent591home_details"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Rent591HomeDetailTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rent591homedetailtag.FieldID, rent591homedetailtag.FieldPostID:
			values[i] = new(sql.NullInt64)
		case rent591homedetailtag.FieldValue:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Rent591HomeDetailTag fields.
func (rdt *Rent591HomeDetailTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rent591homedetailtag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rdt.ID = int(value.Int64)
		case rent591homedetailtag.FieldPostID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field post_id", values[i])
			} else if value.Valid {
				rdt.PostID = int(value.Int64)
			}
		case rent591homedetailtag.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				rdt.Value = value.String
			}
		default:
			rdt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Rent591HomeDetailTag.
// This includes values selected through modifiers, order, etc.
func (rdt *Rent591HomeDetailTag) GetValue(name string) (ent.Value, error) {
	return rdt.selectValues.Get(name)
}

// QueryRent591homeDetails queries the "rent591home_details" edge of the Rent591HomeDetailTag entity.
func (rdt *Rent591HomeDetailTag) QueryRent591homeDetails() *Rent591HomeDetailQuery {
	return NewRent591HomeDetailTagClient(rdt.config).QueryRent591homeDetails(rdt)
}

// Update returns a builder for updating this Rent591HomeDetailTag.
// Note that you need to call Rent591HomeDetailTag.Unwrap() before calling this method if this Rent591HomeDetailTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (rdt *Rent591HomeDetailTag) Update() *Rent591HomeDetailTagUpdateOne {
	return NewRent591HomeDetailTagClient(rdt.config).UpdateOne(rdt)
}

// Unwrap unwraps the Rent591HomeDetailTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rdt *Rent591HomeDetailTag) Unwrap() *Rent591HomeDetailTag {
	_tx, ok := rdt.config.driver.(*txDriver)
	if !ok {
		panic("ent: Rent591HomeDetailTag is not a transactional entity")
	}
	rdt.config.driver = _tx.drv
	return rdt
}

// String implements the fmt.Stringer.
func (rdt *Rent591HomeDetailTag) String() string {
	var builder strings.Builder
	builder.WriteString("Rent591HomeDetailTag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rdt.ID))
	builder.WriteString("post_id=")
	builder.WriteString(fmt.Sprintf("%v", rdt.PostID))
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(rdt.Value)
	builder.WriteByte(')')
	return builder.String()
}

// Rent591HomeDetailTags is a parsable slice of Rent591HomeDetailTag.
type Rent591HomeDetailTags []*Rent591HomeDetailTag