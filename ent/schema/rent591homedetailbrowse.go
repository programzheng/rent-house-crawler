package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetailBrowse holds the schema definition for the Rent591HomeDetailBrowse entity.
type Rent591HomeDetailBrowse struct {
	ent.Schema
}

// Fields of the Rent591HomeDetailBrowse.
func (Rent591HomeDetailBrowse) Fields() []ent.Field {
	return []ent.Field{
		field.Int("pc"),
		field.Int("mobile"),
	}
}

// Edges of the Rent591HomeDetailBrowse.
func (Rent591HomeDetailBrowse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rent591home_details", Rent591HomeDetail.Type).
			Ref("rent591home_detail_browses"),
	}
}
