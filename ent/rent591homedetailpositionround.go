// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetail"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailpositionround"
)

// Rent591HomeDetailPositionRound is the model entity for the Rent591HomeDetailPositionRound schema.
type Rent591HomeDetailPositionRound struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Active holds the value of the "active" field.
	Active int `json:"active,omitempty"`
	// CommunityName holds the value of the "community_name" field.
	CommunityName string `json:"community_name,omitempty"`
	// CommunityID holds the value of the "community_id" field.
	CommunityID int `json:"community_id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Lat holds the value of the "lat" field.
	Lat float64 `json:"lat,omitempty"`
	// Lng holds the value of the "lng" field.
	Lng float64 `json:"lng,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Rent591HomeDetailPositionRoundQuery when eager-loading is set.
	Edges                                                 Rent591HomeDetailPositionRoundEdges `json:"edges"`
	rent591home_detail_rent591home_detail_position_rounds *int
	selectValues                                          sql.SelectValues
}

// Rent591HomeDetailPositionRoundEdges holds the relations/edges for other nodes in the graph.
type Rent591HomeDetailPositionRoundEdges struct {
	// Rent591homeDetails holds the value of the rent591home_details edge.
	Rent591homeDetails *Rent591HomeDetail `json:"rent591home_details,omitempty"`
	// Rent591homeDetailPositionRoundDatas holds the value of the rent591home_detail_position_round_datas edge.
	Rent591homeDetailPositionRoundDatas []*Rent591HomeDetailPositionRoundData `json:"rent591home_detail_position_round_datas,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// Rent591homeDetailsOrErr returns the Rent591homeDetails value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Rent591HomeDetailPositionRoundEdges) Rent591homeDetailsOrErr() (*Rent591HomeDetail, error) {
	if e.loadedTypes[0] {
		if e.Rent591homeDetails == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: rent591homedetail.Label}
		}
		return e.Rent591homeDetails, nil
	}
	return nil, &NotLoadedError{edge: "rent591home_details"}
}

// Rent591homeDetailPositionRoundDatasOrErr returns the Rent591homeDetailPositionRoundDatas value or an error if the edge
// was not loaded in eager-loading.
func (e Rent591HomeDetailPositionRoundEdges) Rent591homeDetailPositionRoundDatasOrErr() ([]*Rent591HomeDetailPositionRoundData, error) {
	if e.loadedTypes[1] {
		return e.Rent591homeDetailPositionRoundDatas, nil
	}
	return nil, &NotLoadedError{edge: "rent591home_detail_position_round_datas"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Rent591HomeDetailPositionRound) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rent591homedetailpositionround.FieldLat, rent591homedetailpositionround.FieldLng:
			values[i] = new(sql.NullFloat64)
		case rent591homedetailpositionround.FieldID, rent591homedetailpositionround.FieldActive, rent591homedetailpositionround.FieldCommunityID:
			values[i] = new(sql.NullInt64)
		case rent591homedetailpositionround.FieldTitle, rent591homedetailpositionround.FieldKey, rent591homedetailpositionround.FieldCommunityName, rent591homedetailpositionround.FieldAddress:
			values[i] = new(sql.NullString)
		case rent591homedetailpositionround.ForeignKeys[0]: // rent591home_detail_rent591home_detail_position_rounds
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Rent591HomeDetailPositionRound fields.
func (rdpr *Rent591HomeDetailPositionRound) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rent591homedetailpositionround.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rdpr.ID = int(value.Int64)
		case rent591homedetailpositionround.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				rdpr.Title = value.String
			}
		case rent591homedetailpositionround.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				rdpr.Key = value.String
			}
		case rent591homedetailpositionround.FieldActive:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				rdpr.Active = int(value.Int64)
			}
		case rent591homedetailpositionround.FieldCommunityName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field community_name", values[i])
			} else if value.Valid {
				rdpr.CommunityName = value.String
			}
		case rent591homedetailpositionround.FieldCommunityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field community_id", values[i])
			} else if value.Valid {
				rdpr.CommunityID = int(value.Int64)
			}
		case rent591homedetailpositionround.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				rdpr.Address = value.String
			}
		case rent591homedetailpositionround.FieldLat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field lat", values[i])
			} else if value.Valid {
				rdpr.Lat = value.Float64
			}
		case rent591homedetailpositionround.FieldLng:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field lng", values[i])
			} else if value.Valid {
				rdpr.Lng = value.Float64
			}
		case rent591homedetailpositionround.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field rent591home_detail_rent591home_detail_position_rounds", value)
			} else if value.Valid {
				rdpr.rent591home_detail_rent591home_detail_position_rounds = new(int)
				*rdpr.rent591home_detail_rent591home_detail_position_rounds = int(value.Int64)
			}
		default:
			rdpr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Rent591HomeDetailPositionRound.
// This includes values selected through modifiers, order, etc.
func (rdpr *Rent591HomeDetailPositionRound) Value(name string) (ent.Value, error) {
	return rdpr.selectValues.Get(name)
}

// QueryRent591homeDetails queries the "rent591home_details" edge of the Rent591HomeDetailPositionRound entity.
func (rdpr *Rent591HomeDetailPositionRound) QueryRent591homeDetails() *Rent591HomeDetailQuery {
	return NewRent591HomeDetailPositionRoundClient(rdpr.config).QueryRent591homeDetails(rdpr)
}

// QueryRent591homeDetailPositionRoundDatas queries the "rent591home_detail_position_round_datas" edge of the Rent591HomeDetailPositionRound entity.
func (rdpr *Rent591HomeDetailPositionRound) QueryRent591homeDetailPositionRoundDatas() *Rent591HomeDetailPositionRoundDataQuery {
	return NewRent591HomeDetailPositionRoundClient(rdpr.config).QueryRent591homeDetailPositionRoundDatas(rdpr)
}

// Update returns a builder for updating this Rent591HomeDetailPositionRound.
// Note that you need to call Rent591HomeDetailPositionRound.Unwrap() before calling this method if this Rent591HomeDetailPositionRound
// was returned from a transaction, and the transaction was committed or rolled back.
func (rdpr *Rent591HomeDetailPositionRound) Update() *Rent591HomeDetailPositionRoundUpdateOne {
	return NewRent591HomeDetailPositionRoundClient(rdpr.config).UpdateOne(rdpr)
}

// Unwrap unwraps the Rent591HomeDetailPositionRound entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rdpr *Rent591HomeDetailPositionRound) Unwrap() *Rent591HomeDetailPositionRound {
	_tx, ok := rdpr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Rent591HomeDetailPositionRound is not a transactional entity")
	}
	rdpr.config.driver = _tx.drv
	return rdpr
}

// String implements the fmt.Stringer.
func (rdpr *Rent591HomeDetailPositionRound) String() string {
	var builder strings.Builder
	builder.WriteString("Rent591HomeDetailPositionRound(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rdpr.ID))
	builder.WriteString("title=")
	builder.WriteString(rdpr.Title)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(rdpr.Key)
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", rdpr.Active))
	builder.WriteString(", ")
	builder.WriteString("community_name=")
	builder.WriteString(rdpr.CommunityName)
	builder.WriteString(", ")
	builder.WriteString("community_id=")
	builder.WriteString(fmt.Sprintf("%v", rdpr.CommunityID))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(rdpr.Address)
	builder.WriteString(", ")
	builder.WriteString("lat=")
	builder.WriteString(fmt.Sprintf("%v", rdpr.Lat))
	builder.WriteString(", ")
	builder.WriteString("lng=")
	builder.WriteString(fmt.Sprintf("%v", rdpr.Lng))
	builder.WriteByte(')')
	return builder.String()
}

// Rent591HomeDetailPositionRounds is a parsable slice of Rent591HomeDetailPositionRound.
type Rent591HomeDetailPositionRounds []*Rent591HomeDetailPositionRound
