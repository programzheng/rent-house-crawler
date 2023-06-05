package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetailTag holds the schema definition for the Rent591HomeDetailTag entity.
type Rent591HomeDetailTag struct {
	ent.Schema
}

// Fields of the Rent591HomeDetailTag.
func (Rent591HomeDetailTag) Fields() []ent.Field {
	return []ent.Field{
		//original id
		field.Int("post_id"),
		field.String("value"),
	}
}

// Edges of the Rent591HomeDetailTag.
func (Rent591HomeDetailTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rent591home_details", Rent591HomeDetail.Type),
	}
}
