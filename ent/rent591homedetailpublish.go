// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailpublish"
)

// Rent591HomeDetailPublish is the model entity for the Rent591HomeDetailPublish schema.
type Rent591HomeDetailPublish struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PostID holds the value of the "post_id" field.
	PostID int `json:"post_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// PostTime holds the value of the "post_time" field.
	PostTime string `json:"post_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime string `json:"update_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Rent591HomeDetailPublishQuery when eager-loading is set.
	Edges        Rent591HomeDetailPublishEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Rent591HomeDetailPublishEdges holds the relations/edges for other nodes in the graph.
type Rent591HomeDetailPublishEdges struct {
	// Rent591homeDetails holds the value of the rent591home_details edge.
	Rent591homeDetails []*Rent591HomeDetail `json:"rent591home_details,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// Rent591homeDetailsOrErr returns the Rent591homeDetails value or an error if the edge
// was not loaded in eager-loading.
func (e Rent591HomeDetailPublishEdges) Rent591homeDetailsOrErr() ([]*Rent591HomeDetail, error) {
	if e.loadedTypes[0] {
		return e.Rent591homeDetails, nil
	}
	return nil, &NotLoadedError{edge: "rent591home_details"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Rent591HomeDetailPublish) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rent591homedetailpublish.FieldID, rent591homedetailpublish.FieldPostID:
			values[i] = new(sql.NullInt64)
		case rent591homedetailpublish.FieldName, rent591homedetailpublish.FieldKey, rent591homedetailpublish.FieldPostTime, rent591homedetailpublish.FieldUpdateTime:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Rent591HomeDetailPublish fields.
func (rdp *Rent591HomeDetailPublish) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rent591homedetailpublish.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rdp.ID = int(value.Int64)
		case rent591homedetailpublish.FieldPostID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field post_id", values[i])
			} else if value.Valid {
				rdp.PostID = int(value.Int64)
			}
		case rent591homedetailpublish.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rdp.Name = value.String
			}
		case rent591homedetailpublish.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				rdp.Key = value.String
			}
		case rent591homedetailpublish.FieldPostTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_time", values[i])
			} else if value.Valid {
				rdp.PostTime = value.String
			}
		case rent591homedetailpublish.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				rdp.UpdateTime = value.String
			}
		default:
			rdp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Rent591HomeDetailPublish.
// This includes values selected through modifiers, order, etc.
func (rdp *Rent591HomeDetailPublish) Value(name string) (ent.Value, error) {
	return rdp.selectValues.Get(name)
}

// QueryRent591homeDetails queries the "rent591home_details" edge of the Rent591HomeDetailPublish entity.
func (rdp *Rent591HomeDetailPublish) QueryRent591homeDetails() *Rent591HomeDetailQuery {
	return NewRent591HomeDetailPublishClient(rdp.config).QueryRent591homeDetails(rdp)
}

// Update returns a builder for updating this Rent591HomeDetailPublish.
// Note that you need to call Rent591HomeDetailPublish.Unwrap() before calling this method if this Rent591HomeDetailPublish
// was returned from a transaction, and the transaction was committed or rolled back.
func (rdp *Rent591HomeDetailPublish) Update() *Rent591HomeDetailPublishUpdateOne {
	return NewRent591HomeDetailPublishClient(rdp.config).UpdateOne(rdp)
}

// Unwrap unwraps the Rent591HomeDetailPublish entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rdp *Rent591HomeDetailPublish) Unwrap() *Rent591HomeDetailPublish {
	_tx, ok := rdp.config.driver.(*txDriver)
	if !ok {
		panic("ent: Rent591HomeDetailPublish is not a transactional entity")
	}
	rdp.config.driver = _tx.drv
	return rdp
}

// String implements the fmt.Stringer.
func (rdp *Rent591HomeDetailPublish) String() string {
	var builder strings.Builder
	builder.WriteString("Rent591HomeDetailPublish(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rdp.ID))
	builder.WriteString("post_id=")
	builder.WriteString(fmt.Sprintf("%v", rdp.PostID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(rdp.Name)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(rdp.Key)
	builder.WriteString(", ")
	builder.WriteString("post_time=")
	builder.WriteString(rdp.PostTime)
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(rdp.UpdateTime)
	builder.WriteByte(')')
	return builder.String()
}

// Rent591HomeDetailPublishes is a parsable slice of Rent591HomeDetailPublish.
type Rent591HomeDetailPublishes []*Rent591HomeDetailPublish
